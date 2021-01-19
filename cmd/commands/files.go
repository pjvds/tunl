package commands

import (
	"github.com/urfave/cli/v2"
)

var FilesCommand = &cli.Command{
	Name: "files",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "access-log",
			Value: true,
		},
		&cli.BoolFlag{
			Name:  "qr",
			Usage: "Print QR code of the public address",
		},
	},
	Hidden:    true,
	Usage:     "Expose a directory via a public http address",
	ArgsUsage: "[dir]",
	Action: func(ctx *cli.Context) error {
		println("FILES IS DEPRECATED AND WILL BE REMOVED SOON, USE DIR COMMAND")
		return DirCommand.Action(ctx)
	},
}
