package client

import "net/http"

type TunnelType interface {
	AddHeaders(req *http.Request) error
}

var TypeHTTP = typeHeaderValue("http")
var TypeTCP = typeHeaderValue("tcp")
var TypeTLS = typeHeaderValue("tls")

type typeHeaderValue string

func (t typeHeaderValue) AddHeaders(req *http.Request) error {
	req.Header.Add("X-Tunl-Type", string(t))
	return nil
}
