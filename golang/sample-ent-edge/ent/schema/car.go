package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.String("model"),
		field.Time("registered_at"),
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		// User 型の owner という逆エッジを作成
		edge.From("owner", User.Type).
			// Ref を使って明示的に User スキーマの cars エッジを参照する
			Ref("cars").
			// cars は 一人のオーナーのみが所有することを保証する
			Unique(),
	}
}
