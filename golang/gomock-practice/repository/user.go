package repository

import "gomock-practice/entity"

type UserRepository interface {
	GetByID(id int) (*entity.User, error)
	Save(entity.User) error
}

type DatabaseRepository struct {
}

func (r *DatabaseRepository) GetByID(id int) (*entity.User, error) {
	return &entity.User{ID: id, Name: "sample"}, nil
}

func (r *DatabaseRepository) Save(user entity.User) error {
	return nil
}
