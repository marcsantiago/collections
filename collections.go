package collections

type Type int

const (
	UnknownType Type = iota
	IntSlice
	Int32Slice
	Int64Slice
	Float32Slice
	Float64Slice
	BoolSlice
	StringSlice

	StringIntMap
	StringInt32Map
	StringInt64Map
	StringFloat32Map
	StringFloat64Map
	StringBoolMap
	StringStringMap
)

type Element struct {
	Key   Data
	Value Data
}

// Iterable is anything that can be ranged on
type Iterable interface {
	Iterate() (Element, bool)
}
