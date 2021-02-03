package commands

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"github.com/urfave/cli/v2"
)

var CopyToClipboardFlag = &cli.BoolFlag{
	Name:    "copy",
	Aliases: []string{"c"},
	EnvVars: []string{"TUNL_COPY_ADDRESS"},
	Usage:   "Copies the public address to the clipboard",
}

func CopyAddressToClipboardIfRequired(ctx *cli.Context, address string) {
	if err := clipboard.WriteAll(address); err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
	}
}
