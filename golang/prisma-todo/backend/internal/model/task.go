package model

import (
	"backend/internal/server/types"
	"backend/prisma/db"
	"context"
)

func GetTasks() ([]db.TaskModel, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	tasks, err := client.Task.FindMany().Exec(context.Background())

	return tasks, err
}

func CreateTask(newTask types.Task) (*db.TaskModel, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	created, err := client.Task.CreateOne(
		db.Task.Title.Set(newTask.Title),
		db.Task.Status.Set(newTask.Status),
		db.Task.Desc.Set(newTask.Desc),
	).Exec(context.Background())
	return created, err
}

func GetTask(taskId string) (*db.TaskModel, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	task, err := client.Task.FindUnique(
		db.Task.ID.Equals(taskId),
	).Exec(context.Background())
	return task, err
}

func UpdateTask(taskId string, newTask types.Task) (*db.TaskModel, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	updated, err := client.Task.FindUnique(
		db.Task.ID.Equals(taskId),
	).Update(
		db.Task.Title.Set(newTask.Title),
		db.Task.Status.Set(newTask.Status),
		db.Task.Desc.Set(newTask.Desc),
	).Exec(context.Background())
	return updated, err
}

func DeleteTask(taskId string) error {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	_, err := client.Task.FindUnique(
		db.Task.ID.Equals(taskId),
	).Delete().Exec(context.Background())
	return err
}
