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
			Name:  "server",
			Value: "localhost:8081",
		},
		&cli.BoolFlag{
			Name:  "access-log",
			Value: true,
		},
	},
	Action: func(ctx *cli.Context) error {
		dir := ctx.Args().First()
		if len(dir) == 0 {
			return cli.Exit("missing directory path", 128)
		}

		conn, err := net.Dial("tcp4", ctx.String("server"))
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return nil
		}
		defer conn.Close()

		request, _ := http.NewRequest(http.MethodConnect, "/", nil)
		request.Host = ctx.String("server")
		request.Header.Add("X-Tunl-Type", "http")

		if err := request.Write(conn); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return nil
		}

		reader := bufio.NewReader(conn)
		response, err := http.ReadResponse(reader, request)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return nil
		}

		if response.StatusCode != http.StatusOK {
			fmt.Fprintln(os.Stderr, response.Status)
			return nil
		}

		hostname := response.Header.Get("X-Tunl-Hostname")
		fmt.Println(hostname)

		session, err := yamux.Client(conn, nil)
		if err != nil {
			fmt.Fprintln(os.Stderr, "mux client creation error: "+err.Error())
			return nil
		}
		defer session.Close()

		handler := handlers.LoggingHandler(os.Stdout, http.FileServer(http.Dir(dir)))

		if err := http.Serve(session, handler); err != nil {
			return cli.Exit("serve error: "+err.Error(), 1)
		}

		return nil
	},
}
