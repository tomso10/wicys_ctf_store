package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Item holds the schema definition for the Item entity.
type Item struct {
	ent.Schema
}

// Fields of the Item.
func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("name of the item").
			NotEmpty().
			Unique(),
		field.String("description").
			Comment("description of the item").
			NotEmpty(),
		field.String("image").
			Comment("image of file").
			NotEmpty(),
		field.Int("price").
			Comment("price of the item").
			NonNegative(),
	}
}

// Edges of the Item.
func (Item) Edges() []ent.Edge {
	return nil
}
