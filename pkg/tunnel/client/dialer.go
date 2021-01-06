package client

import (
	"crypto/tls"
	"net"
)

type Dialer interface {
	Dial() (net.Conn, error)
}

type TlsDialer string

func (addr TlsDialer) Dial() (net.Conn, error) {
	return tls.Dial("tcp", string(addr), nil)
}

type TcpDialer string

func (addr TcpDialer) Dial() (net.Conn, error) {
	return net.Dial("tcp", string(addr))
}
