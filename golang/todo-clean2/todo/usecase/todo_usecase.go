package usecase

import (
	"context"
	"time"
	"todo-clean2/domain"
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

func (tu *todoUsecase) Get(c context.Context, id int) (domain.Todo, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	res, err := tu.todoRepo.Get(ctx, id)
	if err != nil {
		return domain.Todo{}, err
	}
	return res, nil
}

func (tu *todoUsecase) GetList(c context.Context) ([]domain.Todo, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	res, err := tu.todoRepo.GetList(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (tu *todoUsecase) Create(c context.Context, t *domain.Todo) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	if err := tu.todoRepo.Create(ctx, t); err != nil {
		return err
	}
	return nil
}
