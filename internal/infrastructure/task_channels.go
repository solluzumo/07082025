package taskInfrastructure

import (
	"context"
	"errors"
)

var maxTaskError = errors.New("Too many tasks exist")

type ChannelTask struct {
	sem chan struct{}
}

func NewChannelTask(maxChann int) *ChannelTask {
	return &ChannelTask{
		sem: make(chan struct{}, maxChann),
	}
}

func (ct *ChannelTask) Acquire(ctx context.Context) error {
	select {
	case ct.sem <- struct{}{}:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	default:
		return maxTaskError
	}
}

func (ct *ChannelTask) Release() {
	<-ct.sem
}
