package todo

// contains all our types and objects

import "context"

type Repository interface {
	Create(ctx context.Context, title string) (Todo, error)
	GetAll(ctx context.Context) ([]Todo, error)
}
