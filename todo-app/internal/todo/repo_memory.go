package todo

import (
	"context"
	"sync"
	"time"
)

type MemoryRepo struct {
	mu    sync.Mutex
	todos []Todo
	next  int64
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{next: 1}
}

func (r *MemoryRepo) Create(ctx context.Context, title string) (Todo, error) {
	r.mu.Lock()

	defer r.mu.Unlock()

	todo := Todo{
		Title:     title,
		CreatedAt: time.Now(),
		Completed: false,
		ID:        r.next,
	}

	r.next++

	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *MemoryRepo) GetAll(ctx context.Context) ([]Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]Todo{}, r.todos...), nil

}
