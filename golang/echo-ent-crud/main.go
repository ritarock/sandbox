package main

import (
	"context"
	"echo-ent-crud/ent"
	"echo-ent-crud/ent/todo"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

func main() {
	client, err := ent.Open("sqlite3", "file:todo.sqlite?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("faild opening connection to sqlit")
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("faild creating schema")
	}

	e := echo.New()

	e.POST("/todo", func(c echo.Context) error {
		t := &Todo{}
		if err := c.Bind(t); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": http.StatusText(http.StatusInternalServerError),
			})
		}
		todo, err := client.Todo.
			Create().
			SetName(t.Name).
			SetStatus(t.Status).
			SetCreatedAt(time.Now()).
			Save(context.Background())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": http.StatusText(http.StatusInternalServerError),
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"name":   todo.Name,
			"status": strconv.FormatBool(todo.Status),
		})
	})

	e.GET("/todo/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		todo, err := client.Todo.
			Query().
			Where(todo.ID(id)).
			Only(context.Background())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": http.StatusText(http.StatusInternalServerError),
			})
		}
		return c.JSON(http.StatusOK, map[string]string{
			"name":       todo.Name,
			"status":     strconv.FormatBool(todo.Status),
			"created_at": todo.CreatedAt.String(),
		})
	})

	e.PUT("/todo/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		t := &Todo{}
		if err := c.Bind(t); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": http.StatusText(http.StatusInternalServerError),
			})
		}

		todo, err := client.Todo.
			UpdateOneID(id).
			SetName(t.Name).
			SetStatus(t.Status).
			Save(context.Background())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": http.StatusText(http.StatusInternalServerError),
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"name":       todo.Name,
			"status":     strconv.FormatBool(todo.Status),
			"created_at": todo.CreatedAt.String(),
		})
	})

	e.DELETE("/todo/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		err := client.Todo.
			DeleteOneID(id).
			Exec(context.Background())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": http.StatusText(http.StatusInternalServerError),
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": fmt.Sprintf("DELETE: %v", id),
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
