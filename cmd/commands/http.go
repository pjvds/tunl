package commands

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/pjvds/tunl/pkg/tunnel"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var HttpCommand = &cli.Command{
	Name: "http",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "access-log",
			Value: true,
		},
	},
	ArgsUsage: "<url>",
	Usage:     "Create a HTTP tunnel",
	Action: func(ctx *cli.Context) error {
		var targetURL *url.URL
		target := ctx.Args().First()
		if len(target) == 0 {
			fmt.Fprint(os.Stderr, "You must specify the <url> argument\n\n")
			cli.ShowCommandHelpAndExit(ctx, ctx.Command.Name, 1)
		}

		if !strings.Contains(target, "://") {
			target = "http://" + target
		}

		parsed, err := url.Parse(target)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid <url> argument value: %v\n\n", err)
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
			fmt.Fprint(os.Stderr, "Host cannot be empty\nSee --host flag for more information.\n\n")

			cli.ShowCommandHelpAndExit(ctx, ctx.Command.Name, 1)
			return cli.Exit("Host cannot be empty.", 1)
		}

		hostURL, err := url.Parse(host)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Host value invalid: %v\nSee --host flag for more information.\n\n", err)

			cli.ShowCommandHelpAndExit(ctx, ctx.Command.Name, 1)
			return nil
		}

		hostnameWithoutPort := hostURL.Hostname()
		if len(hostnameWithoutPort) == 0 {
			fmt.Fprintf(os.Stderr, "Host hostname cannot be empty, see --host flag for more information.\n\n")

			cli.ShowCommandHelpAndExit(ctx, ctx.Command.Name, 1)
			return nil
		}

		tunnel, err := tunnel.OpenHTTP(ctx.Context, zap.NewNop(), hostURL)
		if err != nil {
			return cli.Exit(err.Error(), 18)
		}

		fmt.Printf("%v -> %v\n", tunnel.Address(), target)

		handler := handlers.LoggingHandler(os.Stdout, proxy)

		if err := http.Serve(tunnel, handler); err != nil {
			return cli.Exit("fatal error: "+err.Error(), 1)
		}

		return nil
	},
}
