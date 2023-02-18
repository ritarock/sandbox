package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"gqlgen-echo-ent/ent/task"
	"gqlgen-echo-ent/graph/model"
	"time"
)

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	timeStamp := time.Now().Format("2006-01-02 15:04:05")
	task := model.Task{
		Title:     input.Title,
		Note:      input.Note,
		Completed: 0,
		CreatedAt: timeStamp,
		UpdatedAt: timeStamp,
	}
	r.EntCLient.Task.Create().
		SetTitle(input.Title).
		SetNote(input.Note).
		SetCompleted(0).Save(ctx)

	return &task, nil
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	tasks := []*model.Task{}
	t, _ := r.EntCLient.Task.Query().Select(
		task.FieldTitle,
		task.FieldNote,
		task.FieldCompleted,
	).All(ctx)
	for _, task := range t {
		tasks = append(tasks, &model.Task{
			Title:     task.Title,
			Note:      task.Note,
			Completed: task.Completed,
		})
	}
	return tasks, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
