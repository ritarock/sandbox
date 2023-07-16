package todo

import "todo-clean/entity"

type Reader interface {
	Get(id entity.ID) (*entity.Todo, error)
	List() ([]*entity.Todo, error)
}

type Writer interface {
	Create(e *entity.Todo) (entity.ID, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetTodo(id entity.ID) (*entity.Todo, error)
	GetTodoList() ([]*entity.Todo, error)
	CreateTodo(title string, status bool) (entity.ID, error)
}
