package cmd

import (
	"bufio"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	"github.com/hashicorp/yamux"
	"github.com/urfave/cli/v2"
)

var FilesCommand = &cli.Command{
	Name: "files",
	Flags: []cli.Flag{
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

		conn, hostname, err := DialHost(ctx)
		if err != nil {
			return err
		}
		defer conn.Close()

		request, _ := http.NewRequest(http.MethodConnect, "/", nil)
		request.Host = hostname
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
			return cli.Exit(fmt.Sprintf("Unexpected connect reponse status: %v", response.Status), 128)
		}

		session, err := yamux.Client(conn, nil)
		if err != nil {
			return cli.Exit(fmt.Sprintf("Failed to create multiplex client: %v", err), 128)
		}
		defer session.Close()

		handler := http.FileServer(http.Dir(dir))

		if ctx.Bool("access-log") {
			handler = handlers.LoggingHandler(os.Stderr, handler)
		}

		fmt.Println(response.Header.Get("X-Tunl-Address"))
		if err := http.Serve(session, handler); err != nil {
			return cli.Exit("fatal: "+err.Error(), 128)
		}

		return nil
	},
}
