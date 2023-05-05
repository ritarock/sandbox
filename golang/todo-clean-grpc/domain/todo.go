package domain

import (
	"context"
)

type Todo struct {
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

type TodoUsecase interface {
	Create(ctx context.Context, t *Todo) error
	Get(ctx context.Context, title string) (Todo, error)
}

type TodoRepository interface {
	Create(ctx context.Context, t *Todo) error
	Get(ctx context.Context, title string) (Todo, error)
}
