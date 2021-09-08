package collections

type Data interface {
	Bool() bool
	String() string
	ComparableData
	NumericData
}

type NumericData interface {
	Int() int
	Int32() int32
	Int64() int64
	Float32() float32
	Float64() float64
}

type ComparableData interface {
	Less(other Data) bool
	Equal(other Data) bool
	Greater(other Data) bool
}

type OperableData interface {
	NumericData
	Add(data OperableData) Data
	Sub(data OperableData) Data
	Mul(data OperableData) Data
	Div(data OperableData) Data
	Mod(data OperableData) Data
}
