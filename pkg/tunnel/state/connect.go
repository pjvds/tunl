package state

import (
	"context"

	"github.com/pjvds/tunl/pkg/tunnel/client"
	"go.uber.org/zap"
)

type ConnectCreator interface {
	Connect() (*Connect, error)
}

type Connect struct {
	Server client.ServerInfo

	client.TunnelInfoSetter
	RunningCreator
}

func (s *Connect) Run(ctx context.Context, log *zap.Logger) (State, error) {
	conn, err := s.Server.Dialer.Dial()
	if err != nil {
		return nil, err
	}

	info, err := client.Handshake(conn, s.Server.ControlHostname, nil)
	if err != nil {
		return nil, err
	}

	s.SetTunnelInfo(info)

	return s.Running(conn)
}

func (s *Connect) String() string {
	return "connecting"
}
