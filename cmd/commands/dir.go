package commands

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/gorilla/handlers"
	"github.com/pjvds/tunl/pkg/tunnel"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var DirCommand = &cli.Command{
	Name: "dir",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "access-log",
			Value: true,
		},
		&cli.BoolFlag{
			Name:  "qr",
			Usage: "Print QR code of the public address",
		},
	},
	Usage:     "Expose a directory via a public http address",
	ArgsUsage: "[dir]",
	Action: func(ctx *cli.Context) error {
		dir := ctx.Args().First()
		if len(dir) == 0 {
			dir = "."
		}

		absDir, err := filepath.Abs(dir)
		if err != nil {
			return cli.Exit("invalid dir: "+err.Error(), 1)
		}

		stat, err := os.Stat(absDir)
		if err != nil {
			if os.IsNotExist(err) {
				return cli.Exit("directory doesn't exist", 1)
			}

			return cli.Exit(err.Error(), 1)
		}

		if !stat.IsDir() {
			return cli.Exit(dir+" not a directory", 1)
		}

		host := ctx.String("host")
		if len(host) == 0 {
			fmt.Print("Host cannot be empty\nSee --host flag for more information.\n\n")

			cli.ShowCommandHelpAndExit(ctx, ctx.Command.Name, 1)
			return cli.Exit("Host cannot be empty.", 1)
		}

		hostURL, err := url.Parse(host)
		if err != nil {
			fmt.Printf("Host value invalid: %v\nSee --host flag for more information.\n\n", err)

			cli.ShowCommandHelpAndExit(ctx, ctx.Command.Name, 1)
			return nil
		}

		hostnameWithoutPort := hostURL.Hostname()
		if len(hostnameWithoutPort) == 0 {
			fmt.Print("Host hostname cannot be empty, see --host flag for more information.\n\n")

			cli.ShowCommandHelpAndExit(ctx, ctx.Command.Name, 1)
			return nil
		}

		tunnel, err := tunnel.OpenHTTP(ctx.Context, zap.NewNop(), hostURL)
		if err != nil {
			return cli.Exit(err.Error(), 18)
		}

		handler := http.FileServer(http.Dir(absDir))

		if ctx.Bool("access-log") {
			handler = handlers.LoggingHandler(os.Stderr, handler)
		}

		PrintTunnel(ctx, tunnel.Address(), absDir)

		go func() {
			for state := range tunnel.StateChanges() {
				println(state)
			}
		}()

		if err := http.Serve(tunnel, handler); err != nil {
			return err
		}

		return nil
	},
}
