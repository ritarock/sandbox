package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sample_prisma/prisma/db"

	"github.com/gin-gonic/gin"
)

func main() {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		fmt.Println(err)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			log.Fatal(err)
		}
	}()

	r := gin.Default()

	r.POST("/tasks", func(c *gin.Context) {
		var task db.TaskModel
		if err := c.ShouldBind(&task); err != nil {
			c.String(http.StatusOK, `Bind err`)
		}
		var text *string
		if newText, ok := task.Text(); ok {
			fmt.Println(newText)
			text = &newText
		}
		var completed *bool
		if newCompleted, ok := task.Completed(); ok {
			completed = &newCompleted
		}
		newTask, err := client.Task.CreateOne(
			db.Task.Text.SetIfPresent(text),
			db.Task.Completed.SetIfPresent(completed),
		).Exec(context.Background())
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{"done": newTask})
	})

	r.GET("/tasks", func(c *gin.Context) {
		tasks, err := client.Task.FindMany().OrderBy(
			db.Task.ID.Order(db.ASC),
		).Exec(context.Background())
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(tasks)
		c.JSON(http.StatusOK, tasks)
	})

	r.Run()
}
