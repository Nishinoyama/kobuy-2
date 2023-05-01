// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldBalance holds the string denoting the balance field in the database.
	FieldBalance = "balance"
	// EdgeProvidedGroceries holds the string denoting the provided_groceries edge name in mutations.
	EdgeProvidedGroceries = "provided_groceries"
	// EdgePurchased holds the string denoting the purchased edge name in mutations.
	EdgePurchased = "purchased"
	// EdgePayer holds the string denoting the payer edge name in mutations.
	EdgePayer = "payer"
	// EdgeReceiver holds the string denoting the receiver edge name in mutations.
	EdgeReceiver = "receiver"
	// Table holds the table name of the user in the database.
	Table = "users"
	// ProvidedGroceriesTable is the table that holds the provided_groceries relation/edge.
	ProvidedGroceriesTable = "groceries"
	// ProvidedGroceriesInverseTable is the table name for the Grocery entity.
	// It exists in this package in order to avoid circular dependency with the "grocery" package.
	ProvidedGroceriesInverseTable = "groceries"
	// ProvidedGroceriesColumn is the table column denoting the provided_groceries relation/edge.
	ProvidedGroceriesColumn = "user_provided_groceries"
	// PurchasedTable is the table that holds the purchased relation/edge.
	PurchasedTable = "purchases"
	// PurchasedInverseTable is the table name for the Purchase entity.
	// It exists in this package in order to avoid circular dependency with the "purchase" package.
	PurchasedInverseTable = "purchases"
	// PurchasedColumn is the table column denoting the purchased relation/edge.
	PurchasedColumn = "user_purchased"
	// PayerTable is the table that holds the payer relation/edge.
	PayerTable = "ledgers"
	// PayerInverseTable is the table name for the Ledger entity.
	// It exists in this package in order to avoid circular dependency with the "ledger" package.
	PayerInverseTable = "ledgers"
	// PayerColumn is the table column denoting the payer relation/edge.
	PayerColumn = "user_payer"
	// ReceiverTable is the table that holds the receiver relation/edge.
	ReceiverTable = "ledgers"
	// ReceiverInverseTable is the table name for the Ledger entity.
	// It exists in this package in order to avoid circular dependency with the "ledger" package.
	ReceiverInverseTable = "ledgers"
	// ReceiverColumn is the table column denoting the receiver relation/edge.
	ReceiverColumn = "user_receiver"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldUserID,
	FieldPassword,
	FieldBalance,
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
	// DefaultBalance holds the default value on creation for the "balance" field.
	DefaultBalance int
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByBalance orders the results by the balance field.
func ByBalance(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBalance, opts...).ToFunc()
}

// ByProvidedGroceriesCount orders the results by provided_groceries count.
func ByProvidedGroceriesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProvidedGroceriesStep(), opts...)
	}
}

// ByProvidedGroceries orders the results by provided_groceries terms.
func ByProvidedGroceries(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProvidedGroceriesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPurchasedCount orders the results by purchased count.
func ByPurchasedCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPurchasedStep(), opts...)
	}
}

// ByPurchased orders the results by purchased terms.
func ByPurchased(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPurchasedStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPayerCount orders the results by payer count.
func ByPayerCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPayerStep(), opts...)
	}
}

// ByPayer orders the results by payer terms.
func ByPayer(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPayerStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByReceiverCount orders the results by receiver count.
func ByReceiverCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newReceiverStep(), opts...)
	}
}

// ByReceiver orders the results by receiver terms.
func ByReceiver(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReceiverStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProvidedGroceriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProvidedGroceriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProvidedGroceriesTable, ProvidedGroceriesColumn),
	)
}
func newPurchasedStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PurchasedInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PurchasedTable, PurchasedColumn),
	)
}
func newPayerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PayerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PayerTable, PayerColumn),
	)
}
func newReceiverStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReceiverInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ReceiverTable, ReceiverColumn),
	)
}
