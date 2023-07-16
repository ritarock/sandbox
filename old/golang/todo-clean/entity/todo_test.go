package entity_test

import (
	"bytes"
	"testing"
	"todo-clean/entity"

	"github.com/stretchr/testify/assert"
)

func TestNewTodo(t *testing.T) {
	todo, err := entity.NewTodo("test title", false)
	assert.Nil(t, err)
	assert.Equal(t, todo.Title, "test title")
	assert.NotNil(t, todo.ID)
}

func TestTodoValidate(t *testing.T) {
	type test struct {
		title  string
		status bool
		want   error
	}

	tests := []test{
		{
			title:  string(bytes.Repeat([]byte("t"), 100)),
			status: false,
			want:   nil,
		},
		{
			title:  string(bytes.Repeat([]byte("t"), 101)),
			status: false,
			want:   entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := entity.NewTodo(tc.title, tc.status)
		assert.Equal(t, err, tc.want)
	}
}
