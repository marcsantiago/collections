package collections

import (
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

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

// Bool casts and returns bool
func (i IntValue) Bool() bool {
	if i == 0 {
		return false
	}
	return true
}

// Less compares the other data and returns true if self is less than other
func (i IntValue) Less(other Data) bool {
	switch reflect.TypeOf(other).Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		return int(i) < other.Int()
	case reflect.Float32, reflect.Float64:
		return int(i) < int(other.Float64())
	case reflect.String:
		v, _ := strconv.Atoi(other.String())
		return int(i) < v
	}
	return false
}

// Equal compares the other data and returns true is the same as self's value, not self's value and type
func (i IntValue) Equal(other Data) bool {
	switch reflect.TypeOf(other).Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		return int(i) == other.Int()
	case reflect.Float32, reflect.Float64:
		return int(i) == int(other.Float64())
	case reflect.String:
		v, _ := strconv.Atoi(other.String())
		return int(i) == v
	}
	return false
}

// Greater compares the other data and returns true if self is greater than other
func (i IntValue) Greater(other Data) bool {
	if i.Equal(other) {
		return false
	}
	return !i.Less(other)
}

// Add adds OperableData together
func (i IntValue) Add(data OperableData) Data {
	return IntValue(int(i) + data.Int())
}

// Sub subtracts OperableData together
func (i IntValue) Sub(data OperableData) Data {
	return IntValue(int(i) - data.Int())
}

// Mul multiplies OperableData together
func (i IntValue) Mul(data OperableData) Data {
	return IntValue(int(i) * data.Int())
}

// Div divides OperableData together
func (i IntValue) Div(data OperableData) Data {
	return IntValue(int(i) / data.Int())
}

// Mod modulus' OperableData together
func (i IntValue) Mod(data OperableData) Data {
	return IntValue(int(i) % data.Int())
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

// Bool casts and returns bool
func (i IntValue32) Bool() bool {
	if i == 0 {
		return false
	}
	return true
}

// Less compares the other data and returns true is it's less than self
func (i IntValue32) Less(other Data) bool {
	switch reflect.TypeOf(other).Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		return int32(i) < other.Int32()
	case reflect.Float32, reflect.Float64:
		return int32(i) < int32(other.Float64())
	case reflect.String:
		v, _ := strconv.Atoi(other.String())
		return int32(i) < int32(v)
	}
	return false
}

// Equal compares the other data and returns true is the same as self's value, not self's value and type
func (i IntValue32) Equal(other Data) bool {
	switch reflect.TypeOf(other).Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		return int32(i) == other.Int32()
	case reflect.Float32, reflect.Float64:
		return int32(i) == int32(other.Float64())
	case reflect.String:
		v, _ := strconv.Atoi(other.String())
		return int32(i) == int32(v)
	}
	return false
}

// Greater compares the other data and returns true is it's greater than self
func (i IntValue32) Greater(other Data) bool {
	if i.Equal(other) {
		return false
	}
	return !i.Less(other)
}

// Add adds OperableData together
func (i IntValue32) Add(data OperableData) Data {
	return IntValue32(int32(i) + data.Int32())
}

// Sub subtracts OperableData together
func (i IntValue32) Sub(data OperableData) Data {
	return IntValue32(int32(i) - data.Int32())
}

// Mul multiplies OperableData together
func (i IntValue32) Mul(data OperableData) Data {
	return IntValue32(int32(i) * data.Int32())
}

// Div divides OperableData together
func (i IntValue32) Div(data OperableData) Data {
	return IntValue32(int32(i) / data.Int32())
}

// Mod modulus' OperableData together
func (i IntValue32) Mod(data OperableData) Data {
	return IntValue32(int32(i) % data.Int32())
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

// Bool casts and returns bool
func (i IntValue64) Bool() bool {
	if i == 0 {
		return false
	}
	return true
}

// Less compares the other data and returns true is it's less than self
func (i IntValue64) Less(other Data) bool {
	switch reflect.TypeOf(other).Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		return int64(i) < other.Int64()
	case reflect.Float32, reflect.Float64:
		return int64(i) < int64(other.Float64())
	case reflect.String:
		v, _ := strconv.Atoi(other.String())
		return int64(i) < int64(v)
	}
	return false
}

