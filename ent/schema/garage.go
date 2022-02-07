// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Garage holds the schema definition for the Garage entity.
type Garage struct {
	ent.Schema
}

// Fields of the Garage.
func (Garage) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

func (Garage) Edges() []ent.Edge {
	return []ent.Edge{
		// not sure
	}
}
