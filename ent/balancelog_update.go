// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nishinoyama/kobuy-2/ent/balancelog"
	"github.com/nishinoyama/kobuy-2/ent/predicate"
	"github.com/nishinoyama/kobuy-2/ent/user"
)

// BalanceLogUpdate is the builder for updating BalanceLog entities.
type BalanceLogUpdate struct {
	config
	hooks    []Hook
	mutation *BalanceLogMutation
}

// Where appends a list predicates to the BalanceLogUpdate builder.
func (blu *BalanceLogUpdate) Where(ps ...predicate.BalanceLog) *BalanceLogUpdate {
	blu.mutation.Where(ps...)
	return blu
}

// SetPrice sets the "price" field.
func (blu *BalanceLogUpdate) SetPrice(i int) *BalanceLogUpdate {
	blu.mutation.ResetPrice()
	blu.mutation.SetPrice(i)
	return blu
}

// AddPrice adds i to the "price" field.
func (blu *BalanceLogUpdate) AddPrice(i int) *BalanceLogUpdate {
	blu.mutation.AddPrice(i)
	return blu
}

// SetType sets the "type" field.
func (blu *BalanceLogUpdate) SetType(b balancelog.Type) *BalanceLogUpdate {
	blu.mutation.SetType(b)
	return blu
}

// SetDonorID sets the "donor" edge to the User entity by ID.
func (blu *BalanceLogUpdate) SetDonorID(id int) *BalanceLogUpdate {
	blu.mutation.SetDonorID(id)
	return blu
}

// SetNillableDonorID sets the "donor" edge to the User entity by ID if the given value is not nil.
func (blu *BalanceLogUpdate) SetNillableDonorID(id *int) *BalanceLogUpdate {
	if id != nil {
		blu = blu.SetDonorID(*id)
	}
	return blu
}

// SetDonor sets the "donor" edge to the User entity.
func (blu *BalanceLogUpdate) SetDonor(u *User) *BalanceLogUpdate {
	return blu.SetDonorID(u.ID)
}

// SetReceiverID sets the "receiver" edge to the User entity by ID.
func (blu *BalanceLogUpdate) SetReceiverID(id int) *BalanceLogUpdate {
	blu.mutation.SetReceiverID(id)
	return blu
}

// SetNillableReceiverID sets the "receiver" edge to the User entity by ID if the given value is not nil.
func (blu *BalanceLogUpdate) SetNillableReceiverID(id *int) *BalanceLogUpdate {
	if id != nil {
		blu = blu.SetReceiverID(*id)
	}
	return blu
}

// SetReceiver sets the "receiver" edge to the User entity.
func (blu *BalanceLogUpdate) SetReceiver(u *User) *BalanceLogUpdate {
	return blu.SetReceiverID(u.ID)
}

// Mutation returns the BalanceLogMutation object of the builder.
func (blu *BalanceLogUpdate) Mutation() *BalanceLogMutation {
	return blu.mutation
}

// ClearDonor clears the "donor" edge to the User entity.
func (blu *BalanceLogUpdate) ClearDonor() *BalanceLogUpdate {
	blu.mutation.ClearDonor()
	return blu
}

