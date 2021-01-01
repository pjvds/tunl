package cmd

import (
	"crypto/tls"
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

	var conn net.Conn
	server := fmt.Sprintf("%s:%v", hostnameWithoutPort, port)

	if hostURL.Scheme == "https" {
		tlsConn, err := tls.Dial("tcp", server, nil)
		if err != nil {
			return nil, "", cli.Exit(fmt.Sprintf("Failed to connect to host %s: %v", server, err), 128)
		}
		conn = tlsConn
	} else {
		nonTlsConn, err := net.Dial("tcp", server)
		if err != nil {
			return nil, "", cli.Exit(fmt.Sprintf("Failed to connect to host %s: %v", server, err), 128)
		}
		conn = nonTlsConn
	}

	return conn, hostnameWithoutPort, nil
}
