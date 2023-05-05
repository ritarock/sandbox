package usecase

import (
	"context"
	"time"
	"todo-clean-grpc/domain"
)

type todoUsecase struct {
	todoRepo       domain.TodoRepository
	contextTimeout time.Duration
}

func NewTodoUsecase(t domain.TodoRepository, timeout time.Duration) domain.TodoUsecase {
	return &todoUsecase{
		todoRepo:       t,
		contextTimeout: timeout,
	}
}

func (tu *todoUsecase) Create(c context.Context, t *domain.Todo) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	if err := tu.todoRepo.Create(ctx, t); err != nil {
		return err
	}
	return nil
}

func (tu *todoUsecase) Get(c context.Context, title string) (domain.Todo, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	res, err := tu.todoRepo.Get(ctx, title)
	if err != nil {
		return domain.Todo{}, nil
	}

	return res, nil
}
