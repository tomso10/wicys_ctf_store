package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the Users entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Comment("username of user").
			NotEmpty().
			Unique(),
		field.String("hash").
			Comment("password hash of user").
			Sensitive().
			NotEmpty().
			Unique(),
		field.Enum("permissions").
			Comment("permissions of user").
			Values("blue", "red", "white", "black"),
		field.Int("balance").
			Comment("balance of user").
			NonNegative().
			Default(0),
	}
}

// Edges of the Users.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("transactions", Transaction.Type),
		edge.From("tokens", Token.Type).
			Ref("user"),
	}
}
