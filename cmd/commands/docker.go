package commands

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/url"
	"strconv"
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/pjvds/tunl/pkg/tunnel"
	"go.uber.org/zap"

	"github.com/urfave/cli/v2"
)

var DockerCommand = &cli.Command{
	Name:      "docker",
	ArgsUsage: "<container>[:<port>]",
	Flags: []cli.Flag{
		CopyToClipboardFlag,
		&cli.BoolFlag{
			Name:  "qr",
			Usage: "Print QR code of the public address",
		},
		&cli.BoolFlag{
			Name:  "copy-address",
			Usage: "Copies the public address to the clipboard",
		},
		&cli.BoolFlag{
			Name:  "tls",
			Usage: "Open a public TLS address instead of TCP. TLS will be terminated at your local machine tunl client.",
		},
	},
	Usage: "Expose a docker container port via a public address",
	BashComplete: cli.BashCompleteFunc(func(ctx *cli.Context) {
		docker, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			return
		}

		args := filters.NewArgs()

		if ctx.Args().Present() {
			args.FuzzyMatch("name", ctx.Args().First())
		}

		containers, err := docker.ContainerList(context.Background(), types.ContainerListOptions{Filters: args})
		if err != nil {
			return
		}

		for _, container := range containers {
			for _, name := range container.Names {
				fmt.Println(name)
			}

			fmt.Println(container.ID)
		}
	}),
	Action: func(ctx *cli.Context) error {
		containerAndPort := ctx.Args().First()
		if len(containerAndPort) == 0 {
			fmt.Print("Missing container argument.\n\n")

			cli.ShowCommandHelpAndExit(ctx, ctx.Command.Name, 1)
			return cli.Exit("Container cannot be empty.", 1)
		}

		docker, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			panic(err)
		}

		containerSpec, portSpec, err := net.SplitHostPort(containerAndPort)
		if err != nil {
			fmt.Print("Invalid container argument\n\n")

			cli.ShowCommandHelpAndExit(ctx, ctx.Command.Name, 1)
			return cli.Exit(err.Error(), 1)
		}

		container, err := docker.ContainerInspect(context.Background(), containerSpec)
		if err != nil {
			return cli.Exit(err.Error(), 1)
		}

		var port int

		if len(portSpec) > 0 {
			parsed, err := strconv.Atoi(portSpec)
			if err != nil {
				return cli.Exit("Invalid port: "+portSpec, 1)
			}
			port = parsed
		} else {
			if exposedPorts := container.Config.ExposedPorts; len(exposedPorts) > 0 {
				for exposedPort := range exposedPorts {
					port = exposedPort.Int()
				}
			} else {
				return cli.Exit("Missing port argument and no exposed ports found in container", 1)
			}
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

		tunnel, err := tunnel.OpenTCP(ctx.Context, zap.NewNop(), hostURL, ctx.Bool("tls"))
		if err != nil {
			return cli.Exit(err.Error(), 18)
		}

		PrintTunnel(ctx, tunnel.Address(), fmt.Sprintf("%s:%v", container.Name[1:], port))

		go func() {
			for state := range tunnel.StateChanges() {
				println(state)
			}
		}()

		for {
			conn, err := tunnel.Accept()
			if err != nil {
				return cli.Exit("accept error: "+err.Error(), 1)
			}

			fmt.Println(conn.RemoteAddr())

			go func(conn net.Conn) {
				defer conn.Close()

				target, err := net.Dial("tcp", fmt.Sprintf("%s:%v", container.NetworkSettings.IPAddress, port))
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
