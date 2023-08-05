package usecase

import (
	"gomock-practice/entity"
	"gomock-practice/repository"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUseCase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (uc *UserUsecase) DoSomething(id int) (*entity.User, error) {
	user, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
