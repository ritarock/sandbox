package ent

import (
	"context"
	"todo-clean-arch/domain"
	"todo-clean-arch/ent"
	"todo-clean-arch/ent/todo"
)

type entTodoRepository struct {
	client *ent.Client
}

func NewEntTodoRepository(client *ent.Client) domain.TodoRepository {
	return &entTodoRepository{client}
}

func (e *entTodoRepository) convert(todo ent.Todo) *domain.Todo {
	return &domain.Todo{
		ID:        todo.ID,
		Title:     todo.Title,
		Status:    todo.Status,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
}

func (e *entTodoRepository) Create(ctx context.Context, todo *domain.Todo) (err error) {
	created, err := e.client.
		Todo.
		Create().
		SetTitle(todo.Title).
		SetStatus(todo.Status).
		Save(ctx)
	if err != nil {
		return
	}

	todo.ID = created.ID
	return
}

func (e *entTodoRepository) GetById(ctx context.Context, id int) (domain.Todo, error) {
	searched, err := e.client.
		Todo.
		Query().
		Where(todo.ID(id)).
		First(ctx)
	if err != nil {
		return domain.Todo{}, err
	}
	result := e.convert(*searched)
	return *result, nil
}

func (e *entTodoRepository) GetByTitle(ctx context.Context, title string) (domain.Todo, error) {
	searched, err := e.client.
		Todo.
		Query().
		Where(todo.Title(title)).
		First(ctx)
	if err != nil {
		return domain.Todo{}, err
	}
	result := e.convert(*searched)
	return *result, nil
}

func (e *entTodoRepository) Update(ctx context.Context, todo *domain.Todo) error {
	_, err := e.client.
		Todo.
		UpdateOneID(todo.ID).
		SetTitle(todo.Title).
		SetStatus(todo.Status).
		SetUpdatedAt(todo.UpdatedAt).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (e *entTodoRepository) Delete(ctx context.Context, id int) error {
	if err := e.client.Todo.DeleteOneID(id).Exec(ctx); err != nil {
		return err
	}
	return nil
}
