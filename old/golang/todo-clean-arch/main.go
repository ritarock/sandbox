package main

import (
	"context"
	"log"
	"time"
	"todo-clean-arch/ent"
	_todoDelivery "todo-clean-arch/todo/delivery/http"
	_todoRepo "todo-clean-arch/todo/repository/ent"
	_todoUsecase "todo-clean-arch/todo/usecase"

	_ "github.com/mattn/go-sqlite3"

	"github.com/labstack/echo/v4"
)

const (
	DRIVER      = "sqlite3"
	DATA_SOURCE = "file:data.sqlite?cache=shared&_fk=1"
)

func main() {
	client, err := ent.Open(DRIVER, DATA_SOURCE)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	todoRepo := _todoRepo.NewEntTodoRepository(client)
	timeoutContext := time.Duration(2) * time.Second
	todoUsecase := _todoUsecase.NewTodoUsecase(todoRepo, timeoutContext)
	_todoDelivery.NewTodoHandler(e, todoUsecase)

	log.Fatal(e.Start(":8080"))
}
