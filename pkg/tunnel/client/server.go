package client

import (
	"net"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

type ServerInfo struct {
	Host            *url.URL
	ControlHostname string
	Dialer          Dialer
}

func ParseHostURL(host *url.URL) (ServerInfo, error) {
	port := host.Port()
	if len(port) == 0 {
		defaultPort, err := net.LookupPort("tcp", host.Scheme)
		if err != nil {
			return ServerInfo{}, err
		}

		port = strconv.Itoa(defaultPort)
	}

	address := net.JoinHostPort(host.Hostname(), port)

	var dialer Dialer
	switch host.Scheme {
	case "http":
		dialer = TcpDialer(address)
	case "https":
		dialer = TlsDialer(address)
	default:
		return ServerInfo{}, errors.Errorf("scheme %v not valid, use http or https", host.Scheme)
	}

	return ServerInfo{
		Host:            host,
		ControlHostname: host.Hostname(),
		Dialer:          dialer,
	}, nil
}
