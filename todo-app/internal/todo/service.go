package todo

// this is the actual service provided

import (
	"context"
	"errors"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateTodo(ctx context.Context, title string) (Todo, error) {
	if title == "" {
		return Todo{}, errors.New("Invalid title")
	}
	return s.repo.Create(ctx, title)
}

func (s *Service) ListTodos(ctx context.Context) ([]Todo, error) {
	return s.repo.GetAll(ctx)
}
