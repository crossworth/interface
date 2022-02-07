// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/bug/ent/car"
	"entgo.io/bug/ent/plane"
	"entgo.io/bug/ent/predicate"
	"entgo.io/bug/ent/vehicle"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCar     = "Car"
	TypePlane   = "Plane"
	TypeVehicle = "Vehicle"
)

// CarMutation represents an operation that mutates the Car nodes in the graph.
type CarMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Car, error)
	predicates    []predicate.Car
}

var _ ent.Mutation = (*CarMutation)(nil)

// carOption allows management of the mutation configuration using functional options.
type carOption func(*CarMutation)

// newCarMutation creates new mutation for the Car entity.
func newCarMutation(c config, op Op, opts ...carOption) *CarMutation {
	m := &CarMutation{
		config:        c,
		op:            op,
		typ:           TypeCar,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCarID sets the ID field of the mutation.
func withCarID(id int) carOption {
	return func(m *CarMutation) {
		var (
			err   error
			once  sync.Once
			value *Car
		)
		m.oldValue = func(ctx context.Context) (*Car, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Car.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCar sets the old Car of the mutation.
func withCar(node *Car) carOption {
	return func(m *CarMutation) {
		m.oldValue = func(context.Context) (*Car, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CarMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CarMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CarMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *CarMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Car.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *CarMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *CarMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Car entity.
// If the Car object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CarMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *CarMutation) ResetName() {
	m.name = nil
}

// Where appends a list predicates to the CarMutation builder.
func (m *CarMutation) Where(ps ...predicate.Car) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *CarMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Car).
func (m *CarMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CarMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, car.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CarMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case car.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CarMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case car.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Car field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CarMutation) SetField(name string, value ent.Value) error {
	switch name {
	case car.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Car field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CarMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CarMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CarMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Car numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CarMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CarMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CarMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Car nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CarMutation) ResetField(name string) error {
	switch name {
	case car.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Car field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CarMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CarMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CarMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CarMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CarMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CarMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CarMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Car unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CarMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Car edge %s", name)
}

// PlaneMutation represents an operation that mutates the Plane nodes in the graph.
type PlaneMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Plane, error)
	predicates    []predicate.Plane
}

var _ ent.Mutation = (*PlaneMutation)(nil)

// planeOption allows management of the mutation configuration using functional options.
type planeOption func(*PlaneMutation)

// newPlaneMutation creates new mutation for the Plane entity.
func newPlaneMutation(c config, op Op, opts ...planeOption) *PlaneMutation {
	m := &PlaneMutation{
		config:        c,
		op:            op,
		typ:           TypePlane,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withPlaneID sets the ID field of the mutation.
func withPlaneID(id int) planeOption {
	return func(m *PlaneMutation) {
		var (
			err   error
			once  sync.Once
			value *Plane
		)
		m.oldValue = func(ctx context.Context) (*Plane, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Plane.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withPlane sets the old Plane of the mutation.
func withPlane(node *Plane) planeOption {
	return func(m *PlaneMutation) {
		m.oldValue = func(context.Context) (*Plane, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m PlaneMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m PlaneMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *PlaneMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *PlaneMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Plane.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *PlaneMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *PlaneMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Plane entity.
// If the Plane object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PlaneMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *PlaneMutation) ResetName() {
	m.name = nil
}

// Where appends a list predicates to the PlaneMutation builder.
func (m *PlaneMutation) Where(ps ...predicate.Plane) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *PlaneMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Plane).
func (m *PlaneMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *PlaneMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, plane.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *PlaneMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case plane.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *PlaneMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case plane.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Plane field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PlaneMutation) SetField(name string, value ent.Value) error {
	switch name {
	case plane.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Plane field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *PlaneMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *PlaneMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PlaneMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Plane numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *PlaneMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *PlaneMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *PlaneMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Plane nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *PlaneMutation) ResetField(name string) error {
	switch name {
	case plane.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Plane field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *PlaneMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *PlaneMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *PlaneMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *PlaneMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *PlaneMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *PlaneMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *PlaneMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Plane unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *PlaneMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Plane edge %s", name)
}

// VehicleMutation represents an operation that mutates the Vehicle nodes in the graph.
type VehicleMutation struct {
	config
	op            Op
	typ           string
	id            *string
	_type         *string
	name          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Vehicle, error)
	predicates    []predicate.Vehicle
}

var _ ent.Mutation = (*VehicleMutation)(nil)

// vehicleOption allows management of the mutation configuration using functional options.
type vehicleOption func(*VehicleMutation)

// newVehicleMutation creates new mutation for the Vehicle entity.
func newVehicleMutation(c config, op Op, opts ...vehicleOption) *VehicleMutation {
	m := &VehicleMutation{
		config:        c,
		op:            op,
		typ:           TypeVehicle,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withVehicleID sets the ID field of the mutation.
func withVehicleID(id string) vehicleOption {
	return func(m *VehicleMutation) {
		var (
			err   error
			once  sync.Once
			value *Vehicle
		)
		m.oldValue = func(ctx context.Context) (*Vehicle, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Vehicle.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withVehicle sets the old Vehicle of the mutation.
func withVehicle(node *Vehicle) vehicleOption {
	return func(m *VehicleMutation) {
		m.oldValue = func(context.Context) (*Vehicle, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m VehicleMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m VehicleMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Vehicle entities.
func (m *VehicleMutation) SetID(id string) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *VehicleMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *VehicleMutation) IDs(ctx context.Context) ([]string, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []string{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Vehicle.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetType sets the "type" field.
func (m *VehicleMutation) SetType(s string) {
	m._type = &s
}

// GetType returns the value of the "type" field in the mutation.
func (m *VehicleMutation) GetType() (r string, exists bool) {
	v := m._type
	if v == nil {
		return
	}
	return *v, true
}

// OldType returns the old "type" field's value of the Vehicle entity.
// If the Vehicle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VehicleMutation) OldType(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldType: %w", err)
	}
	return oldValue.Type, nil
}

// ResetType resets all changes to the "type" field.
func (m *VehicleMutation) ResetType() {
	m._type = nil
}

// SetName sets the "name" field.
func (m *VehicleMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *VehicleMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Vehicle entity.
// If the Vehicle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VehicleMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *VehicleMutation) ResetName() {
	m.name = nil
}

// Where appends a list predicates to the VehicleMutation builder.
func (m *VehicleMutation) Where(ps ...predicate.Vehicle) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *VehicleMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Vehicle).
func (m *VehicleMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *VehicleMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m._type != nil {
		fields = append(fields, vehicle.FieldType)
	}
	if m.name != nil {
		fields = append(fields, vehicle.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *VehicleMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case vehicle.FieldType:
		return m.GetType()
	case vehicle.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *VehicleMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case vehicle.FieldType:
		return m.OldType(ctx)
	case vehicle.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Vehicle field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *VehicleMutation) SetField(name string, value ent.Value) error {
	switch name {
	case vehicle.FieldType:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetType(v)
		return nil
	case vehicle.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Vehicle field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *VehicleMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *VehicleMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *VehicleMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Vehicle numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *VehicleMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *VehicleMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *VehicleMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Vehicle nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *VehicleMutation) ResetField(name string) error {
	switch name {
	case vehicle.FieldType:
		m.ResetType()
		return nil
	case vehicle.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Vehicle field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *VehicleMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *VehicleMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *VehicleMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *VehicleMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *VehicleMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *VehicleMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *VehicleMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Vehicle unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *VehicleMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Vehicle edge %s", name)
}
