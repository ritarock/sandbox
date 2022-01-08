package main

import (
	"context"
	"fmt"
	"log"
	"sample-ent/ent"
	"sample-ent/ent/todo"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	task1, err := client.Todo.Create().SetText("Add GraphQL Example").Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a todo: %v", err)
	}

	task2, err := client.Todo.Create().SetText("Add Tracing Example").Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a todo: %v", err)
	}

	// task2 の親に task1 をセットする
	if err := task2.Update().SetParent(task1).Exec(ctx); err != nil {
		log.Fatalf("failed connecting todo2 to its parent: %v", err)
	}

	// すべての TODO を取得する
	items, err := client.Todo.Query().All(ctx)
	if err != nil {
		log.Fatalf("failed querying todos: %v", err)
	}
	fmt.Println("// すべての TODO を取得する")
	for _, t := range items {
		fmt.Printf("%d: %q\n", t.ID, t.Text)
	}

	// 親を持つ TODO を取得する
	items, err = client.Todo.Query().Where(todo.HasParent()).All(ctx)
	if err != nil {
		log.Fatalf("failed querying todos: %v", err)
	}
	fmt.Println("// 親を持つ TODO を取得する")
	for _, t := range items {
		fmt.Printf("%d: %q\n", t.ID, t.Text)
	}

	// 他の TODO に依存しておらず、子を持つ TODO を取得する
	items, err = client.Todo.Query().
		Where(
			todo.Not(
				todo.HasParent(),
			),
			todo.HasChildren(),
		).
		All(ctx)
	if err != nil {
		log.Fatalf("failed querying todos: %v", err)
	}
	fmt.Println("// 他の TODO に依存しておらず、子を持つ TODO を取得する")
	for _, t := range items {
		fmt.Printf("%d: %q\n", t.ID, t.Text)
	}

	// 子から親を取得する
	parent, err := client.Todo.Query(). // すべての TODO を取得する
						Where(todo.HasParent()). // 親 TODO を持つ TODO のみにフィルターする
						QueryParent().           // 親 TODO について走査する
						Only(ctx)                // 1 件のみ取得する
	if err != nil {
		log.Fatalf("failed querying todos: %v", err)
	}
	fmt.Printf("%d: %q\n", parent.ID, parent.Text)
}
