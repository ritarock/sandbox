## ent
### インストール
ent と DB の client をインストールする。
```bash
$ go get -d entgo.io/ent/cmd/ent
$ go get github.com/mattn/go-sqlite3
```

### スキーマの作成
```bash
$ go run entgo.io/ent/cmd/ent init Todo
```

`ent/schema/todo.go` を編集する。
```go
// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default(""),
		field.Bool("status").
			Default(false),
		field.Time("created_at").
			Default(func() time.Time {
				return time.Now()
			}),
	}
}
```

### コード生成
shema ファイルを元にコードを生成するので編集後は実行する必要がある。
```bash
$ go generate ./ent
```

### マイグレーション部分を実装
```go
func main() {

	client, err := ent.Open("sqlite3", "file:todo.sqlite?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("faild opening connection to sqlit")
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("faild creating schema")
	}
}
```

## echo
`Todo` の struct を定義。
```go
type Todo struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}
```

今回は CRUD 可能な Todo アプリを作りたいので以下のハンドラを作成する。
```go
e := echo.New()

e.POST("/todo", func(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
})
e.GET("/todo/:id", func(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
})
e.PUT("/todo/:id", func(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
})
e.DELETE("/todo/:id", func(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
})

e.Logger.Fatal(e.Start(":8080"))
```

## create を実装
```go
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
```

実行。
```bash
$ curl --location --request POST 'http://localhost:8080/todo' \
       --header 'Content-Type: application/json' \
       --data-raw '{"name": "todo1"}'
{"name":"todo1","status":"false"}
```

## READ を実装
```go
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
```

実行。
```bash
$ curl http://localhost:8080/todo/1
{"created_at":"2022-01-09 15:30:02.059031 +0900 +0900","name":"todo1","status":"false"}
```

## UPDATE を実装
```go
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
```

```bash
$ curl --location --request PUT 'http://localhost:8080/todo/1' \
       --header 'Content-Type: application/json' \
       --data-raw '{"name": "todo1", "status": true}'
{"created_at":"2022-01-09 15:30:02.059031 +0900 +0900","name":"todo1","status":"true"}
```

## DELETE を実装
```go
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
```

実行
```bash
curl --location --request DELETE 'http://localhost:8080/todo/1'
{"message":"DELETE: 1"}
```

### コード全体
```go
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
```
