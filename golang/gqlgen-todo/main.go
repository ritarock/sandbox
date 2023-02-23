package main

import (
	"context"
	"gqlgen-todo/ent"
	"gqlgen-todo/graph"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
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
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	graphqlHandler := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{
					EntClient: client,
				},
			},
		),
	)
	playgroudHandler := playground.Handler("GraphQL", "/query")

	e.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("playground", func(c echo.Context) error {
		playgroudHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.Logger.Fatal(e.Start(":8080"))
}
