package sqlite

import (
	"context"
	"time"
	"todo-clean2/domain"

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

func (s *sqliteTodoRepository) Get(ctx context.Context, id int) (domain.Todo, error) {
	todo := domain.Todo{}
	if err := s.Conn.GetContext(ctx, &todo, "SELECT * FROM WHERE id = ?", id); err != nil {
		return domain.Todo{}, err
	}
	return todo, nil
}

func (s *sqliteTodoRepository) GetList(ctx context.Context) ([]domain.Todo, error) {
	todos := []domain.Todo{}
	if err := s.Conn.SelectContext(ctx, &todos, "SELECT * FROM todo"); err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *sqliteTodoRepository) Create(ctx context.Context, t *domain.Todo) error {
	_, err := s.Conn.ExecContext(ctx, "INSERT INTO todo (titile, status, created_at) VALUES (?, ?, ?, ?)",
		t.Title,
		t.Status,
		time.Now().Format(time.DateOnly),
	)
	if err != nil {
		return err
	}
	return nil
}
