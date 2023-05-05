package main

import (
	"log"
	"time"
	_todoHttpHandler "todo-clean-grpc/delivery/http"
	_todoRepository "todo-clean-grpc/repository/sqlite"
	_todoUsecase "todo-clean-grpc/usecase"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbCon, err := sqlx.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		panic(err)
	}

	defer func() {
		err := dbCon.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	migrate(dbCon)

	e := echo.New()
	tRepo := _todoRepository.NewSqliteTodoRepository(dbCon)
	timeoutContext := time.Duration(10) * time.Second
	tUsecase := _todoUsecase.NewTodoUsecase(tRepo, timeoutContext)

	_todoHttpHandler.NewTodoHandler(e, tUsecase)
	log.Fatal(e.Start(":8080"))
}

func migrate(dbCon *sqlx.DB) {
	schema := `CREATE TABLE todo (
		title string,
		status bool);`
	dbCon.Exec(schema)
}
