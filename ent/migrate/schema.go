// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GroceriesColumns holds the columns for the "groceries" table.
	GroceriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "price", Type: field.TypeInt},
		{Name: "unit", Type: field.TypeInt},
		{Name: "expiration_date", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_provided_groceries", Type: field.TypeInt, Nullable: true},
	}
	// GroceriesTable holds the schema information for the "groceries" table.
	GroceriesTable = &schema.Table{
		Name:       "groceries",
		Columns:    GroceriesColumns,
		PrimaryKey: []*schema.Column{GroceriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "groceries_users_provided_groceries",
				Columns:    []*schema.Column{GroceriesColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PurchasesColumns holds the columns for the "purchases" table.
	PurchasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeInt},
		{Name: "amount", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "grocery_purchased", Type: field.TypeInt, Nullable: true},
		{Name: "user_purchased", Type: field.TypeInt, Nullable: true},
	}
	// PurchasesTable holds the schema information for the "purchases" table.
	PurchasesTable = &schema.Table{
		Name:       "purchases",
		Columns:    PurchasesColumns,
		PrimaryKey: []*schema.Column{PurchasesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "purchases_groceries_purchased",
				Columns:    []*schema.Column{PurchasesColumns[4]},
				RefColumns: []*schema.Column{GroceriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "purchases_users_purchased",
				Columns:    []*schema.Column{PurchasesColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GroceriesTable,
		PurchasesTable,
		UsersTable,
	}
)

func init() {
	GroceriesTable.ForeignKeys[0].RefTable = UsersTable
	PurchasesTable.ForeignKeys[0].RefTable = GroceriesTable
	PurchasesTable.ForeignKeys[1].RefTable = UsersTable
}
