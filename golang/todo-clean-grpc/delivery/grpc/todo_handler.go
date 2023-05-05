package grpc

import (
	"context"
	"todo-clean-grpc/domain"
	todopb "todo-clean-grpc/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type todoHandler struct {
	todopb.UnimplementedTodoServiceServer
	TUsecase domain.TodoUsecase
}

func NewTodoGrpcserver(gserver *grpc.Server, tu domain.TodoUsecase) {
	todoServer := &todoHandler{
		TUsecase: tu,
	}
	todopb.RegisterTodoServiceServer(gserver, todoServer)
	reflection.Register(gserver)
}

func (th *todoHandler) Create(ctx context.Context, t *todopb.TodoRequest) (*todopb.TodoResponse, error) {
	todo := domain.Todo{
		Title:  t.GetTitle(),
		Status: t.GetStatus(),
	}
	if err := th.TUsecase.Create(ctx, &todo); err != nil {
		return nil, err
	}
	return &todopb.TodoResponse{
		Title:  todo.Title,
		Status: todo.Status,
	}, nil
}
func (th *todoHandler) Get(ctx context.Context, t *todopb.TodoRequest) (*todopb.TodoResponse, error) {
	todo, _ := th.TUsecase.Get(ctx, t.GetTitle())
	return &todopb.TodoResponse{
		Title:  todo.Title,
		Status: todo.Status,
	}, nil
}
