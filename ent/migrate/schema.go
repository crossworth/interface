// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CarsColumns holds the columns for the "cars" table.
	CarsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// CarsTable holds the schema information for the "cars" table.
	CarsTable = &schema.Table{
		Name:       "cars",
		Columns:    CarsColumns,
		PrimaryKey: []*schema.Column{CarsColumns[0]},
	}
	// GaragesColumns holds the columns for the "garages" table.
	GaragesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, SchemaType: map[string]string{"mysql": " varchar(20)"}},
		{Name: "type", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
	}
	// GaragesTable holds the schema information for the "garages" table.
	GaragesTable = &schema.Table{
		Name:       "garages",
		Columns:    GaragesColumns,
		PrimaryKey: []*schema.Column{GaragesColumns[0]},
	}
	// PlanesColumns holds the columns for the "planes" table.
	PlanesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// PlanesTable holds the schema information for the "planes" table.
	PlanesTable = &schema.Table{
		Name:       "planes",
		Columns:    PlanesColumns,
		PrimaryKey: []*schema.Column{PlanesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CarsTable,
		GaragesTable,
		PlanesTable,
	}
)

func init() {
}
