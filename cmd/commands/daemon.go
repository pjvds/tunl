package commands

import (
	"bytes"
	"crypto/tls"
	"html/template"
	"io"

	"bufio"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hashicorp/yamux"
	"github.com/inconshreveable/go-vhost"
	"github.com/pjvds/tunl/pkg/tunnel/certs"
	"github.com/rs/xid"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"github.com/yelinaung/go-haikunator"
	"go.uber.org/zap"
)

func createToken(signKey string, id string) (string, error) {
	claims := jwt.StandardClaims{
		Id:        xid.New().String(),
		Issuer:    "tunl",
		Subject:   id,
		Audience:  "tunnels",
		ExpiresAt: time.Now().Add(24 * time.Hour).UTC().Unix(),
		IssuedAt:  time.Now().UTC().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signKey))
}

func verifyToken(signKey string, tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signKey), nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "invalid token")
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if !claims.VerifyExpiresAt(time.Now().UTC().Unix(), false) {
		return nil, errors.New("token expired")
	}

	return claims, nil
}

var DaemonCommand = &cli.Command{
	Name:   "daemon",
	Hidden: true,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "bind",
			Value: ":8080",
		},
		&cli.StringSliceFlag{
			Name: "tls-certs",
		},
		&cli.StringFlag{
			Name:  "control",
			Value: "_.tunl.es",
		},
		&cli.StringFlag{
			Name:  "domain",
			Value: "tunl.es",
		},
		&cli.StringFlag{
			Name:  "address-template",
			Value: "https://{{.Id}}.{{.Domain}}",
		},
		&cli.StringFlag{
			Name:  "sign-key",
			Value: xid.New().String(),
		},
	},
	Action: func(ctx *cli.Context) error {
		signKey := ctx.String("sign-key")
		if len(signKey) == 0 {
			logger.Error("sign-key cannot be empty")
			return nil
		}

		bind := ctx.String("bind")
		if len(bind) == 0 {
			logger.Error("bind flag value cannot be empty")
			return nil
		}

		addressTemplateInput := ctx.String("address-template")
		if len(addressTemplateInput) == 0 {
			logger.Error("address-template value cannot be empty")
			return nil
		}

		addressTemplate, err := template.New("address-template").Parse(addressTemplateInput)
		if err != nil {
			logger.Error("address-template value invalid: " + err.Error())
			return nil
		}

		var listener net.Listener
		if certGlobs := ctx.StringSlice("tls-certs"); len(certGlobs) > 0 {
			certs, err := certs.LoadCertificates(certGlobs)
			if err != nil {
				logger.Error("load certificate error", zap.Error(err), zap.Strings("certs", certGlobs))
				return nil
			}

			tlsListener, err := tls.Listen("tcp", bind, &tls.Config{
				Certificates: certs,
			})
			if err != nil {
				logger.Error("listen error failed to listen", zap.Error(err), zap.String("bind", bind))
				return nil
			}
			listener = tlsListener
		} else {
			nonTlsListener, err := net.Listen("tcp", bind)
			if err != nil {
				logger.Error("listen error failed to listen", zap.Error(err), zap.String("bind", bind))
				return nil
			}
			listener = nonTlsListener
		}

		logger.Debug("listener created", zap.String("address", listener.Addr().String()))

		mux, err := vhost.NewHTTPMuxer(listener, 30*time.Second)
		if err != nil {
			logger.Error("vhost mux creation error", zap.Error(err))
			return nil
		}
		defer mux.Close()

		haikunator := haikunator.New(time.Now().UTC().UnixNano())

		failed := make(chan error)

		go func() {
			logger.Debug("creating control vhost", zap.String("control", ctx.String("control")))

			control, err := mux.Listen(ctx.String("control"))
			if err != nil {
				failed <- errors.Wrap(err, "control vhost listener creation error")
				return
			}

			logger.Debug("control vhost created", zap.String("control", ctx.String("control")))

			failed <- http.Serve(control, http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
				if request.Method != http.MethodConnect {
					http.Error(response, "method not allowed", http.StatusMethodNotAllowed)
					return
				}

				conn, _, err := response.(http.Hijacker).Hijack()
				if err != nil {
					logger.Debug("http hijack error", zap.Error(err))
					http.Error(response, "internal server error", http.StatusInternalServerError)
					return
				}
				defer conn.Close()

				if request.Header.Get("X-Tunl-Type") == "tcp" {
					id := xid.New().String()
					logger.Debug("negociating tcp tunnel")

					tunnelLog := logger.With(zap.String("tunnel-id", id))

					vhost, err := net.Listen("tcp", ":0")
					if err != nil {
						tunnelLog.Debug("tcp listen error", zap.Error(err))
						http.Error(response, "invalid token", http.StatusUnauthorized)
						return
					}
					defer vhost.Close()

					_, port, err := net.SplitHostPort(vhost.Addr().String())
					if err != nil {
						tunnelLog.Debug("cannot get port from vhost address", zap.String("addr", vhost.Addr().String()), zap.Error(err))
						http.Error(response, "invalid token", http.StatusUnauthorized)
						return
					}

					address := fmt.Sprintf("%s:%s", ctx.String("domain"), port)

					accepted := &http.Response{
						StatusCode: http.StatusOK,
						Header: http.Header{
							"X-Tunl-Address": []string{address},
						},
					}
					if err := accepted.Write(conn); err != nil {
						logger.Error("failed to accept control stream", zap.Error(err))
						return
					}

					session, err := yamux.Server(conn, nil)
					if err != nil {
						logger.Debug("mux server creation error", zap.Error(err))
						http.Error(response, "internal server error", http.StatusInternalServerError)
						return
					}
					defer session.Close()

					for {
						accepted, err := vhost.Accept()
						if err != nil {
							break
						}

						go func(accepted net.Conn) {
							tunnelLog.Debug("accepted " + accepted.RemoteAddr().String())
							target, err := session.Open()
							if err != nil {
								return
							}

							go io.Copy(accepted, target)
							go io.Copy(target, accepted)
						}(accepted)
					}

					logger.Debug("tcp tunnel closed", zap.String("address", address))
					return
				}

				id := haikunator.Haikunate()
				if claimedID := request.URL.Query().Get("id"); len(claimedID) > 0 {
					logger.Debug("validating claim for tunnel", zap.String("id", claimedID))
					token := request.Header.Get("X-Tunl-Token")

					claims, err := verifyToken(signKey, token)
					if err != nil {
						logger.Debug("invalid token", zap.String("token", token), zap.Error(err))
						http.Error(response, "invalid token", http.StatusUnauthorized)
						return
					}

					logger.Debug("claim valid", zap.String("id", claimedID))
					id = claims.Subject
				}

				hostname := id + "." + ctx.String("domain")
				buffer := bytes.Buffer{}

				if err := addressTemplate.Execute(&buffer, struct {
					Id     string
					Domain string
				}{Id: id, Domain: ctx.String("domain")}); err != nil {
					logger.Debug("address-template execution error", zap.Error(err))
					http.Error(response, "internal server error", http.StatusInternalServerError)
					return
				}

				token, err := createToken(signKey, id)
				if err != nil {
					logger.Error("failed to create token", zap.Error(err))
					http.Error(response, "internal server error", http.StatusInternalServerError)
				}

				started := time.Now()
				vhost, err := mux.Listen(hostname)
				if err != nil {
					logger.Error("vhost listen error", zap.Error(err))
					return
				}
				defer vhost.Close()

				accepted := &http.Response{
					StatusCode: http.StatusOK,
					Header: http.Header{
						"X-Tunl-Id":      []string{id},
						"X-Tunl-Address": []string{buffer.String()},
						"X-Tunl-Token":   []string{token},
					},
				}
				if err := accepted.Write(conn); err != nil {
					logger.Error("failed to accept control stream", zap.Error(err))
					return
				}

				session, err := yamux.Server(conn, nil)
				if err != nil {
					logger.Debug("mux server creation error", zap.Error(err))
					http.Error(response, "internal server error", http.StatusInternalServerError)
					return
				}
				defer session.Close()

				logger.Debug("handshake success")

				go http.Serve(vhost, http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
					upstream, err := session.Open()
					if err != nil {
						http.Error(response, "bad gateway: "+err.Error(), http.StatusBadGateway)
						http.Error(response, err.Error(), 500)
						return
					}

					defer upstream.Close()
					logger.Debug("session opened", zap.String("remote", upstream.RemoteAddr().String()), zap.String("local", upstream.LocalAddr().String()))

					if err := request.WriteProxy(upstream); err != nil {
						http.Error(response, "failed to write request: "+err.Error(), http.StatusBadGateway)
						return
					}

					upstreamResponse, err := http.ReadResponse(bufio.NewReader(upstream), request)
					if err != nil {
						http.Error(response, "failed to read response: "+err.Error(), http.StatusBadGateway)
						return
					}

					for header, values := range upstreamResponse.Header {
						for _, value := range values {
							response.Header().Add(header, value)
						}
					}

					response.WriteHeader(upstreamResponse.StatusCode)
					if upstreamResponse.Body != nil {
						io.Copy(response, upstreamResponse.Body)
					}
				}))

				<-session.CloseChan()
				logger.Debug("tunnel closed", zap.String("id", id), zap.Duration("time-online", time.Since(started)))
			}))
		}()

		go func() {
			for {
				conn, err := mux.NextError()
				logger.Debug("mux error", zap.Error(err))

				if conn != nil {
					fmt.Fprintln(conn, "500", err.Error())
					conn.Close()
				}
			}
		}()

		if err := <-failed; err != nil {
			logger.Error("fatal error", zap.Error(err))
		}

		return nil
	},
}
