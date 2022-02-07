// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/bug/ent/car"
	"entgo.io/ent/dialect/sql"
)

// Car is the model entity for the Car schema.
type Car struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Car) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case car.FieldID:
			values[i] = new(sql.NullInt64)
		case car.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Car", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Car fields.
func (c *Car) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case car.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case car.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Car.
// Note that you need to call Car.Unwrap() before calling this method if this Car
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Car) Update() *CarUpdateOne {
	return (&CarClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Car entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Car) Unwrap() *Car {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Car is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Car) String() string {
	var builder strings.Builder
	builder.WriteString("Car(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Cars is a parsable slice of Car.
type Cars []*Car

func (c Cars) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
