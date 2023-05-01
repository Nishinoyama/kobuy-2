// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/nishinoyama/kobuy-2/pkg/ent/grocery"
	"github.com/nishinoyama/kobuy-2/pkg/ent/purchase"
	"github.com/nishinoyama/kobuy-2/pkg/ent/schema"
	"github.com/nishinoyama/kobuy-2/pkg/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	groceryFields := schema.Grocery{}.Fields()
	_ = groceryFields
	// groceryDescName is the schema descriptor for name field.
	groceryDescName := groceryFields[0].Descriptor()
	// grocery.NameValidator is a validator for the "name" field. It is called by the builders before save.
	grocery.NameValidator = groceryDescName.Validators[0].(func(string) error)
	// groceryDescPrice is the schema descriptor for price field.
	groceryDescPrice := groceryFields[1].Descriptor()
	// grocery.PriceValidator is a validator for the "price" field. It is called by the builders before save.
	grocery.PriceValidator = groceryDescPrice.Validators[0].(func(int) error)
	// groceryDescUnit is the schema descriptor for unit field.
	groceryDescUnit := groceryFields[2].Descriptor()
	// grocery.UnitValidator is a validator for the "unit" field. It is called by the builders before save.
	grocery.UnitValidator = groceryDescUnit.Validators[0].(func(int) error)
	// groceryDescExpirationDate is the schema descriptor for expiration_date field.
	groceryDescExpirationDate := groceryFields[3].Descriptor()
	// grocery.DefaultExpirationDate holds the default value on creation for the expiration_date field.
	grocery.DefaultExpirationDate = groceryDescExpirationDate.Default.(time.Time)
	// groceryDescCreatedAt is the schema descriptor for created_at field.
	groceryDescCreatedAt := groceryFields[4].Descriptor()
	// grocery.DefaultCreatedAt holds the default value on creation for the created_at field.
	grocery.DefaultCreatedAt = groceryDescCreatedAt.Default.(time.Time)
	purchaseFields := schema.Purchase{}.Fields()
	_ = purchaseFields
	// purchaseDescPrice is the schema descriptor for price field.
	purchaseDescPrice := purchaseFields[0].Descriptor()
	// purchase.PriceValidator is a validator for the "price" field. It is called by the builders before save.
	purchase.PriceValidator = purchaseDescPrice.Validators[0].(func(int) error)
	// purchaseDescAmount is the schema descriptor for amount field.
	purchaseDescAmount := purchaseFields[1].Descriptor()
	// purchase.AmountValidator is a validator for the "amount" field. It is called by the builders before save.
	purchase.AmountValidator = purchaseDescAmount.Validators[0].(func(int) error)
	// purchaseDescCreatedAt is the schema descriptor for created_at field.
	purchaseDescCreatedAt := purchaseFields[2].Descriptor()
	// purchase.DefaultCreatedAt holds the default value on creation for the created_at field.
	purchase.DefaultCreatedAt = purchaseDescCreatedAt.Default.(time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescBalance is the schema descriptor for balance field.
	userDescBalance := userFields[3].Descriptor()
	// user.DefaultBalance holds the default value on creation for the balance field.
	user.DefaultBalance = userDescBalance.Default.(int)
}
