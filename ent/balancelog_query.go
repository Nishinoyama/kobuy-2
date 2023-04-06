// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nishinoyama/kobuy-2/ent/balancelog"
	"github.com/nishinoyama/kobuy-2/ent/predicate"
	"github.com/nishinoyama/kobuy-2/ent/user"
)

// BalanceLogQuery is the builder for querying BalanceLog entities.
type BalanceLogQuery struct {
	config
	ctx          *QueryContext
	order        []OrderFunc
	inters       []Interceptor
	predicates   []predicate.BalanceLog
	withDonor    *UserQuery
	withReceiver *UserQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BalanceLogQuery builder.
func (blq *BalanceLogQuery) Where(ps ...predicate.BalanceLog) *BalanceLogQuery {
	blq.predicates = append(blq.predicates, ps...)
	return blq
}

// Limit the number of records to be returned by this query.
func (blq *BalanceLogQuery) Limit(limit int) *BalanceLogQuery {
	blq.ctx.Limit = &limit
	return blq
}

// Offset to start from.
func (blq *BalanceLogQuery) Offset(offset int) *BalanceLogQuery {
	blq.ctx.Offset = &offset
	return blq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (blq *BalanceLogQuery) Unique(unique bool) *BalanceLogQuery {
	blq.ctx.Unique = &unique
	return blq
}

// Order specifies how the records should be ordered.
func (blq *BalanceLogQuery) Order(o ...OrderFunc) *BalanceLogQuery {
	blq.order = append(blq.order, o...)
	return blq
}

// QueryDonor chains the current query on the "donor" edge.
func (blq *BalanceLogQuery) QueryDonor() *UserQuery {
	query := (&UserClient{config: blq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := blq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := blq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(balancelog.Table, balancelog.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, balancelog.DonorTable, balancelog.DonorColumn),
		)
		fromU = sqlgraph.SetNeighbors(blq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryReceiver chains the current query on the "receiver" edge.
func (blq *BalanceLogQuery) QueryReceiver() *UserQuery {
	query := (&UserClient{config: blq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := blq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := blq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(balancelog.Table, balancelog.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, balancelog.ReceiverTable, balancelog.ReceiverColumn),
		)
		fromU = sqlgraph.SetNeighbors(blq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BalanceLog entity from the query.
// Returns a *NotFoundError when no BalanceLog was found.
func (blq *BalanceLogQuery) First(ctx context.Context) (*BalanceLog, error) {
	nodes, err := blq.Limit(1).All(setContextOp(ctx, blq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{balancelog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (blq *BalanceLogQuery) FirstX(ctx context.Context) *BalanceLog {
	node, err := blq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BalanceLog ID from the query.
// Returns a *NotFoundError when no BalanceLog ID was found.
func (blq *BalanceLogQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = blq.Limit(1).IDs(setContextOp(ctx, blq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{balancelog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (blq *BalanceLogQuery) FirstIDX(ctx context.Context) int {
	id, err := blq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BalanceLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BalanceLog entity is found.
// Returns a *NotFoundError when no BalanceLog entities are found.
func (blq *BalanceLogQuery) Only(ctx context.Context) (*BalanceLog, error) {
	nodes, err := blq.Limit(2).All(setContextOp(ctx, blq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{balancelog.Label}
	default:
		return nil, &NotSingularError{balancelog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (blq *BalanceLogQuery) OnlyX(ctx context.Context) *BalanceLog {
	node, err := blq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BalanceLog ID in the query.
// Returns a *NotSingularError when more than one BalanceLog ID is found.
// Returns a *NotFoundError when no entities are found.
func (blq *BalanceLogQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = blq.Limit(2).IDs(setContextOp(ctx, blq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{balancelog.Label}
	default:
		err = &NotSingularError{balancelog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (blq *BalanceLogQuery) OnlyIDX(ctx context.Context) int {
	id, err := blq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BalanceLogs.
func (blq *BalanceLogQuery) All(ctx context.Context) ([]*BalanceLog, error) {
	ctx = setContextOp(ctx, blq.ctx, "All")
	if err := blq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BalanceLog, *BalanceLogQuery]()
	return withInterceptors[[]*BalanceLog](ctx, blq, qr, blq.inters)
}

// AllX is like All, but panics if an error occurs.
func (blq *BalanceLogQuery) AllX(ctx context.Context) []*BalanceLog {
	nodes, err := blq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BalanceLog IDs.
func (blq *BalanceLogQuery) IDs(ctx context.Context) (ids []int, err error) {
	if blq.ctx.Unique == nil && blq.path != nil {
		blq.Unique(true)
	}
	ctx = setContextOp(ctx, blq.ctx, "IDs")
	if err = blq.Select(balancelog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (blq *BalanceLogQuery) IDsX(ctx context.Context) []int {
	ids, err := blq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (blq *BalanceLogQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, blq.ctx, "Count")
	if err := blq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, blq, querierCount[*BalanceLogQuery](), blq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (blq *BalanceLogQuery) CountX(ctx context.Context) int {
	count, err := blq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (blq *BalanceLogQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, blq.ctx, "Exist")
	switch _, err := blq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (blq *BalanceLogQuery) ExistX(ctx context.Context) bool {
	exist, err := blq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BalanceLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (blq *BalanceLogQuery) Clone() *BalanceLogQuery {
	if blq == nil {
		return nil
	}
	return &BalanceLogQuery{
		config:       blq.config,
		ctx:          blq.ctx.Clone(),
		order:        append([]OrderFunc{}, blq.order...),
		inters:       append([]Interceptor{}, blq.inters...),
		predicates:   append([]predicate.BalanceLog{}, blq.predicates...),
		withDonor:    blq.withDonor.Clone(),
		withReceiver: blq.withReceiver.Clone(),
		// clone intermediate query.
		sql:  blq.sql.Clone(),
		path: blq.path,
	}
}

// WithDonor tells the query-builder to eager-load the nodes that are connected to
// the "donor" edge. The optional arguments are used to configure the query builder of the edge.
func (blq *BalanceLogQuery) WithDonor(opts ...func(*UserQuery)) *BalanceLogQuery {
	query := (&UserClient{config: blq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	blq.withDonor = query
	return blq
}

// WithReceiver tells the query-builder to eager-load the nodes that are connected to
// the "receiver" edge. The optional arguments are used to configure the query builder of the edge.
func (blq *BalanceLogQuery) WithReceiver(opts ...func(*UserQuery)) *BalanceLogQuery {
	query := (&UserClient{config: blq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	blq.withReceiver = query
	return blq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Price int `json:"price,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BalanceLog.Query().
//		GroupBy(balancelog.FieldPrice).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (blq *BalanceLogQuery) GroupBy(field string, fields ...string) *BalanceLogGroupBy {
	blq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BalanceLogGroupBy{build: blq}
	grbuild.flds = &blq.ctx.Fields
	grbuild.label = balancelog.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Price int `json:"price,omitempty"`
//	}
//
//	client.BalanceLog.Query().
//		Select(balancelog.FieldPrice).
//		Scan(ctx, &v)
func (blq *BalanceLogQuery) Select(fields ...string) *BalanceLogSelect {
	blq.ctx.Fields = append(blq.ctx.Fields, fields...)
	sbuild := &BalanceLogSelect{BalanceLogQuery: blq}
	sbuild.label = balancelog.Label
	sbuild.flds, sbuild.scan = &blq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BalanceLogSelect configured with the given aggregations.
func (blq *BalanceLogQuery) Aggregate(fns ...AggregateFunc) *BalanceLogSelect {
	return blq.Select().Aggregate(fns...)
}

func (blq *BalanceLogQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range blq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, blq); err != nil {
				return err
			}
		}
	}
	for _, f := range blq.ctx.Fields {
		if !balancelog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if blq.path != nil {
		prev, err := blq.path(ctx)
		if err != nil {
			return err
		}
		blq.sql = prev
	}
	return nil
}

func (blq *BalanceLogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BalanceLog, error) {
	var (
		nodes       = []*BalanceLog{}
		withFKs     = blq.withFKs
		_spec       = blq.querySpec()
		loadedTypes = [2]bool{
			blq.withDonor != nil,
			blq.withReceiver != nil,
		}
	)
	if blq.withDonor != nil || blq.withReceiver != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, balancelog.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BalanceLog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BalanceLog{config: blq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, blq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := blq.withDonor; query != nil {
		if err := blq.loadDonor(ctx, query, nodes, nil,
			func(n *BalanceLog, e *User) { n.Edges.Donor = e }); err != nil {
			return nil, err
		}
	}
	if query := blq.withReceiver; query != nil {
		if err := blq.loadReceiver(ctx, query, nodes, nil,
			func(n *BalanceLog, e *User) { n.Edges.Receiver = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (blq *BalanceLogQuery) loadDonor(ctx context.Context, query *UserQuery, nodes []*BalanceLog, init func(*BalanceLog), assign func(*BalanceLog, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*BalanceLog)
	for i := range nodes {
		if nodes[i].user_donor == nil {
			continue
		}
		fk := *nodes[i].user_donor
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_donor" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (blq *BalanceLogQuery) loadReceiver(ctx context.Context, query *UserQuery, nodes []*BalanceLog, init func(*BalanceLog), assign func(*BalanceLog, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*BalanceLog)
	for i := range nodes {
		if nodes[i].user_receiver == nil {
			continue
		}
		fk := *nodes[i].user_receiver
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_receiver" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (blq *BalanceLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := blq.querySpec()
	_spec.Node.Columns = blq.ctx.Fields
	if len(blq.ctx.Fields) > 0 {
		_spec.Unique = blq.ctx.Unique != nil && *blq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, blq.driver, _spec)
}

func (blq *BalanceLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(balancelog.Table, balancelog.Columns, sqlgraph.NewFieldSpec(balancelog.FieldID, field.TypeInt))
	_spec.From = blq.sql
	if unique := blq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if blq.path != nil {
		_spec.Unique = true
	}
	if fields := blq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, balancelog.FieldID)
		for i := range fields {
			if fields[i] != balancelog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := blq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := blq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := blq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := blq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (blq *BalanceLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(blq.driver.Dialect())
	t1 := builder.Table(balancelog.Table)
	columns := blq.ctx.Fields
	if len(columns) == 0 {
		columns = balancelog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if blq.sql != nil {
		selector = blq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if blq.ctx.Unique != nil && *blq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range blq.predicates {
		p(selector)
	}
	for _, p := range blq.order {
		p(selector)
	}
	if offset := blq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := blq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BalanceLogGroupBy is the group-by builder for BalanceLog entities.
type BalanceLogGroupBy struct {
	selector
	build *BalanceLogQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (blgb *BalanceLogGroupBy) Aggregate(fns ...AggregateFunc) *BalanceLogGroupBy {
	blgb.fns = append(blgb.fns, fns...)
	return blgb
}

// Scan applies the selector query and scans the result into the given value.
func (blgb *BalanceLogGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, blgb.build.ctx, "GroupBy")
	if err := blgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BalanceLogQuery, *BalanceLogGroupBy](ctx, blgb.build, blgb, blgb.build.inters, v)
}

func (blgb *BalanceLogGroupBy) sqlScan(ctx context.Context, root *BalanceLogQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(blgb.fns))
	for _, fn := range blgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*blgb.flds)+len(blgb.fns))
		for _, f := range *blgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*blgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := blgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BalanceLogSelect is the builder for selecting fields of BalanceLog entities.
type BalanceLogSelect struct {
	*BalanceLogQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bls *BalanceLogSelect) Aggregate(fns ...AggregateFunc) *BalanceLogSelect {
	bls.fns = append(bls.fns, fns...)
	return bls
}

// Scan applies the selector query and scans the result into the given value.
func (bls *BalanceLogSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bls.ctx, "Select")
	if err := bls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BalanceLogQuery, *BalanceLogSelect](ctx, bls.BalanceLogQuery, bls, bls.inters, v)
}

func (bls *BalanceLogSelect) sqlScan(ctx context.Context, root *BalanceLogQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bls.fns))
	for _, fn := range bls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
