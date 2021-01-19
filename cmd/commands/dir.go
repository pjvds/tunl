package commands

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/goji/httpauth"
	"github.com/gorilla/handlers"
	"github.com/gorilla/sessions"
	"github.com/pjvds/tunl/pkg/templates"
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
		&cli.StringFlag{
			Name: "password",
		},
		&cli.StringFlag{
			Name:  "basic-auth",
			Usage: "Adds HTTP basic access authentication",
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

		handler := http.FileServer(http.Dir(absDir))

		if password := ctx.String("password"); len(password) > 0 {
			next := handler

			var store = sessions.NewCookieStore([]byte("foobar"))

			handler = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				session, _ := store.Get(r, "tunl-pass")

				if session.Values["pass"] == password {
					next.ServeHTTP(rw, r)
					return
				}

				if r.Method == http.MethodPost {
					if r.ParseForm(); err != nil {
						http.Error(rw, err.Error(), http.StatusBadRequest)
						return
					}

					pass := r.FormValue("password")
					if pass != password {
						templates.Password(rw, templates.PasswordInput{
							Message: "invalid password",
						})
						return
					}

					session.Values["pass"] = pass
					session.Save(r, rw)

					http.Redirect(rw, r, "/", http.StatusSeeOther)
					return
				}
				templates.Password(rw, templates.PasswordInput{})

			})

		}

		if basicAuth := ctx.String("basic-auth"); len(basicAuth) > 0 {
			split := strings.Split(basicAuth, ":")
			if len(split) != 2 {
				return cli.Exit("invalid basic-auth value", 1)
			}

			user := split[0]
			password := split[1]

			if len(user) == 0 {
				return cli.Exit("invalid basic-auth value: empty user", 1)
			}
			if len(password) == 0 {
				return cli.Exit("invalid basic-auth value: empty password", 1)
			}

			handler = httpauth.SimpleBasicAuth(user, password)(handler)
		}

		if ctx.Bool("access-log") {
			handler = handlers.LoggingHandler(os.Stderr, handler)
		}

		tunnel, err := tunnel.OpenHTTP(ctx.Context, zap.NewNop(), hostURL)
		if err != nil {
			return cli.Exit(err.Error(), 18)
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
