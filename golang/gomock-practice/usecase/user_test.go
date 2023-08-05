package usecase

import (
	"testing"

	"gomock-practice/entity"
	mockRepo "gomock-practice/repository/mock"

	"go.uber.org/mock/gomock"
)

func TestUserUsecase_DoSomething(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock の作成
	mockRepo := mockRepo.NewMockUserRepository(ctrl)
	usecase := NewUseCase(mockRepo)

	id := 1
	name := "mock user"
	mockEntity := &entity.User{ID: id, Name: name}
	mockRepo.EXPECT().GetByID(id).Return(mockEntity, nil).Times(1)

	entity, err := usecase.DoSomething(id)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if entity == nil {
		t.Errorf("entity is nil")
	}
	if entity.ID != id {
		t.Errorf("expected: %v, actuacl: %v", id, entity.ID)
	}
	if entity.Name != name {
		t.Errorf("expected: %v, actuacl: %v", name, entity.Name)
	}
}
