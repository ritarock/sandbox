// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gqlgen-echo-ent/ent/predicate"
	"gqlgen-echo-ent/ent/task"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (tu *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetTitle sets the "title" field.
func (tu *TaskUpdate) SetTitle(s string) *TaskUpdate {
	tu.mutation.SetTitle(s)
	return tu
}

// SetNote sets the "note" field.
func (tu *TaskUpdate) SetNote(s string) *TaskUpdate {
	tu.mutation.SetNote(s)
	return tu
}

// SetCompleted sets the "completed" field.
func (tu *TaskUpdate) SetCompleted(i int) *TaskUpdate {
	tu.mutation.ResetCompleted()
	tu.mutation.SetCompleted(i)
	return tu
}

// SetNillableCompleted sets the "completed" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableCompleted(i *int) *TaskUpdate {
	if i != nil {
		tu.SetCompleted(*i)
	}
	return tu
}

// AddCompleted adds i to the "completed" field.
func (tu *TaskUpdate) AddCompleted(i int) *TaskUpdate {
	tu.mutation.AddCompleted(i)
	return tu
}

// SetCreatedAt sets the "created_at" field.
func (tu *TaskUpdate) SetCreatedAt(t time.Time) *TaskUpdate {
	tu.mutation.SetCreatedAt(t)
	return tu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableCreatedAt(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetCreatedAt(*t)
	}
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TaskUpdate) SetUpdatedAt(t time.Time) *TaskUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableUpdatedAt(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetUpdatedAt(*t)
	}
	return tu
}

// Mutation returns the TaskMutation object of the builder.
func (tu *TaskUpdate) Mutation() *TaskMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TaskUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, TaskMutation](ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaskUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaskUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Title(); ok {
		_spec.SetField(task.FieldTitle, field.TypeString, value)
	}
	if value, ok := tu.mutation.Note(); ok {
		_spec.SetField(task.FieldNote, field.TypeString, value)
	}
	if value, ok := tu.mutation.Completed(); ok {
		_spec.SetField(task.FieldCompleted, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedCompleted(); ok {
		_spec.AddField(task.FieldCompleted, field.TypeInt, value)
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.SetField(task.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(task.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskMutation
}

// SetTitle sets the "title" field.
func (tuo *TaskUpdateOne) SetTitle(s string) *TaskUpdateOne {
	tuo.mutation.SetTitle(s)
	return tuo
}

// SetNote sets the "note" field.
func (tuo *TaskUpdateOne) SetNote(s string) *TaskUpdateOne {
	tuo.mutation.SetNote(s)
	return tuo
}

// SetCompleted sets the "completed" field.
func (tuo *TaskUpdateOne) SetCompleted(i int) *TaskUpdateOne {
	tuo.mutation.ResetCompleted()
	tuo.mutation.SetCompleted(i)
	return tuo
}

// SetNillableCompleted sets the "completed" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableCompleted(i *int) *TaskUpdateOne {
	if i != nil {
		tuo.SetCompleted(*i)
	}
	return tuo
}

// AddCompleted adds i to the "completed" field.
func (tuo *TaskUpdateOne) AddCompleted(i int) *TaskUpdateOne {
	tuo.mutation.AddCompleted(i)
	return tuo
}

// SetCreatedAt sets the "created_at" field.
func (tuo *TaskUpdateOne) SetCreatedAt(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetCreatedAt(t)
	return tuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableCreatedAt(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetCreatedAt(*t)
	}
	return tuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TaskUpdateOne) SetUpdatedAt(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableUpdatedAt(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetUpdatedAt(*t)
	}
	return tuo
}

// Mutation returns the TaskMutation object of the builder.
func (tuo *TaskUpdateOne) Mutation() *TaskMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (tuo *TaskUpdateOne) Where(ps ...predicate.Task) *TaskUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TaskUpdateOne) Select(field string, fields ...string) *TaskUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Task entity.
func (tuo *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	return withHooks[*Task, TaskMutation](ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TaskUpdateOne) sqlSave(ctx context.Context) (_node *Task, err error) {
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Task.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, task.FieldID)
		for _, f := range fields {
			if !task.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != task.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Title(); ok {
		_spec.SetField(task.FieldTitle, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Note(); ok {
		_spec.SetField(task.FieldNote, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Completed(); ok {
		_spec.SetField(task.FieldCompleted, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedCompleted(); ok {
		_spec.AddField(task.FieldCompleted, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.SetField(task.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(task.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Task{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}