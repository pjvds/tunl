package cmd

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gorilla/handlers"
	"github.com/hashicorp/yamux"
	"github.com/urfave/cli/v2"
)

var HttpCommand = &cli.Command{
	Name: "http",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "server",
			Value: "localhost:8081",
		},
		&cli.BoolFlag{
			Name:  "access-log",
			Value: true,
		},
	},
	Action: func(ctx *cli.Context) error {
		var targetURL *url.URL
		target := ctx.Args().First()
		if len(target) == 0 {
			return cli.Exit("missing target url", 128)
		}

		parsed, err := url.Parse(target)
		if err != nil {
			return cli.Exit("invalid target url", 128)
		}
		targetURL = parsed

		proxy := httputil.NewSingleHostReverseProxy(targetURL)
		originalDirector := proxy.Director

		proxy.Director = func(request *http.Request) {
			originalDirector(request)
			request.Host = targetURL.Host
		}

		conn, err := net.Dial("tcp4", ctx.String("server"))
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return nil
		}
		defer conn.Close()

		session, err := yamux.Client(conn, nil)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return nil
		}
		defer session.Close()

		control, err := session.OpenStream()
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return nil
		}

		control.Write([]byte(targetURL.String() + "\n"))

		reader := bufio.NewReader(control)
		hostname, _, err := reader.ReadLine()

		fmt.Fprintln(os.Stderr, "exposed at: "+string(hostname))

		handler := handlers.LoggingHandler(os.Stdout, proxy)

		if err := http.Serve(session, handler); err != nil {
			return cli.Exit("serve error: "+err.Error(), 1)
		}

		return nil
	},
}
