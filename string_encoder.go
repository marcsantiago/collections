package collections

import (
	"bytes"
	"reflect"
	"strconv"
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

// StringEncoder writes data into a bytes buffer, used in the String method
func StringEncoder(encoder *bytes.Buffer, data Data, t Type) {
	// rip for optimization
	const size = 0
	b := make([]byte, size)
	switch t {
	case IntType, Int32Type, Int64Type:
		encoder.Write(strconv.AppendInt(b, data.Int64(), 10))
	case Float32Type:
		encoder.Write(strconv.AppendFloat(b, data.Float64(), 'f', -1, 32))
	case Float64Type:
		encoder.Write(strconv.AppendFloat(b, data.Float64(), 'f', -1, 64))
	case StringType:
		encoder.Write([]byte("\"" + data.String() + "\""))
	}
}
