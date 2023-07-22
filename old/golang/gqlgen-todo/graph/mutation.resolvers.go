package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"gqlgen-todo/graph/model"
	"strconv"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	created, _ := r.EntClient.Todo.Create().
		SetTitle(input.Title).
		SetNote(input.Note).
		Save(ctx)

	return &model.Todo{
		ID:        strconv.Itoa(created.ID),
		Title:     created.Title,
		Note:      created.Note,
		Completed: created.Completed,
		CreatedAt: created.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: created.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }