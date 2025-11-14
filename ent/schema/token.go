package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Token holds the schema definition for the Token entity.
type Token struct {
	ent.Schema
}

// Fields of the Token.
func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.String("value").
			Comment("value of token").
			NotEmpty().
			Unique(),
		field.Enum("type").
			Comment("type of token").
			Values(
				"purple",
				"koth",
				"sponsor",
				"ctf",
			),
	}
}

// Edges of the Token.
func (Token) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type),
	}
}
