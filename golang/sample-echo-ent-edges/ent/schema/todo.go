package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

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

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return nil
}
