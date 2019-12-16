package collections

import "strconv"

// IntValue type alias for int
type IntValue int

// Int casts and returns int value
func (i IntValue) Int() int {
	return int(i)
}

// Int32 casts and returns Int32 value
func (i IntValue) Int32() int32 {
	return int32(i)
}

// Int64 casts and returns Int64 value
func (i IntValue) Int64() int64 {
	return int64(i)
}

// Float32 casts and returns Float32 value
func (i IntValue) Float32() float32 {
	return float32(i)
}

// Float64 casts and returns Float64 value
func (i IntValue) Float64() float64 {
	return float64(i)
}

// String casts and returns string value
func (i IntValue) String() string {
	return strconv.Itoa(int(i))
}

// IntValues type alias for a slice of IntValue
type IntValues []int

func (i IntValues) Data() []Data {
	d := make([]Data, len(i))
	for j := range i {
		d[j] = IntValue(i[j])
	}
	return d
}

// IntValue32 type alias for int
type IntValue32 int32

// Int casts and returns int value
func (i IntValue32) Int() int {
	return int(i)
}

// Int32 casts and returns Int32 value
func (i IntValue32) Int32() int32 {
	return int32(i)
}

// Int64 casts and returns Int64 value
func (i IntValue32) Int64() int64 {
	return int64(i)
}

// Float32 casts and returns Float32 value
func (i IntValue32) Float32() float32 {
	return float32(i)
}

// Float64 casts and returns Float64 value
func (i IntValue32) Float64() float64 {
	return float64(i)
}

// String casts and returns string value
func (i IntValue32) String() string {
	return strconv.Itoa(int(i))
}

// IntValues32 type alias for a slice of IntValue
type IntValues32 []int32

func (i IntValues32) Data() []Data {
	d := make([]Data, len(i))
	for j := range i {
		d[j] = IntValue32(i[j])
	}
	return d
}

// IntValue64 type alias for int
type IntValue64 int64

// Int casts and returns int value
func (i IntValue64) Int() int {
	return int(i)
}

// Int32 casts and returns Int32 value
func (i IntValue64) Int32() int32 {
	return int32(i)
}

// Int64 casts and returns Int64 value
func (i IntValue64) Int64() int64 {
	return int64(i)
}

// Float32 casts and returns Float32 value
func (i IntValue64) Float32() float32 {
	return float32(i)
}

// Float64 casts and returns Float64 value
func (i IntValue64) Float64() float64 {
	return float64(i)
}

// String casts and returns string value
func (i IntValue64) String() string {
	return strconv.Itoa(int(i))
}

// IntValues64 type alias for a slice of IntValue
type IntValues64 []int64

func (i IntValues64) Data() []Data {
	d := make([]Data, len(i))
	for j := range i {
		d[j] = IntValue64(i[j])
	}
	return d
}

// StringValue type alias for string
type StringValue string

// Int casts and returns int value
func (s StringValue) Int() int {
	n, _ := strconv.Atoi(string(s))
	return n
}

// Int32 casts and returns Int32 value
func (s StringValue) Int32() int32 {
	n, _ := strconv.Atoi(string(s))
	return int32(n)
}

// Int64 casts and returns Int64 value
func (s StringValue) Int64() int64 {
	n, _ := strconv.Atoi(string(s))
	return int64(n)
}

// Float32 casts and returns Float32 value
func (s StringValue) Float32() float32 {
	n, _ := strconv.ParseFloat(string(s), 32)
	return float32(n)
}

// Float64 casts and returns Float64 value
func (s StringValue) Float64() float64 {
	n, _ := strconv.ParseFloat(string(s), 64)
	return n
}

// String casts and returns string value
func (s StringValue) String() string {
	return string(s)
}

// StringValues type alias for a slice of StringValue
type StringValues []StringValue

func (s StringValues) Data() []Data {
	d := make([]Data, len(s))
	for j := range s {
		d[j] = StringValue(s[j])
	}
	return d
}
