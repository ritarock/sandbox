package domain

import (
	"context"
	"time"
)

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type TodoUsecase interface {
	Get(ctx context.Context, id int) (Todo, error)
	GetList(ctx context.Context) ([]Todo, error)
	Create(ctx context.Context, t *Todo) error
}

type TodoRepository interface {
	Get(ctx context.Context, id int) (Todo, error)
	GetList(ctx context.Context) ([]Todo, error)
	Create(ctx context.Context, t *Todo) error
}
