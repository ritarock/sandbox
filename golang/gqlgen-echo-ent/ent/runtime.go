// Code generated by ent, DO NOT EDIT.

package ent

import (
	"gqlgen-echo-ent/ent/schema"
	"gqlgen-echo-ent/ent/task"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescCompleted is the schema descriptor for completed field.
	taskDescCompleted := taskFields[2].Descriptor()
	// task.DefaultCompleted holds the default value on creation for the completed field.
	task.DefaultCompleted = taskDescCompleted.Default.(int)
	// taskDescCreatedAt is the schema descriptor for created_at field.
	taskDescCreatedAt := taskFields[3].Descriptor()
	// task.DefaultCreatedAt holds the default value on creation for the created_at field.
	task.DefaultCreatedAt = taskDescCreatedAt.Default.(time.Time)
	// taskDescUpdatedAt is the schema descriptor for updated_at field.
	taskDescUpdatedAt := taskFields[4].Descriptor()
	// task.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	task.DefaultUpdatedAt = taskDescUpdatedAt.Default.(time.Time)
}
