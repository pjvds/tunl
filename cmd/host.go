package cmd

import (
	"fmt"
	"net"
	"net/url"
	"strconv"

	"github.com/urfave/cli/v2"
)

func DialHost(ctx *cli.Context) (net.Conn, string, error) {
	host := ctx.String("host")
	if len(host) == 0 {
		fmt.Println("Host cannot be empty, see --host flag for more information.\n")

		cli.ShowCommandHelpAndExit(ctx, ctx.Command.Name, 1)
		return nil, "", cli.Exit("Host cannot be empty.", 1)
	}

	hostURL, err := url.Parse(host)
	if err != nil {
		fmt.Println("Host value invalid: %v, see --host flag for more information.\n", err)

		cli.ShowCommandHelpAndExit(ctx, ctx.Command.Name, 1)
	}

	port, err := net.LookupPort("tcp", hostURL.Scheme)
	if err != nil {
		fmt.Println("Host scheme value is invalid: %v, see --host flag for more information.\n", err)

		cli.ShowCommandHelpAndExit(ctx, ctx.Command.Name, 1)
	}

	if specifiedPort := hostURL.Port(); len(specifiedPort) > 0 {
		port, _ = strconv.Atoi(specifiedPort)
	}

	hostnameWithoutPort := hostURL.Hostname()
	if len(hostnameWithoutPort) == 0 {
		fmt.Println("Host hostname cannot be empty: %v, see --host flag for more information.\n", err)

		cli.ShowCommandHelpAndExit(ctx, ctx.Command.Name, 1)
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%v", hostnameWithoutPort, port))
	if err != nil {
		return nil, "", cli.Exit(fmt.Sprintf("Failed to connect to host %s: %v", host, err), 128)
	}

	return conn, hostnameWithoutPort, nil
}
