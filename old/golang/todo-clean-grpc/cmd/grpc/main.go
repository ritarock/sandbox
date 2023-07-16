package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
	_deliveryGprc "todo-clean-grpc/delivery/grpc"
	_todoRepository "todo-clean-grpc/repository/sqlite"
	_todoUsecase "todo-clean-grpc/usecase"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
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

	s := grpc.NewServer()
	tRepo := _todoRepository.NewSqliteTodoRepository(dbCon)
	timeoutContext := time.Duration(10) * time.Second
	tUsecase := _todoUsecase.NewTodoUsecase(tRepo, timeoutContext)
	_deliveryGprc.NewTodoGrpcserver(s, tUsecase)

	port := 8000
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	s.GracefulStop()
}

func migrate(dbCon *sqlx.DB) {
	schema := `CREATE TABLE todo (
		title string,
		status bool);`
	dbCon.Exec(schema)
}
