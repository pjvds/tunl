package commands

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/handlers"
	"github.com/pjvds/tunl/assets/favicon"
	"github.com/pjvds/tunl/pkg/fallback"

	"github.com/hashicorp/yamux"
	"github.com/pjvds/backoff"
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

		id := ""
		token := ""
		delay := backoff.Exp(1*time.Second, 1*time.Second)

	RECONNECT:
		conn, hostname, err := DialHost(ctx)
		if err != nil {
			fmt.Fprintln(os.Stderr, "connect error:", err)

			delay.Delay()
			goto RECONNECT
		}
		defer conn.Close()

		request, _ := http.NewRequest(http.MethodConnect, "/?id="+url.QueryEscape(id), nil)
		request.Host = hostname
		request.Header.Add("X-Tunl-Type", "http")
		request.Header.Add("X-Tunl-Token", token)

		if err := request.Write(conn); err != nil {
			fmt.Fprintln(os.Stderr, "handshake request error:", err)

			delay.Delay()
			goto RECONNECT
		}

		reader := bufio.NewReader(conn)
		response, err := http.ReadResponse(reader, request)
		if err != nil {
			fmt.Fprintln(os.Stderr, "handshake response error:", err)

			delay.Delay()
			goto RECONNECT

		}

		if response.StatusCode != http.StatusOK {
			return cli.Exit(fmt.Sprintf("Unexpected connect reponse status: %v", response.Status), 128)
		}

		session, err := yamux.Client(conn, nil)
		if err != nil {
			fmt.Fprintln(os.Stdout, "mux creation error: "+err.Error())
			delay.Delay()
			goto RECONNECT
		}

		err = func() error {
			defer session.Close()
			handler := fallback.Fallback(http.FileServer(favicon.AssetFile()), http.FileServer(http.Dir(absDir)))

			if ctx.Bool("access-log") {
				handler = handlers.LoggingHandler(os.Stderr, handler)
			}

			id = response.Header.Get("X-Tunl-Id")
			token = response.Header.Get("X-Tunl-Token")

			fmt.Fprintln(os.Stdout, response.Header.Get("X-Tunl-Address"), "->", absDir)

			if err := http.Serve(session, handler); err != nil {
				return err
			}

			return nil
		}()

		if err != nil {
			fmt.Fprintln(os.Stdout, "disconnected: "+err.Error())
			delay.Delay()
			goto RECONNECT
		}

		return nil
	},
}
