package repository

import "context"

type SemaphoreRepository interface {
	Acquire(ctx context.Context) error
	Release()
}