// ClearReceiver clears the "receiver" edge to the User entity.
func (blu *BalanceLogUpdate) ClearReceiver() *BalanceLogUpdate {
	blu.mutation.ClearReceiver()
	return blu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (blu *BalanceLogUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, BalanceLogMutation](ctx, blu.sqlSave, blu.mutation, blu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (blu *BalanceLogUpdate) SaveX(ctx context.Context) int {
	affected, err := blu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (blu *BalanceLogUpdate) Exec(ctx context.Context) error {
	_, err := blu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (blu *BalanceLogUpdate) ExecX(ctx context.Context) {
	if err := blu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (blu *BalanceLogUpdate) check() error {
	if v, ok := blu.mutation.GetType(); ok {
		if err := balancelog.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "BalanceLog.type": %w`, err)}
		}
	}
	return nil
}

func (blu *BalanceLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := blu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(balancelog.Table, balancelog.Columns, sqlgraph.NewFieldSpec(balancelog.FieldID, field.TypeInt))
	if ps := blu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := blu.mutation.Price(); ok {
		_spec.SetField(balancelog.FieldPrice, field.TypeInt, value)
	}
	if value, ok := blu.mutation.AddedPrice(); ok {
		_spec.AddField(balancelog.FieldPrice, field.TypeInt, value)
	}
	if value, ok := blu.mutation.GetType(); ok {
		_spec.SetField(balancelog.FieldType, field.TypeEnum, value)
	}
	if blu.mutation.DonorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   balancelog.DonorTable,
			Columns: []string{balancelog.DonorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := blu.mutation.DonorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   balancelog.DonorTable,
			Columns: []string{balancelog.DonorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if blu.mutation.ReceiverCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   balancelog.ReceiverTable,
			Columns: []string{balancelog.ReceiverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := blu.mutation.ReceiverIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   balancelog.ReceiverTable,
			Columns: []string{balancelog.ReceiverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, blu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{balancelog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	blu.mutation.done = true
	return n, nil
}

// BalanceLogUpdateOne is the builder for updating a single BalanceLog entity.
type BalanceLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BalanceLogMutation
}

// SetPrice sets the "price" field.
func (bluo *BalanceLogUpdateOne) SetPrice(i int) *BalanceLogUpdateOne {
	bluo.mutation.ResetPrice()
	bluo.mutation.SetPrice(i)
	return bluo
}

// AddPrice adds i to the "price" field.
func (bluo *BalanceLogUpdateOne) AddPrice(i int) *BalanceLogUpdateOne {
	bluo.mutation.AddPrice(i)
	return bluo
}

// SetType sets the "type" field.
func (bluo *BalanceLogUpdateOne) SetType(b balancelog.Type) *BalanceLogUpdateOne {
	bluo.mutation.SetType(b)
	return bluo
}

// SetDonorID sets the "donor" edge to the User entity by ID.
func (bluo *BalanceLogUpdateOne) SetDonorID(id int) *BalanceLogUpdateOne {
	bluo.mutation.SetDonorID(id)
	return bluo
}

// SetNillableDonorID sets the "donor" edge to the User entity by ID if the given value is not nil.
func (bluo *BalanceLogUpdateOne) SetNillableDonorID(id *int) *BalanceLogUpdateOne {
	if id != nil {
		bluo = bluo.SetDonorID(*id)
	}
	return bluo
}

// SetDonor sets the "donor" edge to the User entity.
func (bluo *BalanceLogUpdateOne) SetDonor(u *User) *BalanceLogUpdateOne {
	return bluo.SetDonorID(u.ID)
}

// SetReceiverID sets the "receiver" edge to the User entity by ID.
func (bluo *BalanceLogUpdateOne) SetReceiverID(id int) *BalanceLogUpdateOne {
	bluo.mutation.SetReceiverID(id)
	return bluo
}

// SetNillableReceiverID sets the "receiver" edge to the User entity by ID if the given value is not nil.
func (bluo *BalanceLogUpdateOne) SetNillableReceiverID(id *int) *BalanceLogUpdateOne {
	if id != nil {
		bluo = bluo.SetReceiverID(*id)
	}
	return bluo
}

// SetReceiver sets the "receiver" edge to the User entity.
func (bluo *BalanceLogUpdateOne) SetReceiver(u *User) *BalanceLogUpdateOne {
	return bluo.SetReceiverID(u.ID)
}

// Mutation returns the BalanceLogMutation object of the builder.
func (bluo *BalanceLogUpdateOne) Mutation() *BalanceLogMutation {
	return bluo.mutation
}

// ClearDonor clears the "donor" edge to the User entity.
func (bluo *BalanceLogUpdateOne) ClearDonor() *BalanceLogUpdateOne {
	bluo.mutation.ClearDonor()
	return bluo
}

// ClearReceiver clears the "receiver" edge to the User entity.
func (bluo *BalanceLogUpdateOne) ClearReceiver() *BalanceLogUpdateOne {
	bluo.mutation.ClearReceiver()
	return bluo
}

// Where appends a list predicates to the BalanceLogUpdate builder.
func (bluo *BalanceLogUpdateOne) Where(ps ...predicate.BalanceLog) *BalanceLogUpdateOne {
	bluo.mutation.Where(ps...)
	return bluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bluo *BalanceLogUpdateOne) Select(field string, fields ...string) *BalanceLogUpdateOne {
	bluo.fields = append([]string{field}, fields...)
	return bluo
}

// Save executes the query and returns the updated BalanceLog entity.
func (bluo *BalanceLogUpdateOne) Save(ctx context.Context) (*BalanceLog, error) {
	return withHooks[*BalanceLog, BalanceLogMutation](ctx, bluo.sqlSave, bluo.mutation, bluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bluo *BalanceLogUpdateOne) SaveX(ctx context.Context) *BalanceLog {
	node, err := bluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bluo *BalanceLogUpdateOne) Exec(ctx context.Context) error {
	_, err := bluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bluo *BalanceLogUpdateOne) ExecX(ctx context.Context) {
	if err := bluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bluo *BalanceLogUpdateOne) check() error {
	if v, ok := bluo.mutation.GetType(); ok {
		if err := balancelog.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "BalanceLog.type": %w`, err)}
		}
	}
	return nil
}

func (bluo *BalanceLogUpdateOne) sqlSave(ctx context.Context) (_node *BalanceLog, err error) {
	if err := bluo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(balancelog.Table, balancelog.Columns, sqlgraph.NewFieldSpec(balancelog.FieldID, field.TypeInt))
	id, ok := bluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BalanceLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, balancelog.FieldID)
		for _, f := range fields {
			if !balancelog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != balancelog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bluo.mutation.Price(); ok {
		_spec.SetField(balancelog.FieldPrice, field.TypeInt, value)
	}
	if value, ok := bluo.mutation.AddedPrice(); ok {
		_spec.AddField(balancelog.FieldPrice, field.TypeInt, value)
	}
	if value, ok := bluo.mutation.GetType(); ok {
		_spec.SetField(balancelog.FieldType, field.TypeEnum, value)
	}
	if bluo.mutation.DonorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   balancelog.DonorTable,
			Columns: []string{balancelog.DonorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bluo.mutation.DonorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   balancelog.DonorTable,
			Columns: []string{balancelog.DonorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bluo.mutation.ReceiverCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   balancelog.ReceiverTable,
			Columns: []string{balancelog.ReceiverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bluo.mutation.ReceiverIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   balancelog.ReceiverTable,
			Columns: []string{balancelog.ReceiverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BalanceLog{config: bluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{balancelog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bluo.mutation.done = true
	return _node, nil
}