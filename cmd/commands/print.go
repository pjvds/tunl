package commands

import (
	"fmt"
	"os"

	"github.com/mdp/qrterminal/v3"
	"github.com/urfave/cli/v2"
)

// PrintTunnel prints public address to stdout and target to stderr
func PrintTunnel(ctx *cli.Context, publicAddress string, target string) {
	fmt.Fprint(os.Stdout, publicAddress)
	fmt.Fprintln(os.Stderr, " ->", target)

	if ctx.Bool("qr") {
		qrterminal.GenerateHalfBlock(publicAddress, qrterminal.L, os.Stdout)
	}

	CopyAddressToClipboardIfRequired(ctx, publicAddress)
}
