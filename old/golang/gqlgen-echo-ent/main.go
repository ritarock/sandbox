package main

import (
	"context"
	"gqlgen-echo-ent/ent"
	"gqlgen-echo-ent/graph"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

const (
	DRIVER      = "sqlite3"
	DATA_SOURCE = "file:ent?mode=memory&cache=shared&_fk=1"
)

func main() {
	ctx := context.Background()
	client, _ := ent.Open(DRIVER, DATA_SOURCE)
	client.Schema.Create(ctx)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})

	graphqlHandler := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: &graph.Resolver{EntCLient: client}},
		),
	)
	playgroundHandler := playground.Handler("GraphQL", "/query")

	e.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.Logger.Fatal(e.Start(":8000"))
}
