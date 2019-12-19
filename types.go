package collections

type Type int

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
