// Code generated by entc, DO NOT EDIT.

package car

const (
	// Label holds the string label denoting the car type in the database.
	Label = "car"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldGarageID holds the string denoting the garage_id field in the database.
	FieldGarageID = "garage_id"
	// EdgeGarage holds the string denoting the garage edge name in mutations.
	EdgeGarage = "garage"
	// Table holds the table name of the car in the database.
	Table = "cars"
	// GarageTable is the table that holds the garage relation/edge.
	GarageTable = "cars"
	// GarageInverseTable is the table name for the Garage entity.
	// It exists in this package in order to avoid circular dependency with the "garage" package.
	GarageInverseTable = "garages"
	// GarageColumn is the table column denoting the garage relation/edge.
	GarageColumn = "garage_id"
)

// Columns holds all SQL columns for car fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldGarageID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
