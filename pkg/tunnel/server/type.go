package server

type ClientInfo struct {
	Hostname string
	Address  string
}

type TunnelEndpoint interface {
	Id() string
	Serve() error
}
