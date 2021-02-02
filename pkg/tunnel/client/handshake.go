package client

import (
	"bufio"
	"net"
	"net/http"

	"github.com/Masterminds/semver"
	"github.com/pkg/errors"
)

type TunnelInfo struct {
	Id      string
	Token   string
	Address string
	Version *semver.Version
}

type TunnelInfoSetter interface {
	SetTunnelInfo(info TunnelInfo)
}

func Handshake(conn net.Conn, hostname string, tunnel *TunnelInfo, t TunnelType) (TunnelInfo, error) {
	url := "/"
	if tunnel != nil && len(tunnel.Id) > 0 {
		url += "?id=" + tunnel.Id
	}

	request, _ := http.NewRequest(http.MethodConnect, url, nil)
	t.AddHeaders(request)

	// this hostname is used to redirect this request to the control mux (e.q.: _.tunl.es)
	request.Host = hostname

	if tunnel != nil && len(tunnel.Token) > 0 {
		request.Header.Add("X-Tunl-Token", tunnel.Token)
	}

	if err := request.Write(conn); err != nil {
		return TunnelInfo{}, err
	}

	reader := bufio.NewReader(conn)
	response, err := http.ReadResponse(reader, request)
	if err != nil {
		return TunnelInfo{}, err
	}

	if response.StatusCode != http.StatusOK {
		return TunnelInfo{}, errors.Errorf("unexpected reponsee status %v", response.Status)
	}

	id := response.Header.Get("X-Tunl-Id")
	token := response.Header.Get("X-Tunl-Token")
	address := response.Header.Get("X-Tunl-Address")
	version, _ := semver.NewVersion("X-Tunl-Version")

	return TunnelInfo{
		Id:      id,
		Token:   token,
		Address: address,
		Version: version,
	}, nil
}
