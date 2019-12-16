package collections

type Type int

const (
	UnknownType Type = iota
	IntSliceType
	Int32SliceType
	Int64SliceType
	// not implemented as floats cannot be map keys
	//Float32SliceType
	//Float64SliceType
	StringSliceType
)
