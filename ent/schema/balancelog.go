package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BalanceLog holds the schema definition for the BalanceLog entity.
type BalanceLog struct {
	ent.Schema
}

// Fields of the BalanceLog.
func (BalanceLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int("price"),
		field.Enum("type").Values("purchase", "cash", "etc"),
	}
}

// Edges of the BalanceLog.
func (BalanceLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("donor", User.Type).Ref("donor").Unique(),
		edge.From("receiver", User.Type).Ref("receiver").Unique(),
	}
}
