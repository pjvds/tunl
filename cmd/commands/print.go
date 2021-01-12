package commands

import (
	"fmt"
	"os"
)

// PrintTunnel prints public address to stdout and target to stderr
func PrintTunnel(publicAddress string, target string) {
	fmt.Fprint(os.Stdout, publicAddress)
	fmt.Fprintln(os.Stderr, " ->", target)
}
