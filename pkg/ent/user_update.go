// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nishinoyama/kobuy-2/pkg/ent/grocery"
	"github.com/nishinoyama/kobuy-2/pkg/ent/ledger"
	"github.com/nishinoyama/kobuy-2/pkg/ent/predicate"
	"github.com/nishinoyama/kobuy-2/pkg/ent/purchase"
	"github.com/nishinoyama/kobuy-2/pkg/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetUserID sets the "user_id" field.
func (uu *UserUpdate) SetUserID(s string) *UserUpdate {
	uu.mutation.SetUserID(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetBalance sets the "balance" field.
func (uu *UserUpdate) SetBalance(i int) *UserUpdate {
	uu.mutation.ResetBalance()
	uu.mutation.SetBalance(i)
	return uu
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (uu *UserUpdate) SetNillableBalance(i *int) *UserUpdate {
	if i != nil {
		uu.SetBalance(*i)
	}
	return uu
}

// AddBalance adds i to the "balance" field.
func (uu *UserUpdate) AddBalance(i int) *UserUpdate {
	uu.mutation.AddBalance(i)
	return uu
}

// AddProvidedGroceryIDs adds the "provided_groceries" edge to the Grocery entity by IDs.
func (uu *UserUpdate) AddProvidedGroceryIDs(ids ...int) *UserUpdate {
	uu.mutation.AddProvidedGroceryIDs(ids...)
	return uu
}

// AddProvidedGroceries adds the "provided_groceries" edges to the Grocery entity.
func (uu *UserUpdate) AddProvidedGroceries(g ...*Grocery) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.AddProvidedGroceryIDs(ids...)
}

// AddPurchasedIDs adds the "purchased" edge to the Purchase entity by IDs.
func (uu *UserUpdate) AddPurchasedIDs(ids ...int) *UserUpdate {
	uu.mutation.AddPurchasedIDs(ids...)
	return uu
}

// AddPurchased adds the "purchased" edges to the Purchase entity.
func (uu *UserUpdate) AddPurchased(p ...*Purchase) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddPurchasedIDs(ids...)
}

// AddPayerIDs adds the "payer" edge to the Ledger entity by IDs.
func (uu *UserUpdate) AddPayerIDs(ids ...int) *UserUpdate {
	uu.mutation.AddPayerIDs(ids...)
	return uu
}

// AddPayer adds the "payer" edges to the Ledger entity.
func (uu *UserUpdate) AddPayer(l ...*Ledger) *UserUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return uu.AddPayerIDs(ids...)
}

// AddReceiverIDs adds the "receiver" edge to the Ledger entity by IDs.
func (uu *UserUpdate) AddReceiverIDs(ids ...int) *UserUpdate {
	uu.mutation.AddReceiverIDs(ids...)
	return uu
}

// AddReceiver adds the "receiver" edges to the Ledger entity.
func (uu *UserUpdate) AddReceiver(l ...*Ledger) *UserUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return uu.AddReceiverIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearProvidedGroceries clears all "provided_groceries" edges to the Grocery entity.
func (uu *UserUpdate) ClearProvidedGroceries() *UserUpdate {
	uu.mutation.ClearProvidedGroceries()
	return uu
}

// RemoveProvidedGroceryIDs removes the "provided_groceries" edge to Grocery entities by IDs.
func (uu *UserUpdate) RemoveProvidedGroceryIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveProvidedGroceryIDs(ids...)
	return uu
}

// RemoveProvidedGroceries removes "provided_groceries" edges to Grocery entities.
func (uu *UserUpdate) RemoveProvidedGroceries(g ...*Grocery) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.RemoveProvidedGroceryIDs(ids...)
}

// ClearPurchased clears all "purchased" edges to the Purchase entity.
func (uu *UserUpdate) ClearPurchased() *UserUpdate {
	uu.mutation.ClearPurchased()
	return uu
}

// RemovePurchasedIDs removes the "purchased" edge to Purchase entities by IDs.
func (uu *UserUpdate) RemovePurchasedIDs(ids ...int) *UserUpdate {
	uu.mutation.RemovePurchasedIDs(ids...)
	return uu
}

// RemovePurchased removes "purchased" edges to Purchase entities.
func (uu *UserUpdate) RemovePurchased(p ...*Purchase) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemovePurchasedIDs(ids...)
}

// ClearPayer clears all "payer" edges to the Ledger entity.
func (uu *UserUpdate) ClearPayer() *UserUpdate {
	uu.mutation.ClearPayer()
	return uu
}

