// Code generated by entc, DO NOT EDIT.

package todo

import (
	"time"
)

const (
	// Label holds the string label denoting the todo type in the database.
	Label = "todo"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldTodoID holds the string denoting the todo_id field in the database.
	FieldTodoID = "todo_id"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the todo in the database.
	Table = "todos"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "todos"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "todo_id"
)

// Columns holds all SQL columns for todo fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldStatus,
	FieldCreatedAt,
	FieldTodoID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