// Equal compares the other data and returns true is the same as self's value, not self's value and type
func (i IntValue64) Equal(other Data) bool {
	switch reflect.TypeOf(other).Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		return int64(i) == other.Int64()
	case reflect.Float32, reflect.Float64:
		return int64(i) == int64(other.Float64())
	case reflect.String:
		v, _ := strconv.Atoi(other.String())
		return int64(i) == int64(v)
	}
	return false
}

// Add adds OperableData together
func (i IntValue64) Add(data OperableData) Data {
	return IntValue64(int64(i) + data.Int64())
}

// Sub subtracts OperableData together
func (i IntValue64) Sub(data OperableData) Data {
	return IntValue64(int64(i) - data.Int64())
}

// Mul multiplies OperableData together
func (i IntValue64) Mul(data OperableData) Data {
	return IntValue64(int64(i) * data.Int64())
}

// Div divides OperableData together
func (i IntValue64) Div(data OperableData) Data {
	return IntValue64(int64(i) / data.Int64())
}

// Mod modulus' OperableData together
func (i IntValue64) Mod(data OperableData) Data {
	return IntValue64(int64(i) % data.Int64())
}

// Greater compares the other data and returns true is it's greater than self
func (i IntValue64) Greater(other Data) bool {
	if i.Equal(other) {
		return false
	}
	return !i.Less(other)
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

// FloatValue32 type alias for int
type FloatValue32 float32

// Int casts and returns int value
func (i FloatValue32) Int() int {
	return int(i)
}

// Int32 casts and returns Int32 value
func (i FloatValue32) Int32() int32 {
	return int32(i)
}

// Int64 casts and returns Int64 value
func (i FloatValue32) Int64() int64 {
	return int64(i)
}

// Float32 casts and returns Float32 value
func (i FloatValue32) Float32() float32 {
	return float32(i)
}

// Float64 casts and returns Float64 value
func (i FloatValue32) Float64() float64 {
	return float64(i)
}

// String casts and returns string value
func (i FloatValue32) String() string {
	return strconv.FormatFloat(float64(i), 'f', -1, 32)
}

// Bool casts and returns bool
func (i FloatValue32) Bool() bool {
	if i == 0.0 {
		return false
	}
	return true
}

// Less compares the other data and returns true is it's less than self
func (i FloatValue32) Less(other Data) bool {
	switch reflect.TypeOf(other).Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		return float32(i) < other.Float32()
	case reflect.Float32, reflect.Float64:
		return float32(i) < other.Float32()
	case reflect.String:
		v, _ := strconv.Atoi(other.String())
		return float32(i) < float32(v)
	}
	return false
}

// Equal compares the other data and returns true is the same as self's value, not self's value and type
func (i FloatValue32) Equal(other Data) bool {
	switch reflect.TypeOf(other).Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		return float32(i) == other.Float32()
	case reflect.Float32, reflect.Float64:
		return float32(i) == other.Float32()
	case reflect.String:
		v, _ := strconv.Atoi(other.String())
		return float32(i) == float32(v)
	}
	return false
}

// Greater compares the other data and returns true is it's greater than self
func (i FloatValue32) Greater(other Data) bool {
	if i.Equal(other) {
		return false
	}
	return !i.Less(other)
}

// Add adds OperableData together
func (i FloatValue32) Add(data OperableData) Data {
	return FloatValue32(float32(i) + data.Float32())
}

// Sub subtracts OperableData together
func (i FloatValue32) Sub(data OperableData) Data {
	return FloatValue32(float32(i) - data.Float32())
}

