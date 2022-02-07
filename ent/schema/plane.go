// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Plane holds the schema definition for the Plane entity.
type Plane struct {
	ent.Schema
}

// Fields of the Plane.
func (Plane) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}
