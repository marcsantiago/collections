package collections

type Type int

const (
	UnknownType Type = iota
	IntSliceType
	//Need to implement more types in the counter package
	//Int32SliceType
	//Int64SliceType
	//Float32SliceType
	//Float64SliceType
	//BoolSliceType
	//StringSliceType
)

type Element struct {
	Key   Data
	Value Data
}

// Iterable is anything that can be ranged on
type Iterable interface {
	Iterate() <-chan Element
}
