// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("garage_id").Optional().Nillable(),
	}
}

func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("garage", Garage.Type).Field("garage_id").Unique(),
	}
}
