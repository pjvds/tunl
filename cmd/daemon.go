package cmd

import (
	"crypto/tls"
	"io"
	"net/url"

	"bufio"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/hashicorp/yamux"
	"github.com/inconshreveable/go-vhost"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"github.com/yelinaung/go-haikunator"
	"go.uber.org/zap"
)

var DaemonCommand = &cli.Command{
	Name:   "daemon",
	Hidden: true,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "bind",
			Value: ":8080",
		},
		&cli.StringFlag{
			Name: "tls-cert-file",
		},
		&cli.StringFlag{
			Name: "tls-key-file",
		},
		&cli.StringFlag{
			Name:  "control",
			Value: "_.tunl.es",
		},
		&cli.StringFlag{
			Name:  "domain",
			Value: "tunl.es",
		},
	},
	Action: func(ctx *cli.Context) error {
		bind := ctx.String("bind")
		if len(bind) == 0 {
			logger.Error("bind flag value cannot be empty")
			return nil
		}

		var mux *vhost.VhostMuxer

		if certFile := ctx.String("tls-cert-file"); len(certFile) > 0 {
			keyFile := ctx.String("tls-key-file")

			cert, err := tls.LoadX509KeyPair(certFile, keyFile)
			if err != nil {
				logger.Error("load certificate error", zap.Error(err), zap.String("cert", certFile), zap.String("key", keyFile))
				return nil
			}

			tlsListener, err := tls.Listen("tcp", bind, &tls.Config{
				Certificates: []tls.Certificate{cert},
			})
			if err != nil {
				logger.Error("listen error failed to listen", zap.Error(err), zap.String("bind", bind))
				return nil
			}

			tlsMux, err := vhost.NewHTTPMuxer(tlsListener, 30*time.Second)
			if err != nil {
				logger.Error("vhost tls mux creation error", zap.Error(err))
				return nil
			}
			defer mux.Close()

			mux = tlsMux.VhostMuxer
		} else {
			nonTlsListener, err := net.Listen("tcp", bind)
			if err != nil {
				logger.Error("listen error failed to listen", zap.Error(err), zap.String("bind", bind))
				return nil
			}

			httpMux, err := vhost.NewHTTPMuxer(nonTlsListener, 30*time.Second)
			if err != nil {
				logger.Error("vhost http mux creation error", zap.Error(err))
				return nil
			}
			defer httpMux.Close()

			mux = httpMux.VhostMuxer
		}

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

				hostname := haikunator.Haikunate() + "." + ctx.String("domain")
				vhost, err := mux.Listen(hostname)
				if err != nil {
					logger.Error("vhost listen error", zap.Error(err))
					return
				}
				defer vhost.Close()

				accepted := &http.Response{
					StatusCode: http.StatusOK,
					Header: http.Header{
						"X-Tunl-Hostname": []string{hostname},
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

				targetURL, err := url.Parse("https://httpstatuses.com")
				if err != nil {
					logger.Error("handshake error", zap.Error(err))
					return
				}

				logger.Debug("handshake success", zap.String("target-url", targetURL.String()))

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
