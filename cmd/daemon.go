package cmd

import (
	"io"
	"net/url"

	"bufio"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/yamux"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"github.com/yelinaung/go-haikunator"
	"go.uber.org/zap"
)

type Tunnel struct {
	id      string
	handler http.Handler
	session *yamux.Session
}

var DaemonCommand = &cli.Command{
	Name:   "daemon",
	Hidden: true,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "http",
			Value: ":8080",
		},
		&cli.StringFlag{
			Name:  "control",
			Value: ":8081",
		},
	},
	Action: func(ctx *cli.Context) error {
		bind := ctx.String("control")
		if len(bind) == 0 {
			logger.Error("missing control argument")
			return nil
		}
		logger.Info("binding control", zap.String("control", bind))

		listener, err := net.Listen("tcp", bind)
		if err != nil {
			logger.Error("failed to listen", zap.Error(err), zap.String("bind", bind))
			return nil
		}

		tunnels := make(map[string]*Tunnel)
		haikunator := haikunator.New(time.Now().UTC().UnixNano())

		failed := make(chan error)

		http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
			logger.Debug("handling http request", zap.String("method", request.Method), zap.String("host", request.Host))

			parts := strings.Split(request.Host, ".")
			tunnel, ok := tunnels[parts[0]]
			if !ok {
				http.Error(response, fmt.Sprintf("tunnel %v not found", parts[0]), http.StatusNotFound)
				return
			}

			logger.Debug("invoking http handler of tunnel", zap.String("id", parts[0]))

			tunnel.handler.ServeHTTP(response, request)
		})

		go func() {
			for {
				conn, err := listener.Accept()
				if err != nil {
					failed <- errors.Wrap(err, "failed to accept connection")
					return
				}

				logger.Debug("new connection")

				go func(conn net.Conn) {
					defer conn.Close()

					session, err := yamux.Server(conn, nil)
					if err != nil {
						logger.Debug("mux server creation error", zap.Error(err))
						return
					}
					defer session.Close()

					control, err := session.AcceptStream()
					if err != nil {
						logger.Error("accept control stream error", zap.Error(err))
						return
					}
					defer control.Close()

					reader := bufio.NewReader(control)
					line, _, err := reader.ReadLine()
					if err != nil {
						logger.Error("failed to read handshake", zap.Error(err))
						return
					}

					targetURL, err := url.Parse(string(line))
					if err != nil {
						logger.Error("handshake error", zap.Error(err))
						return
					}

					logger.Debug("handshake success", zap.String("target-url", targetURL.String()))

					var hostname string
					for {
						hostname = haikunator.Haikunate()

						if _, exists := tunnels[hostname]; exists {
							continue
						}
						break
					}

					if _, err := control.Write([]byte(hostname + "\n")); err != nil {
						logger.Error("failed to accept control stream", zap.Error(err))
						return
					}

					tunnel := &Tunnel{
						id:      hostname,
						session: session,
						handler: http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
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
						}),
					}

					logger.Info("tunnel created", zap.String("hostname", hostname))
					tunnels[hostname] = tunnel

					<-session.CloseChan()
					delete(tunnels, hostname)

					logger.Info("session closed", zap.String("hostname", hostname))
				}(conn)
			}
		}()

		go func() {
			logger.Info("binding http listener", zap.String("address", ctx.String("http")))
			failed <- http.ListenAndServe(ctx.String("http"), nil)
		}()

		if err := <-failed; err != nil {
			logger.Error("fatal error", zap.Error(err))
		}

		return nil
	},
}
