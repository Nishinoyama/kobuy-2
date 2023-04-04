package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Grocery holds the schema definition for the Grocery entity.
type Grocery struct {
	ent.Schema
}

// Fields of the Grocery.
func (Grocery) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
	}
}

// Edges of the Grocery.
func (Grocery) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("provider", User.Type).Ref("provided_groceries").Unique(),
	}
}
