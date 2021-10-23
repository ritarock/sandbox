package model

import (
	"backend/internal/server/types"
	"backend/prisma/db"
	"context"
)

func GetComments(taskId string) ([]db.CommentModel, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	comments, err := client.Comment.FindMany(
		db.Comment.TaskID.Equals(taskId),
	).Exec(context.Background())

	return comments, err
}

func CreateComment(taskId string, newComment types.Comment) (*db.CommentModel, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	created, err := client.Comment.CreateOne(
		db.Comment.Content.Set(newComment.Content),
		db.Comment.Task.Link(
			db.Task.ID.Equals(taskId),
		),
	).Exec(context.Background())

	return created, err
}

func GetComment(commentId string) (*db.CommentModel, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	comment, err := client.Comment.FindUnique(
		db.Comment.ID.Equals(commentId),
	).Exec(context.Background())

	return comment, err
}

func UpdateComment(commentId string, newComment types.Comment) (*db.CommentModel, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	updated, err := client.Comment.FindUnique(
		db.Comment.ID.Equals(commentId),
	).Update(
		db.Comment.Content.Set(newComment.Content),
	).Exec(context.Background())

	return updated, err
}

func DeleteComment(commentId string) error {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	_, err := client.Comment.FindUnique(
		db.Comment.ID.Equals(commentId),
	).Delete().Exec(context.Background())

	return err
}
