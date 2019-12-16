package collections

type Type int

const (
	UnknownType Type = iota
	IntSliceType
	Int32SliceType
	Int64SliceType
	Float32SliceType
	Float64SliceType
	StringSliceType
)
