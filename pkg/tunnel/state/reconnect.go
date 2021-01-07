package state

import (
	"context"
	"time"

	"github.com/pjvds/tunl/pkg/tunnel/client"
	"go.uber.org/zap"
)

type ReconnectCreator interface {
	Reconnect(attempt int) (*Reconnect, error)
}

type Reconnect struct {
	Server  client.ServerInfo
	Type    client.TunnelType
	Token   client.TokenRefresher
	Tunnel  client.TunnelInfo
	Attempt int

	ReconnectCreator
	RunningCreator
}

func (s *Reconnect) Run(ctx context.Context, log *zap.Logger) (State, error) {
	conn, err := s.Server.Dialer.Dial()
	if err != nil {
		time.Sleep(5 * time.Second)
		return s.Reconnect(s.Attempt + 1)
	}

	info, err := client.Handshake(conn, s.Server.ControlHostname, &s.Tunnel, s.Type)
	if err != nil {
		time.Sleep(5 * time.Second)
		return s.Reconnect(s.Attempt + 1)
	}

	// Set token and enter running state
	s.Token.SetToken(info.Token)

	return s.Running(conn)
}

func (s *Reconnect) String() string {
	if s.Attempt > 1 {
		return ""
	}

	return "connecting"
}
