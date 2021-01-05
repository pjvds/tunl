package main

import (
	"log"
	"os"

	"github.com/pjvds/tunl/cmd/commands"
	"github.com/pjvds/tunl/pkg/version"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "tunl",
		Version: version.String(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Value: "https://_.tunl.es",
			},
		},
		Usage: "expose local file and services via a public tunnel",
		Commands: []*cli.Command{
			commands.DaemonCommand,
			commands.FilesCommand,
			commands.HttpCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
