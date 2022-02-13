package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sample-echo-ent-edges/ent"
	"sample-echo-ent-edges/ent/todo"
	"sample-echo-ent-edges/ent/user"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

type User struct {
	Name string `json:"name"`
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

	e.GET("/todos", func(c echo.Context) error {
		todos, err := client.Todo.
			Query().
			All(context.Background())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": http.StatusInternalServerError,
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"todos": todos,
		})
	})

	e.POST("/todo", func(c echo.Context) error {
		t := &Todo{}
		if err := c.Bind(t); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": http.StatusInternalServerError,
			})
		}
		todo, err := client.Todo.
			Create().
			SetName(t.Name).
			SetStatus(t.Status).
			SetCreatedAt(time.Now()).
			Save(context.Background())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": http.StatusInternalServerError,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"name":   todo.Name,
			"status": todo.Status,
		})
	})

	e.GET("/todo/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		todo, err := client.Todo.
			Query().
			Where(todo.ID(id)).
			Only(context.Background())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": http.StatusInternalServerError,
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"name":       todo.Name,
			"status":     todo.Status,
			"created_at": todo.CreatedAt,
		})
	})

	e.PUT("/todo/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		t := &Todo{}
		if err := c.Bind(t); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": http.StatusInternalServerError,
			})
		}

		todo, err := client.Todo.
			UpdateOneID(id).
			SetName(t.Name).
			SetStatus(t.Status).
			Save(context.Background())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": http.StatusInternalServerError,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
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
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": http.StatusInternalServerError,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": fmt.Sprintf("DELETE: %v", id),
		})
	})

	e.GET("/users", func(c echo.Context) error {
		users, err := client.User.
			Query().
			All(context.Background())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": http.StatusInternalServerError,
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"users": users,
		})
	})

	e.POST("/user", func(c echo.Context) error {
		u := &User{}
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": http.StatusInternalServerError,
			})
		}
		user, err := client.User.
			Create().
			SetName(u.Name).
			Save(context.Background())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": http.StatusInternalServerError,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"name": user.Name,
		})
	})

	e.GET("user/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		user, err := client.User.
			Query().
			Where(user.ID(id)).
			Only(context.Background())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": http.StatusInternalServerError,
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"name": user.Name,
		})
	})

	e.GET("add_todo/:todo_id/:user_id", func(c echo.Context) error {
		todoId, _ := strconv.Atoi(c.Param("todo_id"))
		userId, _ := strconv.Atoi(c.Param("user_id"))

		user, err := client.User.
			UpdateOneID(userId).
			AddTodoIDs(todoId).
			Save(context.Background())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": http.StatusInternalServerError,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"name": user.Name,
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
