package commands

import (
	"fmt"
	"io"
	"net"
	"net/url"
	"os"
	"sync"

	"github.com/pjvds/tunl/pkg/tunnel"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var TcpCommand = &cli.Command{
	Name: "tcp",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "access-log",
			Value: true,
		},
	},
	Usage:     "Expose a TCP service via a public address",
	ArgsUsage: "<host:port>",
	Action: func(ctx *cli.Context) error {
		target := ctx.Args().First()
		if len(target) == 0 {
			fmt.Fprint(os.Stderr, "You must specify the <url> argument\n\n")
			cli.ShowCommandHelpAndExit(ctx, ctx.Command.Name, 1)
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

		tunnel, err := tunnel.OpenTCP(ctx.Context, zap.NewNop(), hostURL)
		if err != nil {
			return cli.Exit(err.Error(), 18)
		}

		PrintTunnel(tunnel.Address(), ctx.Args().First())

		for {
			conn, err := tunnel.Accept()
			if err != nil {
				return cli.Exit("fatal error: "+err.Error(), 1)
			}

			fmt.Println(conn.RemoteAddr())

			go func(conn net.Conn) {
				defer conn.Close()

				target, err := net.Dial("tcp", ctx.Args().First())
				if err != nil {
					println(err.Error())
					return
				}

				var work sync.WaitGroup
				work.Add(2)

				go func() {
					defer work.Done()
					io.Copy(conn, target)
				}()

				go func() {
					defer work.Done()
					io.Copy(target, conn)
				}()

				work.Wait()
			}(conn)
		}
	},
}
