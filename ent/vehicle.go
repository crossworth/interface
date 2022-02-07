// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/bug/ent/garage"
	"entgo.io/bug/ent/vehicle"
	"entgo.io/ent/dialect/sql"
)

// Vehicle is the model entity for the Vehicle schema.
type Vehicle struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// GarageID holds the value of the "garage_id" field.
	GarageID *int `json:"garage_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VehicleQuery when eager-loading is set.
	Edges VehicleEdges `json:"edges"`
}

// VehicleEdges holds the relations/edges for other nodes in the graph.
type VehicleEdges struct {
	// Garage holds the value of the garage edge.
	Garage *Garage `json:"garage,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// GarageOrErr returns the Garage value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e VehicleEdges) GarageOrErr() (*Garage, error) {
	if e.loadedTypes[0] {
		if e.Garage == nil {
			// The edge garage was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: garage.Label}
		}
		return e.Garage, nil
	}
	return nil, &NotLoadedError{edge: "garage"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Vehicle) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case vehicle.FieldGarageID:
			values[i] = new(sql.NullInt64)
		case vehicle.FieldID, vehicle.FieldType, vehicle.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Vehicle", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Vehicle fields.
func (v *Vehicle) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vehicle.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				v.ID = value.String
			}
		case vehicle.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				v.Type = value.String
			}
		case vehicle.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				v.Name = value.String
			}
		case vehicle.FieldGarageID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field garage_id", values[i])
			} else if value.Valid {
				v.GarageID = new(int)
				*v.GarageID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryGarage queries the "garage" edge of the Vehicle entity.
func (v *Vehicle) QueryGarage() *GarageQuery {
	return (&VehicleClient{config: v.config}).QueryGarage(v)
}

// Update returns a builder for updating this Vehicle.
// Note that you need to call Vehicle.Unwrap() before calling this method if this Vehicle
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Vehicle) Update() *VehicleUpdateOne {
	return (&VehicleClient{config: v.config}).UpdateOne(v)
}

// Unwrap unwraps the Vehicle entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Vehicle) Unwrap() *Vehicle {
	tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: Vehicle is not a transactional entity")
	}
	v.config.driver = tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Vehicle) String() string {
	var builder strings.Builder
	builder.WriteString("Vehicle(")
	builder.WriteString(fmt.Sprintf("id=%v", v.ID))
	builder.WriteString(", type=")
	builder.WriteString(v.Type)
	builder.WriteString(", name=")
	builder.WriteString(v.Name)
	if v := v.GarageID; v != nil {
		builder.WriteString(", garage_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Vehicles is a parsable slice of Vehicle.
type Vehicles []*Vehicle

func (v Vehicles) config(cfg config) {
	for _i := range v {
		v[_i].config = cfg
	}
}
