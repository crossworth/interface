// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/bug/ent/garage"
	"entgo.io/ent/dialect/sql"
)

// Garage is the model entity for the Garage schema.
type Garage struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Garage) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case garage.FieldID, garage.FieldType, garage.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Garage", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Garage fields.
func (ga *Garage) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case garage.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ga.ID = value.String
			}
		case garage.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				ga.Type = value.String
			}
		case garage.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ga.Name = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Garage.
// Note that you need to call Garage.Unwrap() before calling this method if this Garage
// was returned from a transaction, and the transaction was committed or rolled back.
func (ga *Garage) Update() *GarageUpdateOne {
	return (&GarageClient{config: ga.config}).UpdateOne(ga)
}

// Unwrap unwraps the Garage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ga *Garage) Unwrap() *Garage {
	tx, ok := ga.config.driver.(*txDriver)
	if !ok {
		panic("ent: Garage is not a transactional entity")
	}
	ga.config.driver = tx.drv
	return ga
}

// String implements the fmt.Stringer.
func (ga *Garage) String() string {
	var builder strings.Builder
	builder.WriteString("Garage(")
	builder.WriteString(fmt.Sprintf("id=%v", ga.ID))
	builder.WriteString(", type=")
	builder.WriteString(ga.Type)
	builder.WriteString(", name=")
	builder.WriteString(ga.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Garages is a parsable slice of Garage.
type Garages []*Garage

func (ga Garages) config(cfg config) {
	for _i := range ga {
		ga[_i].config = cfg
	}
}