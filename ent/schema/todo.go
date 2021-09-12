package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Int("is_done").Default(0),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		//Create an inverse-edge called "owner" of type `User`
		//this is called the backref(Inverse edge)
        // and reference it to the "todos" edge (in User schema)
        // explicitly using the `Ref` method.
		edge.From("owner", User.Type).Ref("todos").
		//Here we are telling th ent framework that a todo can only have one user who owned it 
		Unique(),
	}
}
