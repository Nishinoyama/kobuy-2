package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("user_id").Unique(),
		field.String("password").Sensitive(),
		field.Int("balance").Default(0).StructTag(`json:"balance"`),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("provided_groceries", Grocery.Type),
		edge.To("purchased", Purchase.Type),
		edge.To("payer", Ledger.Type),
		edge.To("receiver", Ledger.Type),
	}
}
