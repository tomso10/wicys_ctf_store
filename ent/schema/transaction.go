package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.Time("time").
			Comment("time of the transcation").
			Default(time.Now).
			Immutable(),
		field.String("instructions").
			Comment("instructions of the transaction"),
		field.Int("price").
			Comment("price of the transaction").
			Immutable(),
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("item", Item.Type).
			Unique().
			Required(),
		edge.From("user", User.Type).
			Ref("transactions").
			Unique().
			Required(),
	}
}