// Mul multiplies OperableData together
func (i FloatValue32) Mul(data OperableData) Data {
	return FloatValue32(float32(i) * data.Float32())
}

// Div divides OperableData together
func (i FloatValue32) Div(data OperableData) Data {
	return FloatValue32(float32(i) / data.Float32())
}

// Mod modulus' OperableData together, **note** mod only works on integer values to floats are cast to ints first
func (i FloatValue32) Mod(data OperableData) Data {
	return FloatValue32(int(i) % data.Int())
}

// FloatValues32 type alias for a slice of IntValue
type FloatValues32 []float32

func (i FloatValues32) Data() []Data {
	d := make([]Data, len(i))
	for j := range i {
		d[j] = FloatValue32(i[j])
	}
	return d
}

// FloatValue64 type alias for int
type FloatValue64 float64

// Int casts and returns int value
func (i FloatValue64) Int() int {
	return int(i)
}

// Int32 casts and returns Int32 value
func (i FloatValue64) Int32() int32 {
	return int32(i)
}

// Int64 casts and returns Int64 value
func (i FloatValue64) Int64() int64 {
	return int64(i)
}

// Float32 casts and returns Float32 value
func (i FloatValue64) Float32() float32 {
	return float32(i)
}

// Float64 casts and returns Float64 value
func (i FloatValue64) Float64() float64 {
	return float64(i)
}

// String casts and returns string value
func (i FloatValue64) String() string {
	return strconv.FormatFloat(float64(i), 'f', -1, 64)
}

// Bool casts and returns bool
func (i FloatValue64) Bool() bool {
	if i == 0.0 {
		return false
	}
	return true
}

// Less compares the other data and returns true is it's less than self
func (i FloatValue64) Less(other Data) bool {
	switch reflect.TypeOf(other).Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		return float64(i) < other.Float64()
	case reflect.Float32, reflect.Float64:
		return float64(i) < other.Float64()
	case reflect.String:
		v, _ := strconv.Atoi(other.String())
		return float64(i) < float64(v)
	}
	return false
}

// Equal compares the other data and returns true is the same as self's value, not self's value and type
func (i FloatValue64) Equal(other Data) bool {
	switch reflect.TypeOf(other).Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		return float64(i) == other.Float64()
	case reflect.Float32, reflect.Float64:
		return float64(i) == other.Float64()
	case reflect.String:
		v, _ := strconv.Atoi(other.String())
		return float64(i) == float64(v)
	}
	return false
}

// Greater compares the other data and returns true is it's greater than self
func (i FloatValue64) Greater(other Data) bool {
	if i.Equal(other) {
		return false
	}
	return !i.Less(other)
}

// Add adds OperableData together
func (i FloatValue64) Add(data OperableData) Data {
	return FloatValue64(float64(i) + data.Float64())
}

// Sub subtracts OperableData together
func (i FloatValue64) Sub(data OperableData) Data {
	return FloatValue64(float64(i) - data.Float64())
}

// Mul multiplies OperableData together
func (i FloatValue64) Mul(data OperableData) Data {
	return FloatValue64(float64(i) * data.Float64())
}

// Div divides OperableData together
func (i FloatValue64) Div(data OperableData) Data {
	return FloatValue64(float64(i) / data.Float64())
}

// Mod modulus' OperableData together, **note** mod only works on integer values to floats are cast to ints first
func (i FloatValue64) Mod(data OperableData) Data {
	return FloatValue64(int(i) % data.Int())
}

// FloatValues64 type alias for a slice of IntValue
type FloatValues64 []float32

