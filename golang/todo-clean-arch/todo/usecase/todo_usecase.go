package usecase

import (
	"context"
	"time"
	"todo-clean-arch/domain"
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

func (t *todoUsecase) Create(c context.Context, todo *domain.Todo) error {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()
	if err := t.todoRepo.Create(ctx, todo); err != nil {
		return err
	}
	return nil
}

func (t *todoUsecase) GetById(c context.Context, id int) (res domain.Todo, err error) {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()

	res, err = t.todoRepo.GetById(ctx, id)
	if err != nil {
		return
	}

	return
}

func (t *todoUsecase) GetByTitle(c context.Context, title string) (res domain.Todo, err error) {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()

	res, err = t.todoRepo.GetByTitle(ctx, title)
	if err != nil {
		return
	}

	return
}

func (t *todoUsecase) Update(c context.Context, todo *domain.Todo) error {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()

	if err := t.todoRepo.Update(ctx, todo); err != nil {
		return err
	}

	return nil
}

func (t *todoUsecase) Delete(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()

	if err := t.todoRepo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
