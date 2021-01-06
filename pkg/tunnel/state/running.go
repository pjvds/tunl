package state

import (
	"context"
	"net"

	"github.com/hashicorp/yamux"
	"go.uber.org/zap"
)

type RunningCreator interface {
	Running(conn net.Conn) (*Running, error)
}

type Running struct {
	Conn     net.Conn
	Accepted chan net.Conn

	DisconnectedCreator
}

func (s *Running) Run(ctx context.Context, log *zap.Logger) (State, error) {
	defer s.Conn.Close()
	session, err := yamux.Client(s.Conn, nil)
	if err != nil {
		log.Debug("mux creation error", zap.Error(err))
		return s.Disconnected(err)
	}
	defer session.Close()

	for {
		stream, err := session.Accept()
		if err != nil {
			log.Debug("session accept error", zap.Error(err))
			return s.Disconnected(err)
		}

		log.Debug("session connection accepted", zap.String("remote-addr", stream.RemoteAddr().String()))

		select {
		case s.Accepted <- stream:
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

func (s *Running) String() string {
	return "connected"
}
