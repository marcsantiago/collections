package collections

import (
	"bytes"
	"strconv"
)

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
