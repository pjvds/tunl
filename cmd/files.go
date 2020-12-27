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
		path := ctx.Args().First()
		if len(path) == 0 {
			path = "."
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

		control.Write([]byte("http://tunl-file-server\n"))

		reader := bufio.NewReader(control)
		hostname, _, err := reader.ReadLine()

		fmt.Fprintln(os.Stderr, "exposed at: "+string(hostname))

		fmt.Println("serving path:", path)

		fileserver := http.FileServer(http.Dir(path))
		handler := handlers.LoggingHandler(os.Stdout, fileserver)

		return http.Serve(session, handler)
	},
}
