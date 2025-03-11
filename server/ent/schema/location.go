package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/schema"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Location holds the schema definition for the Location entity.
type Location struct {
	ent.Schema
}

// Fields of the Location.
func (Location) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").Default("Unknown Location"),
	}
}

// Edges of the Location.
func (Location) Edges() []ent.Edge {
	return nil
}

func (Location) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
