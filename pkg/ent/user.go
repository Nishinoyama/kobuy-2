// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/nishinoyama/kobuy-2/pkg/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// Balance holds the value of the "balance" field.
	Balance int `json:"balance"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// ProvidedGroceries holds the value of the provided_groceries edge.
	ProvidedGroceries []*Grocery `json:"provided_groceries,omitempty"`
	// Purchased holds the value of the purchased edge.
	Purchased []*Purchase `json:"purchased,omitempty"`
	// Payer holds the value of the payer edge.
	Payer []*Ledger `json:"payer,omitempty"`
	// Receiver holds the value of the receiver edge.
	Receiver []*Ledger `json:"receiver,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ProvidedGroceriesOrErr returns the ProvidedGroceries value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ProvidedGroceriesOrErr() ([]*Grocery, error) {
	if e.loadedTypes[0] {
		return e.ProvidedGroceries, nil
	}
	return nil, &NotLoadedError{edge: "provided_groceries"}
}

// PurchasedOrErr returns the Purchased value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PurchasedOrErr() ([]*Purchase, error) {
	if e.loadedTypes[1] {
		return e.Purchased, nil
	}
	return nil, &NotLoadedError{edge: "purchased"}
}

// PayerOrErr returns the Payer value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PayerOrErr() ([]*Ledger, error) {
	if e.loadedTypes[2] {
		return e.Payer, nil
	}
	return nil, &NotLoadedError{edge: "payer"}
}

// ReceiverOrErr returns the Receiver value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ReceiverOrErr() ([]*Ledger, error) {
	if e.loadedTypes[3] {
		return e.Receiver, nil
	}
	return nil, &NotLoadedError{edge: "receiver"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID, user.FieldBalance:
			values[i] = new(sql.NullInt64)
		case user.FieldName, user.FieldUserID, user.FieldPassword:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				u.UserID = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldBalance:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field balance", values[i])
			} else if value.Valid {
				u.Balance = int(value.Int64)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryProvidedGroceries queries the "provided_groceries" edge of the User entity.
func (u *User) QueryProvidedGroceries() *GroceryQuery {
	return NewUserClient(u.config).QueryProvidedGroceries(u)
}

// QueryPurchased queries the "purchased" edge of the User entity.
func (u *User) QueryPurchased() *PurchaseQuery {
	return NewUserClient(u.config).QueryPurchased(u)
}

// QueryPayer queries the "payer" edge of the User entity.
func (u *User) QueryPayer() *LedgerQuery {
	return NewUserClient(u.config).QueryPayer(u)
}

// QueryReceiver queries the "receiver" edge of the User entity.
func (u *User) QueryReceiver() *LedgerQuery {
	return NewUserClient(u.config).QueryReceiver(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(u.UserID)
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("balance=")
	builder.WriteString(fmt.Sprintf("%v", u.Balance))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
