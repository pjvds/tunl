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
	ArgsUsage: "<container> [port]",
	Usage:     "Create a TCP tunnel to a local docker container",
	BashComplete: cli.BashCompleteFunc(func(ctx *cli.Context) {
		docker, err := client.NewClientWithOpts(client.FromEnv)
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
		id := ctx.Args().First()
		if len(id) == 0 {
			fmt.Print("Missing container id argument.\n\n")

			cli.ShowCommandHelpAndExit(ctx, ctx.Command.Name, 1)
			return cli.Exit("Host cannot be empty.", 1)
		}

		docker, err := client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			panic(err)
		}
		container, err := docker.ContainerInspect(context.Background(), id)
		if err != nil {
			return cli.Exit(err.Error(), 1)
		}

		var port int

		if p := ctx.Args().Get(1); len(p) > 0 {
			parsed, err := strconv.Atoi(p)
			if err != nil {
				return cli.Exit("Invalid port: "+p, 1)
			}
			port = parsed
		} else {
			if exposedPorts := container.Config.ExposedPorts; len(exposedPorts) > 0 {
				for exposedPort := range exposedPorts {
					port = exposedPort.Int()
				}
			} else {
				return cli.Exit("Missing port argument", 1)
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

		tunnel, err := tunnel.OpenTCP(ctx.Context, zap.NewNop(), hostURL)
		if err != nil {
			return cli.Exit(err.Error(), 18)
		}

		fmt.Printf("%s -> %s:%v\n", tunnel.Address(), container.Name[1:], port)

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
