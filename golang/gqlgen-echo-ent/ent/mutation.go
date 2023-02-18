// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gqlgen-echo-ent/ent/predicate"
	"gqlgen-echo-ent/ent/task"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTask = "Task"
)

// TaskMutation represents an operation that mutates the Task nodes in the graph.
type TaskMutation struct {
	config
	op            Op
	typ           string
	id            *int
	title         *string
	note          *string
	completed     *int
	addcompleted  *int
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Task, error)
	predicates    []predicate.Task
}

var _ ent.Mutation = (*TaskMutation)(nil)

// taskOption allows management of the mutation configuration using functional options.
type taskOption func(*TaskMutation)

// newTaskMutation creates new mutation for the Task entity.
func newTaskMutation(c config, op Op, opts ...taskOption) *TaskMutation {
	m := &TaskMutation{
		config:        c,
		op:            op,
		typ:           TypeTask,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTaskID sets the ID field of the mutation.
func withTaskID(id int) taskOption {
	return func(m *TaskMutation) {
		var (
			err   error
			once  sync.Once
			value *Task
		)
		m.oldValue = func(ctx context.Context) (*Task, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Task.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTask sets the old Task of the mutation.
func withTask(node *Task) taskOption {
	return func(m *TaskMutation) {
		m.oldValue = func(context.Context) (*Task, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TaskMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TaskMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TaskMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TaskMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Task.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTitle sets the "title" field.
func (m *TaskMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *TaskMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Task entity.
// If the Task object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TaskMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *TaskMutation) ResetTitle() {
	m.title = nil
}

// SetNote sets the "note" field.
func (m *TaskMutation) SetNote(s string) {
	m.note = &s
}

// Note returns the value of the "note" field in the mutation.
func (m *TaskMutation) Note() (r string, exists bool) {
	v := m.note
	if v == nil {
		return
	}
	return *v, true
}

// OldNote returns the old "note" field's value of the Task entity.
// If the Task object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TaskMutation) OldNote(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldNote is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldNote requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNote: %w", err)
	}
	return oldValue.Note, nil
}

// ResetNote resets all changes to the "note" field.
func (m *TaskMutation) ResetNote() {
	m.note = nil
}

// SetCompleted sets the "completed" field.
func (m *TaskMutation) SetCompleted(i int) {
	m.completed = &i
	m.addcompleted = nil
}

// Completed returns the value of the "completed" field in the mutation.
func (m *TaskMutation) Completed() (r int, exists bool) {
	v := m.completed
	if v == nil {
		return
	}
	return *v, true
}

// OldCompleted returns the old "completed" field's value of the Task entity.
// If the Task object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TaskMutation) OldCompleted(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCompleted is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCompleted requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCompleted: %w", err)
	}
	return oldValue.Completed, nil
}

// AddCompleted adds i to the "completed" field.
func (m *TaskMutation) AddCompleted(i int) {
	if m.addcompleted != nil {
		*m.addcompleted += i
	} else {
		m.addcompleted = &i
	}
}

// AddedCompleted returns the value that was added to the "completed" field in this mutation.
func (m *TaskMutation) AddedCompleted() (r int, exists bool) {
	v := m.addcompleted
	if v == nil {
		return
	}
	return *v, true
}

// ResetCompleted resets all changes to the "completed" field.
func (m *TaskMutation) ResetCompleted() {
	m.completed = nil
	m.addcompleted = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *TaskMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *TaskMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Task entity.
// If the Task object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TaskMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *TaskMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *TaskMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *TaskMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Task entity.
// If the Task object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TaskMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *TaskMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the TaskMutation builder.
func (m *TaskMutation) Where(ps ...predicate.Task) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the TaskMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *TaskMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Task, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *TaskMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *TaskMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Task).
func (m *TaskMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TaskMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.title != nil {
		fields = append(fields, task.FieldTitle)
	}
	if m.note != nil {
		fields = append(fields, task.FieldNote)
	}
	if m.completed != nil {
		fields = append(fields, task.FieldCompleted)
	}
	if m.created_at != nil {
		fields = append(fields, task.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, task.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TaskMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case task.FieldTitle:
		return m.Title()
	case task.FieldNote:
		return m.Note()
	case task.FieldCompleted:
		return m.Completed()
	case task.FieldCreatedAt:
		return m.CreatedAt()
	case task.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TaskMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case task.FieldTitle:
		return m.OldTitle(ctx)
	case task.FieldNote:
		return m.OldNote(ctx)
	case task.FieldCompleted:
		return m.OldCompleted(ctx)
	case task.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case task.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Task field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TaskMutation) SetField(name string, value ent.Value) error {
	switch name {
	case task.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case task.FieldNote:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNote(v)
		return nil
	case task.FieldCompleted:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCompleted(v)
		return nil
	case task.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case task.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Task field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TaskMutation) AddedFields() []string {
	var fields []string
	if m.addcompleted != nil {
		fields = append(fields, task.FieldCompleted)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TaskMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case task.FieldCompleted:
		return m.AddedCompleted()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TaskMutation) AddField(name string, value ent.Value) error {
	switch name {
	case task.FieldCompleted:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddCompleted(v)
		return nil
	}
	return fmt.Errorf("unknown Task numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TaskMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TaskMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TaskMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Task nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TaskMutation) ResetField(name string) error {
	switch name {
	case task.FieldTitle:
		m.ResetTitle()
		return nil
	case task.FieldNote:
		m.ResetNote()
		return nil
	case task.FieldCompleted:
		m.ResetCompleted()
		return nil
	case task.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case task.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Task field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TaskMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TaskMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TaskMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TaskMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TaskMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TaskMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TaskMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Task unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TaskMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Task edge %s", name)
}
