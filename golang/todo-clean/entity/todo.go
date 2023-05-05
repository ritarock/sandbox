package entity

import (
	"time"
	"unicode/utf8"
)

type Todo struct {
	ID        ID
	Title     string
	Status    bool
	CreatedAt time.Time
}

func NewTodo(title string, status bool) (*Todo, error) {
	t := &Todo{
		ID:        NewID(),
		Title:     title,
		Status:    status,
		CreatedAt: time.Now(),
	}
	if err := t.Validate(); err != nil {
		return nil, ErrInvalidEntity
	}
	return t, nil
}

func (t *Todo) Validate() error {
	if utf8.RuneCountInString(t.Title) > 100 {
		return ErrInvalidEntity
	}
	return nil
}
