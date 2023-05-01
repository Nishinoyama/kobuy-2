// Code generated by ent, DO NOT EDIT.

package purchase

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/nishinoyama/kobuy-2/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Purchase {
	return predicate.Purchase(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Purchase {
	return predicate.Purchase(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Purchase {
	return predicate.Purchase(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Purchase {
	return predicate.Purchase(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Purchase {
	return predicate.Purchase(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Purchase {
	return predicate.Purchase(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Purchase {
	return predicate.Purchase(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Purchase {
	return predicate.Purchase(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Purchase {
	return predicate.Purchase(sql.FieldLTE(FieldID, id))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v int) predicate.Purchase {
	return predicate.Purchase(sql.FieldEQ(FieldPrice, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v int) predicate.Purchase {
	return predicate.Purchase(sql.FieldEQ(FieldAmount, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Purchase {
	return predicate.Purchase(sql.FieldEQ(FieldCreatedAt, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v int) predicate.Purchase {
	return predicate.Purchase(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v int) predicate.Purchase {
	return predicate.Purchase(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...int) predicate.Purchase {
	return predicate.Purchase(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...int) predicate.Purchase {
	return predicate.Purchase(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v int) predicate.Purchase {
	return predicate.Purchase(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v int) predicate.Purchase {
	return predicate.Purchase(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v int) predicate.Purchase {
	return predicate.Purchase(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v int) predicate.Purchase {
	return predicate.Purchase(sql.FieldLTE(FieldPrice, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v int) predicate.Purchase {
	return predicate.Purchase(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v int) predicate.Purchase {
	return predicate.Purchase(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...int) predicate.Purchase {
	return predicate.Purchase(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...int) predicate.Purchase {
	return predicate.Purchase(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v int) predicate.Purchase {
	return predicate.Purchase(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v int) predicate.Purchase {
	return predicate.Purchase(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v int) predicate.Purchase {
	return predicate.Purchase(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v int) predicate.Purchase {
	return predicate.Purchase(sql.FieldLTE(FieldAmount, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Purchase {
	return predicate.Purchase(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Purchase {
	return predicate.Purchase(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Purchase {
	return predicate.Purchase(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Purchase {
	return predicate.Purchase(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Purchase {
	return predicate.Purchase(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Purchase {
	return predicate.Purchase(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Purchase {
	return predicate.Purchase(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Purchase {
	return predicate.Purchase(sql.FieldLTE(FieldCreatedAt, v))
}

// HasBuyer applies the HasEdge predicate on the "buyer" edge.
func HasBuyer() predicate.Purchase {
	return predicate.Purchase(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BuyerTable, BuyerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBuyerWith applies the HasEdge predicate on the "buyer" edge with a given conditions (other predicates).
func HasBuyerWith(preds ...predicate.User) predicate.Purchase {
	return predicate.Purchase(func(s *sql.Selector) {
		step := newBuyerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGrocery applies the HasEdge predicate on the "grocery" edge.
func HasGrocery() predicate.Purchase {
	return predicate.Purchase(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GroceryTable, GroceryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroceryWith applies the HasEdge predicate on the "grocery" edge with a given conditions (other predicates).
func HasGroceryWith(preds ...predicate.Grocery) predicate.Purchase {
	return predicate.Purchase(func(s *sql.Selector) {
		step := newGroceryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Purchase) predicate.Purchase {
	return predicate.Purchase(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Purchase) predicate.Purchase {
	return predicate.Purchase(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Purchase) predicate.Purchase {
	return predicate.Purchase(func(s *sql.Selector) {
		p(s.Not())
	})
}
