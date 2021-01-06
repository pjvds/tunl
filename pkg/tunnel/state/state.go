package state

import (
	"context"

	"go.uber.org/zap"
)

type State interface {
	Run(ctx context.Context, log *zap.Logger) (State, error)
	String() string
}
