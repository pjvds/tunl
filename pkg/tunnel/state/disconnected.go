package state

import (
	"context"

	"go.uber.org/zap"
)

type DisconnectedCreator interface {
	Disconnected(err error) (*Disconnected, error)
}

type Disconnected struct {
	LastErr error
	ReconnectCreator
}

func (s *Disconnected) Run(ctx context.Context, log *zap.Logger) (State, error) {
	return s.Reconnect(1)
}

func (s *Disconnected) String() string {
	return "disconnected"
}
