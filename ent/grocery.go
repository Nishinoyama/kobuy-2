// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/nishinoyama/kobuy-2/ent/grocery"
	"github.com/nishinoyama/kobuy-2/ent/user"
)

// Grocery is the model entity for the Grocery schema.
type Grocery struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Price holds the value of the "price" field.
	Price int `json:"price,omitempty"`
	// Unit holds the value of the "unit" field.
	Unit int `json:"unit,omitempty"`
	// ExpirationDate holds the value of the "expiration_date" field.
	ExpirationDate time.Time `json:"expiration_date,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroceryQuery when eager-loading is set.
	Edges                   GroceryEdges `json:"edges"`
	user_provided_groceries *int
}

// GroceryEdges holds the relations/edges for other nodes in the graph.
type GroceryEdges struct {
	// Provider holds the value of the provider edge.
	Provider *User `json:"provider,omitempty"`
	// Purchased holds the value of the purchased edge.
	Purchased []*Purchase `json:"purchased,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ProviderOrErr returns the Provider value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroceryEdges) ProviderOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Provider == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Provider, nil
	}
	return nil, &NotLoadedError{edge: "provider"}
}

// PurchasedOrErr returns the Purchased value or an error if the edge
// was not loaded in eager-loading.
func (e GroceryEdges) PurchasedOrErr() ([]*Purchase, error) {
	if e.loadedTypes[1] {
		return e.Purchased, nil
	}
	return nil, &NotLoadedError{edge: "purchased"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Grocery) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case grocery.FieldID, grocery.FieldPrice, grocery.FieldUnit:
			values[i] = new(sql.NullInt64)
		case grocery.FieldName:
			values[i] = new(sql.NullString)
		case grocery.FieldExpirationDate, grocery.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case grocery.ForeignKeys[0]: // user_provided_groceries
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Grocery", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Grocery fields.
func (gr *Grocery) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case grocery.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gr.ID = int(value.Int64)
		case grocery.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				gr.Name = value.String
			}
		case grocery.FieldPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				gr.Price = int(value.Int64)
			}
		case grocery.FieldUnit:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field unit", values[i])
			} else if value.Valid {
				gr.Unit = int(value.Int64)
			}
		case grocery.FieldExpirationDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expiration_date", values[i])
			} else if value.Valid {
				gr.ExpirationDate = value.Time
			}
		case grocery.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gr.CreatedAt = value.Time
			}
		case grocery.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_provided_groceries", value)
			} else if value.Valid {
				gr.user_provided_groceries = new(int)
				*gr.user_provided_groceries = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryProvider queries the "provider" edge of the Grocery entity.
func (gr *Grocery) QueryProvider() *UserQuery {
	return NewGroceryClient(gr.config).QueryProvider(gr)
}

// QueryPurchased queries the "purchased" edge of the Grocery entity.
func (gr *Grocery) QueryPurchased() *PurchaseQuery {
	return NewGroceryClient(gr.config).QueryPurchased(gr)
}

// Update returns a builder for updating this Grocery.
// Note that you need to call Grocery.Unwrap() before calling this method if this Grocery
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *Grocery) Update() *GroceryUpdateOne {
	return NewGroceryClient(gr.config).UpdateOne(gr)
}

// Unwrap unwraps the Grocery entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gr *Grocery) Unwrap() *Grocery {
	_tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Grocery is not a transactional entity")
	}
	gr.config.driver = _tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *Grocery) String() string {
	var builder strings.Builder
	builder.WriteString("Grocery(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gr.ID))
	builder.WriteString("name=")
	builder.WriteString(gr.Name)
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", gr.Price))
	builder.WriteString(", ")
	builder.WriteString("unit=")
	builder.WriteString(fmt.Sprintf("%v", gr.Unit))
	builder.WriteString(", ")
	builder.WriteString("expiration_date=")
	builder.WriteString(gr.ExpirationDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(gr.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Groceries is a parsable slice of Grocery.
type Groceries []*Grocery
