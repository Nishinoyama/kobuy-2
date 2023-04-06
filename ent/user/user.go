// Code generated by ent, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeProvidedGroceries holds the string denoting the provided_groceries edge name in mutations.
	EdgeProvidedGroceries = "provided_groceries"
	// EdgePurchased holds the string denoting the purchased edge name in mutations.
	EdgePurchased = "purchased"
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
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
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
