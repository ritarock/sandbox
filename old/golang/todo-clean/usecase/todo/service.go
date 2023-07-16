package todo

import "todo-clean/entity"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetTodo(id entity.ID) (*entity.Todo, error) {
	t, err := s.repo.Get(id)
	if t == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (s *Service) GetTodoList() ([]*entity.Todo, error) {
	todos, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(todos) == 0 {
		return nil, entity.ErrNotFound
	}
	return todos, nil
}

func (s *Service) CreateTodo(title string, status bool) (entity.ID, error) {
	t, err := entity.NewTodo(title, status)
	if err != nil {
		return entity.NewID(), err
	}
	return s.repo.Create(t)
}
