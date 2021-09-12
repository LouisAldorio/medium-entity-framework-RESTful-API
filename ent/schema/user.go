package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
        field.String("email").
            NotEmpty(),
        field.String("name").
            Default("unknown"),
    }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		//To() function will take two parameter (name or the relation, to where the relation is pointed)
		edge.To("todos", Todo.Type),
	}
}
