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
		Name:                 "tunl",
		Version:              version.String(),
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Value: "https://_.tunl.es",
			},
		},
		Usage: "expose your localhost to the public",
		Commands: []*cli.Command{
			commands.DockerCommand,
			commands.DaemonCommand,
			commands.FilesCommand,
			commands.HttpCommand,
			commands.TcpCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
