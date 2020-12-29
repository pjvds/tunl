package cmd

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	"github.com/hashicorp/yamux"
	"github.com/urfave/cli/v2"
)

var FilesCommand = &cli.Command{
	Name: "files",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "host",
			Value:    "_.tunl.es:80",
			Required: true,
		},
		&cli.BoolFlag{
			Name:  "access-log",
			Value: true,
		},
	},
	ArgsUsage: "<dir>",
	Action: func(ctx *cli.Context) error {
		dir := ctx.Args().First()
		if len(dir) == 0 {
			fmt.Println("You must specify the <dir> argument\n")
			cli.ShowCommandHelpAndExit(ctx, ctx.Command.Name, 1)
		}

		host := ctx.String("host")
		if len(host) == 0 {
			fmt.Println("Host cannot be empty, see --host flag for more information.\n")

			cli.ShowCommandHelpAndExit(ctx, ctx.Command.Name, 1)
			return cli.Exit("Host cannot be empty.", 1)
		}

		conn, err := net.Dial("tcp", host)
		if err != nil {
			return cli.Exit(fmt.Sprintf("Failed to connect to host %s: %v", host, err), 128)
		}
		defer conn.Close()

		request, _ := http.NewRequest(http.MethodConnect, "/", nil)
		request.Host = ctx.String("host")
		request.Header.Add("X-Tunl-Type", "http")

		if err := request.Write(conn); err != nil {
			return cli.Exit(fmt.Sprintf("Failed to write connect request: %v", err), 128)
		}

		reader := bufio.NewReader(conn)
		response, err := http.ReadResponse(reader, request)
		if err != nil {
			return cli.Exit(fmt.Sprintf("Failed to read connect response: %v", err), 128)
		}

		if response.StatusCode != http.StatusOK {
			return cli.Exit(fmt.Sprintf("Unexpected connec reponse status %v", response.StatusCode), 128)
		}

		hostname := response.Header.Get("X-Tunl-Hostname")
		fmt.Println(hostname)

		session, err := yamux.Client(conn, nil)
		if err != nil {
			return cli.Exit(fmt.Sprintf("Failed to create multiplex client: %v", err), 128)
		}
		defer session.Close()

		handler := http.FileServer(http.Dir(dir))

		if ctx.Bool("access-log") {
			handler = handlers.LoggingHandler(os.Stderr, handler)
		}

		if err := http.Serve(session, handler); err != nil {
			return cli.Exit("fatal: "+err.Error(), 128)
		}

		return nil
	},
}
