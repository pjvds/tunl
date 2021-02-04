package commands

import (
	"crypto/tls"
	"io"
	"sync"

	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	honeycomb "github.com/getspine/go-metrics-honeycomb"
	"github.com/hashicorp/yamux"
	"github.com/inconshreveable/go-vhost"
	"github.com/pjvds/tunl/pkg/tunnel/certs"
	"github.com/pjvds/tunl/pkg/tunnel/server"
	"github.com/pjvds/tunl/pkg/version"
	"github.com/rcrowley/go-metrics"
	"github.com/rs/xid"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
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
		&cli.StringFlag{
			Name: "metrics.honeycomb.token",
		},
		&cli.StringFlag{
			Name:  "metrics.honeycomb.name",
			Value: "tunl",
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

		if token := ctx.String("metrics.honeycomb.token"); len(token) > 0 {
			dataset := "tunl"
			if value := ctx.String("metrics.honeycomb.dataset"); len(value) > 0 {
				dataset = value
			}

			go honeycomb.Honeycomb(metrics.DefaultRegistry, 10*time.Second, token, dataset, []float64{0.50, 0.75, 0.95, 0.99})
			logger.Info("honeycomb sink configurated", zap.String("dataset", dataset))
		}

		tunnelCount := metrics.GetOrRegisterCounter("tunnel", nil)
		connectionCount := metrics.GetOrRegisterCounter("connections", nil)

		var tlsConfig *tls.Config

		var listener net.Listener
		if certGlobs := ctx.StringSlice("tls-certs"); len(certGlobs) > 0 {
			loadedCerts, err := certs.LoadCertificates(certGlobs)
			if err != nil {
				logger.Error("load certificate error", zap.Error(err), zap.Strings("certs", certGlobs))
				return nil
			}

			tlsConfig = &tls.Config{Certificates: loadedCerts}
		}

		listener, err := net.Listen("tcp", bind)
		if err != nil {
			logger.Error("listen error failed to listen", zap.Error(err), zap.String("bind", bind))
			return nil
		}
		logger.Debug("listener created", zap.String("address", listener.Addr().String()))

		mux, err := vhost.NewVhostMuxer(listener, func(c net.Conn) (vhost.Conn, error) {
			tls, tlsErr := vhost.TLS(c)
			if tlsErr == nil {
				return tls, nil
			}
			http, httpErr := vhost.HTTP(tls)
			if httpErr == nil {
				return http, nil
			}

			logger.Debug("connection is not a TLS connection", zap.Error(tlsErr))
			logger.Debug("connection is not a HTTP connection", zap.Error(httpErr))

			return nil, errors.New("unsupported protocol for multiplexing")
		}, 30*time.Second)
		if err != nil {
			logger.Error("vhost mux creation error", zap.Error(err))
			return nil
		}
		defer mux.Close()

		addresses := server.NewAddresses(logger, ctx.String("domain"), tlsConfig, mux)

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

				tunlType := request.Header.Get("X-Tunl-Type")
				tunlToken := request.Header.Get("X-Tunl-Token")

				var address *server.PublicAddress

				if len(tunlToken) > 0 {
					token, err := verifyToken(signKey, tunlToken)
					if err != nil {
						http.Error(response, err.Error(), http.StatusInternalServerError)
						return
					}

					address, err = addresses.ClaimAddress(tunlType, token.Subject)
					if err != nil {
						http.Error(response, err.Error(), http.StatusInternalServerError)
						return
					}

				} else {
					var err error
					address, err = addresses.NewAddress(tunlType)
					if err != nil {
						http.Error(response, err.Error(), http.StatusInternalServerError)
						return
					}
				}

				defer address.Close()

				token, err := createToken(signKey, address.Address)
				if err != nil {
					logger.Error("failed to create token", zap.Error(err))
					http.Error(response, "internal server error", http.StatusInternalServerError)
					return
				}

				accept := &http.Response{
					StatusCode: http.StatusOK,
					Header: http.Header{
						"X-Tunl-Token":   []string{token},
						"X-Tunl-Address": []string{address.Address},
						"X-Tunl-Version": []string{version.String()},
					},
				}

				if err := accept.Write(conn); err != nil {
					logger.Error("failed to write success response", zap.Error(err))
					return
				}

				session, err := yamux.Server(conn, nil)
				if err != nil {
					logger.Debug("mux server creation error", zap.Error(err))
					return
				}

				started := time.Now()
				defer func() {
					session.Close()
					logger.Debug("tunnel closed", zap.String("address", address.Address), zap.Duration("time-online", time.Since(started)))
				}()

				accepted := make(chan net.Conn)
				go func() {
					defer close(accepted)
					defer tunnelCount.Dec(1)

					tunnelCount.Inc(1)

					for {
						conn, err := address.Accept()
						if err != nil {
							logger.Debug("accept error", zap.Error(err))
							return
						}

						accepted <- conn
					}
				}()

				for {
					select {
					case <-session.CloseChan():
						logger.Debug("session closed")
						return

					case conn, ok := <-accepted:
						if !ok {
							return
						}

						go func(public net.Conn) {
							defer public.Close()
							defer connectionCount.Dec(1)

							connectionCount.Inc(1)

							logger.Debug("accepted "+public.RemoteAddr().String(), zap.String("address", address.Address))

							local, err := session.Open()
							if err != nil {
								logger.Debug("failed to open session", zap.Error(err))
								return
							}
							defer local.Close()

							var work sync.WaitGroup
							work.Add(2)

							var in int64
							var out int64

							go func() {
								in, _ = io.Copy(local, public)
								work.Done()
							}()

							go func() {
								out, _ = io.Copy(public, local)
								work.Done()
							}()

							work.Wait()
							logger.Debug("connection copy finished", zap.Int64("in", in), zap.Int64("out", out), zap.Stringer("public", public.RemoteAddr()), zap.Stringer("local", local.RemoteAddr()))
						}(conn)
					}
				}
			}))
		}()

		go func() {
			for {
				conn, err := mux.NextError()
				if err != nil {
					logger.Debug("mux error", zap.Error(err))
				}

				if conn != nil {
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
