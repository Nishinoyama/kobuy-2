// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/nishinoyama/kobuy-2/ent/grocery"
	"github.com/nishinoyama/kobuy-2/ent/purchase"
	"github.com/nishinoyama/kobuy-2/ent/user"
)

// Purchase is the model entity for the Purchase schema.
type Purchase struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Price holds the value of the "price" field.
	Price int `json:"price,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount int `json:"amount,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PurchaseQuery when eager-loading is set.
	Edges             PurchaseEdges `json:"edges"`
	grocery_purchased *int
	user_purchased    *int
}

// PurchaseEdges holds the relations/edges for other nodes in the graph.
type PurchaseEdges struct {
	// Buyer holds the value of the buyer edge.
	Buyer *User `json:"buyer,omitempty"`
	// Grocery holds the value of the grocery edge.
	Grocery *Grocery `json:"grocery,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// BuyerOrErr returns the Buyer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PurchaseEdges) BuyerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Buyer == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Buyer, nil
	}
	return nil, &NotLoadedError{edge: "buyer"}
}

// GroceryOrErr returns the Grocery value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PurchaseEdges) GroceryOrErr() (*Grocery, error) {
	if e.loadedTypes[1] {
		if e.Grocery == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: grocery.Label}
		}
		return e.Grocery, nil
	}
	return nil, &NotLoadedError{edge: "grocery"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Purchase) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case purchase.FieldID, purchase.FieldPrice, purchase.FieldAmount:
			values[i] = new(sql.NullInt64)
		case purchase.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case purchase.ForeignKeys[0]: // grocery_purchased
			values[i] = new(sql.NullInt64)
		case purchase.ForeignKeys[1]: // user_purchased
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Purchase", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Purchase fields.
func (pu *Purchase) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case purchase.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pu.ID = int(value.Int64)
		case purchase.FieldPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				pu.Price = int(value.Int64)
			}
		case purchase.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				pu.Amount = int(value.Int64)
			}
		case purchase.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pu.CreatedAt = value.Time
			}
		case purchase.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field grocery_purchased", value)
			} else if value.Valid {
				pu.grocery_purchased = new(int)
				*pu.grocery_purchased = int(value.Int64)
			}
		case purchase.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_purchased", value)
			} else if value.Valid {
				pu.user_purchased = new(int)
				*pu.user_purchased = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryBuyer queries the "buyer" edge of the Purchase entity.
func (pu *Purchase) QueryBuyer() *UserQuery {
	return NewPurchaseClient(pu.config).QueryBuyer(pu)
}

// QueryGrocery queries the "grocery" edge of the Purchase entity.
func (pu *Purchase) QueryGrocery() *GroceryQuery {
	return NewPurchaseClient(pu.config).QueryGrocery(pu)
}

// Update returns a builder for updating this Purchase.
// Note that you need to call Purchase.Unwrap() before calling this method if this Purchase
// was returned from a transaction, and the transaction was committed or rolled back.
func (pu *Purchase) Update() *PurchaseUpdateOne {
	return NewPurchaseClient(pu.config).UpdateOne(pu)
}

// Unwrap unwraps the Purchase entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pu *Purchase) Unwrap() *Purchase {
	_tx, ok := pu.config.driver.(*txDriver)
	if !ok {
		panic("ent: Purchase is not a transactional entity")
	}
	pu.config.driver = _tx.drv
	return pu
}

// String implements the fmt.Stringer.
func (pu *Purchase) String() string {
	var builder strings.Builder
	builder.WriteString("Purchase(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pu.ID))
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", pu.Price))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", pu.Amount))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pu.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Purchases is a parsable slice of Purchase.
type Purchases []*Purchase
