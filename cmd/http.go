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
			Name:  "host",
			Value: "_.tunl.es:80",
		},
		&cli.BoolFlag{
			Name:  "access-log",
			Value: true,
		},
	},
	ArgsUsage: "<url>",
	Action: func(ctx *cli.Context) error {
		var targetURL *url.URL
		target := ctx.Args().First()
		if len(target) == 0 {
			fmt.Println("You must specify the <url> argument\n")
			cli.ShowCommandHelpAndExit(ctx, ctx.Command.Name, 1)
		}

		parsed, err := url.Parse(target)
		if err != nil {
			fmt.Printf("Invalid <url> argument value: %v\n\n", err)
			cli.ShowCommandHelpAndExit(ctx, ctx.Command.Name, 1)
		}
		targetURL = parsed

		proxy := httputil.NewSingleHostReverseProxy(targetURL)
		originalDirector := proxy.Director

		proxy.Director = func(request *http.Request) {
			originalDirector(request)
			request.Host = targetURL.Host
		}

		host := ctx.String("host")
		if len(host) == 0 {
			fmt.Println("Host cannot be empty, see --host flag for more information.\n")

			cli.ShowCommandHelpAndExit(ctx, ctx.Command.Name, 1)
			return cli.Exit("Host cannot be empty.", 1)
		}

		conn, err := net.Dial("tcp", host)
		if err != nil {
			return cli.Exit(fmt.Sprintf("Failed connect server: %v", err), 1)
		}
		defer conn.Close()

		request, _ := http.NewRequest(http.MethodConnect, "/", nil)
		request.Host = ctx.String("server")
		request.Header.Add("X-Tunl", targetURL.String())

		if err := request.Write(conn); err != nil {
			return cli.Exit(fmt.Sprintf("Failed to write connect request: %v", err), 1)
		}

		reader := bufio.NewReader(conn)
		response, err := http.ReadResponse(reader, request)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return nil
		}

		if response.StatusCode != http.StatusOK {
			return cli.Exit(fmt.Sprintf("Unexpect connect response status: %v", response.Status), 1)
		}

		hostname := response.Header.Get("X-Tunl-Hostname")
		fmt.Println(hostname)

		session, err := yamux.Client(conn, nil)
		if err != nil {
			return cli.Exit(fmt.Sprintf("Failed to create multiplex client: %v", err), 1)
		}
		defer session.Close()

		handler := handlers.LoggingHandler(os.Stdout, proxy)

		if err := http.Serve(session, handler); err != nil {
			return cli.Exit("fatal error: "+err.Error(), 1)
		}

		return nil
	},
}