func (i FloatValues64) Data() []Data {
	d := make([]Data, len(i))
	for j := range i {
		d[j] = FloatValue64(i[j])
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

// Bool casts and returns bool
func (s StringValue) Bool() bool {
	switch strings.ToLower(string(s)) {
	case "t", "true":
		return true
	}
	return false
}

// Less compares the other data and returns true is it's less than self
func (s StringValue) Less(other Data) bool {
	v, _ := strconv.Atoi(string(s))
	switch reflect.TypeOf(other).Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		return v < other.Int()
	case reflect.Float32, reflect.Float64:
		return float64(v) < other.Float64()
	case reflect.String:
		return string(s) < other.String()
	}
	return false
}

// Equal compares the other data and returns true is the same as self's value, not self's value and type
func (s StringValue) Equal(other Data) bool {
	v, _ := strconv.Atoi(string(s))
	switch reflect.TypeOf(other).Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		return v == other.Int()
	case reflect.Float32, reflect.Float64:
		return float64(v) == other.Float64()
	case reflect.String:
		return string(s) == other.String()
	}
	return false
}

// Greater compares the other data and returns true is it's greater than self
func (s StringValue) Greater(other Data) bool {
	if s.Equal(other) {
		return false
	}
	return !s.Less(other)
}

// StringValues type alias for a slice of StringValue
type StringValues []string

func (s StringValues) Data() []Data {
	d := make([]Data, len(s))
	for j := range s {
		d[j] = StringValue(s[j])
	}
	return d
}

// RuneValue type alias for rune
type RuneValue rune

// Int casts and returns int value
func (s RuneValue) Int() int {
	return int(s)
}

// Int32 casts and returns Int32 value
func (s RuneValue) Int32() int32 {
	return int32(s)
}

// Int64 casts and returns Int64 value
func (s RuneValue) Int64() int64 {
	return int64(s)
}

// Float32 casts and returns Float32 value
func (s RuneValue) Float32() float32 {
	return float32(s)
}

// Float64 casts and returns Float64 value
func (s RuneValue) Float64() float64 {
	return float64(s)
}

// String casts and returns string value
func (s RuneValue) String() string {
	return string(s)
}

// Bool casts and returns bool
func (s RuneValue) Bool() bool {
	switch unicode.ToLower(rune(s)) {
	case 't':
		return true
	}
	return false
}

// Less compares the other data and returns true is it's less than self
func (s RuneValue) Less(other Data) bool {
	switch reflect.TypeOf(other).Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		return int32(s) < other.Int32()
	case reflect.Float32, reflect.Float64:
		return int32(s) < int32(other.Float64())
	case reflect.String:
		return string(s) < other.String()
	}
	return false
}

// Equal compares the other data and returns true is the same as self's value, not self's value and type
func (s RuneValue) Equal(other Data) bool {
	switch reflect.TypeOf(other).Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		return int32(s) == other.Int32()
	case reflect.Float32, reflect.Float64:
		return int32(s) == int32(other.Float64())
	case reflect.String:
		return string(s) == other.String()
	}
	return false
}

// Greater compares the other data and returns true is it's greater than self
func (s RuneValue) Greater(other Data) bool {
	if s.Equal(other) {
		return false
	}
	return !s.Less(other)
}

// Add adds OperableData together
func (s RuneValue) Add(data OperableData) Data {
	return RuneValue(int32(s) + data.Int32())
}

// Sub subtracts OperableData together
func (s RuneValue) Sub(data OperableData) Data {
	return RuneValue(int32(s) - data.Int32())
}

// Mul multiplies OperableData together
func (s RuneValue) Mul(data OperableData) Data {
	return RuneValue(int32(s) * data.Int32())
}

// Div divides OperableData together
func (s RuneValue) Div(data OperableData) Data {
	return RuneValue(int32(s) / data.Int32())
}

// Mod modulus' OperableData together, **note** mod only works on integer values to floats are cast to ints first
func (s RuneValue) Mod(data OperableData) Data {
	return RuneValue(int(s) % data.Int())
}

// RuneValues type alias for a slice of RuneValue
type RuneValues []rune

func (s RuneValues) Data() []Data {
	d := make([]Data, len(s))
	for j := range s {
		d[j] = RuneValue(s[j])
	}
	return d
}
