// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BalanceLogsColumns holds the columns for the "balance_logs" table.
	BalanceLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeInt},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"purchase", "cash", "etc"}},
		{Name: "user_donor", Type: field.TypeInt, Nullable: true},
		{Name: "user_receiver", Type: field.TypeInt, Nullable: true},
	}
	// BalanceLogsTable holds the schema information for the "balance_logs" table.
	BalanceLogsTable = &schema.Table{
		Name:       "balance_logs",
		Columns:    BalanceLogsColumns,
		PrimaryKey: []*schema.Column{BalanceLogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "balance_logs_users_donor",
				Columns:    []*schema.Column{BalanceLogsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "balance_logs_users_receiver",
				Columns:    []*schema.Column{BalanceLogsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GroceriesColumns holds the columns for the "groceries" table.
	GroceriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
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
		{Name: "balance", Type: field.TypeInt, Default: 0},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BalanceLogsTable,
		GroceriesTable,
		PurchasesTable,
		UsersTable,
	}
)

func init() {
	BalanceLogsTable.ForeignKeys[0].RefTable = UsersTable
	BalanceLogsTable.ForeignKeys[1].RefTable = UsersTable
	GroceriesTable.ForeignKeys[0].RefTable = UsersTable
	PurchasesTable.ForeignKeys[0].RefTable = GroceriesTable
	PurchasesTable.ForeignKeys[1].RefTable = UsersTable
}
