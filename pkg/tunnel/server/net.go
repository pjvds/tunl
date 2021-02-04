package server

import "net"

type ListenInterceptor struct {
	Interceptor func(c net.Conn) (net.Conn, error)

	net.Listener
}

// Accept waits for and returns the next connection to the listener.
func (l *ListenInterceptor) Accept() (net.Conn, error) {
	c, err := l.Listener.Accept()
	if err != nil {
		return nil, err
	}

	return l.Interceptor(c)
}
