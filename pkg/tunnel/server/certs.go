package server

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gobwas/glob"
)

func LoadCertificates(paths []string) ([]tls.Certificate, error) {
	certs := []tls.Certificate{}
	keys := []tls.Certificate{}

	for _, p := range paths {
		d, f := filepath.Split(p)

		g := glob.MustCompile(f, '/')

		err := filepath.Walk(d, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !g.Match(info.Name()) {
				return nil
			}

			cert, err := addBlocks(path)
			if err != nil {
				return err
			}

			if len(cert.Certificate) > 0 {
				certs = append(certs, cert)
			}

			if cert.PrivateKey != nil {
				keys = append(keys, cert)
			}

			return nil
		})

		if err != nil {
			return certs, err
		}
	}

	for i := range certs {
		// We don't need to parse the public key for TLS, but we so do anyway
		// to check that it looks sane and matches the private key.
		x509Cert, err := x509.ParseCertificate(certs[i].Certificate[0])
		if err != nil {
			return certs, nil
		}

		switch pub := x509Cert.PublicKey.(type) {
		case *rsa.PublicKey:
			for _, key := range keys {
				priv, ok := key.PrivateKey.(*rsa.PrivateKey)
				if !ok {
					continue
				}
				if pub.N.Cmp(priv.N) != 0 {
					continue
				}

				certs[i].PrivateKey = priv
				break
			}
		case *ecdsa.PublicKey:
			for _, key := range keys {
				priv, ok := key.PrivateKey.(*ecdsa.PrivateKey)
				if !ok {
					continue
				}
				if pub.X.Cmp(priv.X) != 0 || pub.Y.Cmp(priv.Y) != 0 {
					continue
				}

				certs[i].PrivateKey = priv
				break
			}
		case ed25519.PublicKey:
			for _, key := range keys {
				priv, ok := key.PrivateKey.(ed25519.PrivateKey)
				if !ok {
					continue
				}
				if !bytes.Equal(priv.Public().(ed25519.PublicKey), pub) {
					continue
				}

				certs[i].PrivateKey = priv
				break
			}
		default:
			return certs, errors.New("tls: unknown public key algorithm")
		}
	}

	return certs, nil
}

// Attempt to parse the given private key DER block. OpenSSL 0.9.8 generates
// PKCS#1 private keys by default, while OpenSSL 1.0.0 generates PKCS#8 keys.
// OpenSSL ecparam generates SEC1 EC private keys for ECDSA. We try all three.
func parsePrivateKey(der []byte) (crypto.PrivateKey, error) {
	if key, err := x509.ParsePKCS1PrivateKey(der); err == nil {
		return key, nil
	}

	if key, err := x509.ParsePKCS8PrivateKey(der); err == nil {
		switch key := key.(type) {
		case *rsa.PrivateKey, *ecdsa.PrivateKey, ed25519.PrivateKey:
			return key, nil
		default:
			return nil, errors.New("tls: found unknown private key type in PKCS#8 wrapping")
		}
	}

	if key, err := x509.ParseECPrivateKey(der); err == nil {
		return key, nil
	}

	return nil, errors.New("tls: failed to parse private key")
}

func addBlocks(path string) (tls.Certificate, error) {
	cert := tls.Certificate{}

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return cert, err
	}

	for {
		block, rest := pem.Decode(raw)
		if block == nil {
			break
		}
		if block.Type == "CERTIFICATE" {
			cert.Certificate = append(cert.Certificate, block.Bytes)
		} else if !(block.Type == "PRIVATE KEY" || strings.HasSuffix(block.Type, " PRIVATE KEY")) {
		} else if key, err := parsePrivateKey(block.Bytes); err != nil {
			return cert, fmt.Errorf("Failure reading private key from \"%s\": %s", path, err)
		} else {
			cert.PrivateKey = key
		}

		raw = rest
	}

	return cert, nil
}
