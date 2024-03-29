// Code generated by ent, DO NOT EDIT.

package ent

import (
	"gqlgen-todo/ent/schema"
	"gqlgen-todo/ent/todo"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	// todoDescCompleted is the schema descriptor for completed field.
	todoDescCompleted := todoFields[2].Descriptor()
	// todo.DefaultCompleted holds the default value on creation for the completed field.
	todo.DefaultCompleted = todoDescCompleted.Default.(bool)
	// todoDescCreatedAt is the schema descriptor for created_at field.
	todoDescCreatedAt := todoFields[3].Descriptor()
	// todo.DefaultCreatedAt holds the default value on creation for the created_at field.
	todo.DefaultCreatedAt = todoDescCreatedAt.Default.(time.Time)
	// todoDescUpdatedAt is the schema descriptor for updated_at field.
	todoDescUpdatedAt := todoFields[4].Descriptor()
	// todo.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	todo.DefaultUpdatedAt = todoDescUpdatedAt.Default.(time.Time)
}
