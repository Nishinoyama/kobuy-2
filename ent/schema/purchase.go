package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Purchase holds the schema definition for the Purchase entity.
type Purchase struct {
	ent.Schema
}

// Fields of the Purchase.
func (Purchase) Fields() []ent.Field {
	return []ent.Field{
		field.Int("price").NonNegative(),
		field.Int("amount").NonNegative(),
		field.Time("created_at").Default(time.Now()).Immutable(),
	}
}

// Edges of the Purchase.
func (Purchase) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("buyer", User.Type).Ref("purchased").Unique(),
		edge.From("grocery", Grocery.Type).Ref("purchased").Unique(),
	}
}
