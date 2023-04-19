package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Grocery holds the schema definition for the Grocery entity.
type Grocery struct {
	ent.Schema
}

// Fields of the Grocery.
func (Grocery) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Int("price").NonNegative().StructTag(`json:"price"`),
		field.Int("unit").NonNegative().StructTag(`json:"unit"`),
		field.Time("expiration_date").Default(time.Now().Add(time.Hour * 24 * 365)), // almost a year
		field.Time("created_at").Default(time.Now()).Immutable(),
	}
}

// Edges of the Grocery.
func (Grocery) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("provider", User.Type).Ref("provided_groceries").Unique(),
		edge.To("purchased", Purchase.Type),
	}
}
