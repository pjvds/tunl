package commands

import (
	"github.com/urfave/cli/v2"
)

var VersionCommand = &cli.Command{
	Name:  "version",
	Usage: "Print version information",
	Action: func(ctx *cli.Context) error {
		cli.VersionPrinter(ctx)
		return nil
	},
}
