package repository

import (
	"context"
	"todo-clean-grpc/domain"

	"github.com/jmoiron/sqlx"
)

type sqliteTodoRepository struct {
	Conn *sqlx.DB
}

func NewSqliteTodoRepository(conn *sqlx.DB) domain.TodoRepository {
	return &sqliteTodoRepository{
		Conn: conn,
	}
}

func (s *sqliteTodoRepository) Create(ctx context.Context, t *domain.Todo) error {
	_, err := s.Conn.ExecContext(ctx,
		"INSERT INTO todo (title, status) VALUES (?, ?)",
		t.Title,
		t.Status,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *sqliteTodoRepository) Get(ctx context.Context, title string) (domain.Todo, error) {
	todo := domain.Todo{}
	if err := s.Conn.GetContext(ctx, &todo, "SELECT * FROM todo WHERE title = ?", title); err != nil {
		return todo, err
	}

	return todo, nil
}
