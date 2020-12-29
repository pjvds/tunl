package main

import (
	"log"
	"os"

	"github.com/pjvds/tunl/cmd"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "tunl",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Value: "http://_.tunl.es",
			},
		},
		Usage: "expose local file and services via a public tunnel",
		Commands: []*cli.Command{
			cmd.DaemonCommand,
			cmd.FilesCommand,
			cmd.HttpCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
