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
	UpdatedAt time.Time `json:"updated_at"`
}

type TodoUsecase interface {
	Create(ctx context.Context, todo *Todo) error
	GetById(ctx context.Context, id int) (Todo, error)
	GetByTitle(ctx context.Context, title string) (Todo, error)
	Update(ctx context.Context, todo *Todo) error
	Delete(ctx context.Context, id int) error
}

type TodoRepository interface {
	Create(ctx context.Context, todo *Todo) error
	GetById(ctx context.Context, id int) (Todo, error)
	GetByTitle(ctx context.Context, title string) (Todo, error)
	Update(ctx context.Context, todo *Todo) error
	Delete(ctx context.Context, id int) error
}
