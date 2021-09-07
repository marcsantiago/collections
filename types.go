package collections

import "reflect"

// Type is a type alias for an int used for the  "ENUM" definition below
type Type int

// Collection supported types
const (
	UnknownType Type = iota
	IntType
	Int32Type
	Int64Type
	Float32Type
	Float64Type
	StringType

	IntSliceType
	Int32SliceType
	Int64SliceType
	Float32SliceType
	Float64SliceType
	StringSliceType
)

// DetermineDataType gets the internal data type converts it to a supported collections type
// note this does use reflection
func DetermineDataType(data Data) Type {
	switch reflect.TypeOf(data).Kind() {
	case reflect.Int:
		return IntType
	case reflect.Int32:
		return Int32Type
	case reflect.Int64:
		return Int64Type
	case reflect.Float32:
		return Float32Type
	case reflect.Float64:
		return Float64Type
	case reflect.String:
		return StringType
	}
	return UnknownType
}
