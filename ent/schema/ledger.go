package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Ledger holds the schema definition for the Ledger entity.
type Ledger struct {
	ent.Schema
}

// Fields of the Ledger.
func (Ledger) Fields() []ent.Field {
	return []ent.Field{
		field.Int("price"),
		field.Enum("type").Values("purchase", "cash", "etc"),
	}
}

// Edges of the Ledger.
func (Ledger) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("donor", User.Type).Ref("donor").Unique(),
		edge.From("receiver", User.Type).Ref("receiver").Unique(),
	}
}
