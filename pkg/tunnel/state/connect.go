package state

import (
	"context"

	"github.com/pjvds/tunl/pkg/tunnel/client"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ConnectCreator interface {
	Connect() (*Connect, error)
}

type Connect struct {
	Server client.ServerInfo
	Type   client.TunnelType

	client.TunnelInfoSetter
	RunningCreator
}

func (s *Connect) Run(ctx context.Context, log *zap.Logger) (State, error) {
	conn, err := s.Server.Dialer.Dial()
	if err != nil {
		return nil, errors.Wrap(err, "server dial failed")
	}

	info, err := client.Handshake(conn, s.Server.ControlHostname, nil, s.Type)
	if err != nil {
		return nil, errors.Wrap(err, "server handshake failed")
	}

	s.SetTunnelInfo(info)

	return s.Running(conn)
}

func (s *Connect) String() string {
	return "connecting"
}
