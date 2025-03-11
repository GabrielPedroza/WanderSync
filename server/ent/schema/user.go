package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/schema"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").Default("Unknown"),
		field.Int("age").Positive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	// return []ent.Edge{
	//   edge.From()
	// }
	return nil
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