// RemovePayerIDs removes the "payer" edge to Ledger entities by IDs.
func (uu *UserUpdate) RemovePayerIDs(ids ...int) *UserUpdate {
	uu.mutation.RemovePayerIDs(ids...)
	return uu
}

// RemovePayer removes "payer" edges to Ledger entities.
func (uu *UserUpdate) RemovePayer(l ...*Ledger) *UserUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return uu.RemovePayerIDs(ids...)
}

// ClearReceiver clears all "receiver" edges to the Ledger entity.
func (uu *UserUpdate) ClearReceiver() *UserUpdate {
	uu.mutation.ClearReceiver()
	return uu
}

// RemoveReceiverIDs removes the "receiver" edge to Ledger entities by IDs.
func (uu *UserUpdate) RemoveReceiverIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveReceiverIDs(ids...)
	return uu
}

// RemoveReceiver removes "receiver" edges to Ledger entities.
func (uu *UserUpdate) RemoveReceiver(l ...*Ledger) *UserUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return uu.RemoveReceiverIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.UserID(); ok {
		_spec.SetField(user.FieldUserID, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.Balance(); ok {
		_spec.SetField(user.FieldBalance, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedBalance(); ok {
		_spec.AddField(user.FieldBalance, field.TypeInt, value)
	}
	if uu.mutation.ProvidedGroceriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProvidedGroceriesTable,
			Columns: []string{user.ProvidedGroceriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grocery.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedProvidedGroceriesIDs(); len(nodes) > 0 && !uu.mutation.ProvidedGroceriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProvidedGroceriesTable,
			Columns: []string{user.ProvidedGroceriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grocery.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ProvidedGroceriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProvidedGroceriesTable,
			Columns: []string{user.ProvidedGroceriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grocery.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.PurchasedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PurchasedTable,
			Columns: []string{user.PurchasedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(purchase.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedPurchasedIDs(); len(nodes) > 0 && !uu.mutation.PurchasedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PurchasedTable,
			Columns: []string{user.PurchasedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(purchase.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PurchasedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PurchasedTable,
			Columns: []string{user.PurchasedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(purchase.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.PayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PayerTable,
			Columns: []string{user.PayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ledger.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedPayerIDs(); len(nodes) > 0 && !uu.mutation.PayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PayerTable,
			Columns: []string{user.PayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ledger.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PayerTable,
			Columns: []string{user.PayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ledger.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.ReceiverCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReceiverTable,
			Columns: []string{user.ReceiverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ledger.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedReceiverIDs(); len(nodes) > 0 && !uu.mutation.ReceiverCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReceiverTable,
			Columns: []string{user.ReceiverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ledger.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ReceiverIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReceiverTable,
			Columns: []string{user.ReceiverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ledger.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetUserID sets the "user_id" field.
func (uuo *UserUpdateOne) SetUserID(s string) *UserUpdateOne {
	uuo.mutation.SetUserID(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetBalance sets the "balance" field.
func (uuo *UserUpdateOne) SetBalance(i int) *UserUpdateOne {
	uuo.mutation.ResetBalance()
	uuo.mutation.SetBalance(i)
	return uuo
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBalance(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetBalance(*i)
	}
	return uuo
}

// AddBalance adds i to the "balance" field.
func (uuo *UserUpdateOne) AddBalance(i int) *UserUpdateOne {
	uuo.mutation.AddBalance(i)
	return uuo
}

// AddProvidedGroceryIDs adds the "provided_groceries" edge to the Grocery entity by IDs.
func (uuo *UserUpdateOne) AddProvidedGroceryIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddProvidedGroceryIDs(ids...)
	return uuo
}

// AddProvidedGroceries adds the "provided_groceries" edges to the Grocery entity.
func (uuo *UserUpdateOne) AddProvidedGroceries(g ...*Grocery) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.AddProvidedGroceryIDs(ids...)
}

// AddPurchasedIDs adds the "purchased" edge to the Purchase entity by IDs.
func (uuo *UserUpdateOne) AddPurchasedIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddPurchasedIDs(ids...)
	return uuo
}

// AddPurchased adds the "purchased" edges to the Purchase entity.
func (uuo *UserUpdateOne) AddPurchased(p ...*Purchase) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddPurchasedIDs(ids...)
}

// AddPayerIDs adds the "payer" edge to the Ledger entity by IDs.
func (uuo *UserUpdateOne) AddPayerIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddPayerIDs(ids...)
	return uuo
}

// AddPayer adds the "payer" edges to the Ledger entity.
func (uuo *UserUpdateOne) AddPayer(l ...*Ledger) *UserUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return uuo.AddPayerIDs(ids...)
}

// AddReceiverIDs adds the "receiver" edge to the Ledger entity by IDs.
func (uuo *UserUpdateOne) AddReceiverIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddReceiverIDs(ids...)
	return uuo
}

// AddReceiver adds the "receiver" edges to the Ledger entity.
func (uuo *UserUpdateOne) AddReceiver(l ...*Ledger) *UserUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return uuo.AddReceiverIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearProvidedGroceries clears all "provided_groceries" edges to the Grocery entity.
func (uuo *UserUpdateOne) ClearProvidedGroceries() *UserUpdateOne {
	uuo.mutation.ClearProvidedGroceries()
	return uuo
}

// RemoveProvidedGroceryIDs removes the "provided_groceries" edge to Grocery entities by IDs.
func (uuo *UserUpdateOne) RemoveProvidedGroceryIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveProvidedGroceryIDs(ids...)
	return uuo
}

// RemoveProvidedGroceries removes "provided_groceries" edges to Grocery entities.
func (uuo *UserUpdateOne) RemoveProvidedGroceries(g ...*Grocery) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.RemoveProvidedGroceryIDs(ids...)
}

// ClearPurchased clears all "purchased" edges to the Purchase entity.
func (uuo *UserUpdateOne) ClearPurchased() *UserUpdateOne {
	uuo.mutation.ClearPurchased()
	return uuo
}

// RemovePurchasedIDs removes the "purchased" edge to Purchase entities by IDs.
func (uuo *UserUpdateOne) RemovePurchasedIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemovePurchasedIDs(ids...)
	return uuo
}

// RemovePurchased removes "purchased" edges to Purchase entities.
func (uuo *UserUpdateOne) RemovePurchased(p ...*Purchase) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemovePurchasedIDs(ids...)
}

// ClearPayer clears all "payer" edges to the Ledger entity.
func (uuo *UserUpdateOne) ClearPayer() *UserUpdateOne {
	uuo.mutation.ClearPayer()
	return uuo
}

// RemovePayerIDs removes the "payer" edge to Ledger entities by IDs.
func (uuo *UserUpdateOne) RemovePayerIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemovePayerIDs(ids...)
	return uuo
}

// RemovePayer removes "payer" edges to Ledger entities.
func (uuo *UserUpdateOne) RemovePayer(l ...*Ledger) *UserUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return uuo.RemovePayerIDs(ids...)
}

// ClearReceiver clears all "receiver" edges to the Ledger entity.
func (uuo *UserUpdateOne) ClearReceiver() *UserUpdateOne {
	uuo.mutation.ClearReceiver()
	return uuo
}

// RemoveReceiverIDs removes the "receiver" edge to Ledger entities by IDs.
func (uuo *UserUpdateOne) RemoveReceiverIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveReceiverIDs(ids...)
	return uuo
}

// RemoveReceiver removes "receiver" edges to Ledger entities.
func (uuo *UserUpdateOne) RemoveReceiver(l ...*Ledger) *UserUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return uuo.RemoveReceiverIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.UserID(); ok {
		_spec.SetField(user.FieldUserID, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Balance(); ok {
		_spec.SetField(user.FieldBalance, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedBalance(); ok {
		_spec.AddField(user.FieldBalance, field.TypeInt, value)
	}
	if uuo.mutation.ProvidedGroceriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProvidedGroceriesTable,
			Columns: []string{user.ProvidedGroceriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grocery.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedProvidedGroceriesIDs(); len(nodes) > 0 && !uuo.mutation.ProvidedGroceriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProvidedGroceriesTable,
			Columns: []string{user.ProvidedGroceriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grocery.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ProvidedGroceriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProvidedGroceriesTable,
			Columns: []string{user.ProvidedGroceriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grocery.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.PurchasedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PurchasedTable,
			Columns: []string{user.PurchasedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(purchase.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedPurchasedIDs(); len(nodes) > 0 && !uuo.mutation.PurchasedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PurchasedTable,
			Columns: []string{user.PurchasedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(purchase.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PurchasedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PurchasedTable,
			Columns: []string{user.PurchasedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(purchase.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.PayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PayerTable,
			Columns: []string{user.PayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ledger.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedPayerIDs(); len(nodes) > 0 && !uuo.mutation.PayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PayerTable,
			Columns: []string{user.PayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ledger.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PayerTable,
			Columns: []string{user.PayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ledger.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.ReceiverCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReceiverTable,
			Columns: []string{user.ReceiverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ledger.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedReceiverIDs(); len(nodes) > 0 && !uuo.mutation.ReceiverCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReceiverTable,
			Columns: []string{user.ReceiverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ledger.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ReceiverIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReceiverTable,
			Columns: []string{user.ReceiverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ledger.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
