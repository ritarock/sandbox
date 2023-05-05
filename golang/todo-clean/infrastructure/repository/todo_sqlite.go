package repository

import (
	"time"
	"todo-clean/entity"

	"github.com/jmoiron/sqlx"
)

type TodoSqlite struct {
	db *sqlx.DB
}

func NewTodoSqlite(db *sqlx.DB) *TodoSqlite {
	return &TodoSqlite{
		db: db,
	}
}

func (r *TodoSqlite) Get(id *entity.ID) (*entity.Todo, error) {
	todo := entity.Todo{}
	if err := r.db.Get(&todo, "SELECT * FROM todo WHERE id = ?", id); err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoSqlite) List() ([]*entity.Todo, error) {
	todos := []*entity.Todo{}
	if err := r.db.Select(&todos, "SELECT * FROM todo"); err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoSqlite) Create(e *entity.Todo) (entity.ID, error) {
	_, err := r.db.Exec("INSERT INTO todo (id, title, status, created_at) VALUES (?, ?, ?, ?)",
		e.ID,
		e.Title,
		e.Status,
		time.Now().Format(time.DateOnly),
	)
	if err != nil {
		return e.ID, err
	}
	return e.ID, nil
}
