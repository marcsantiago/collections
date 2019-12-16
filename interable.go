package collections

// Iterable is anything that can be ranged on
type Iterable interface {
	Iterate() <-chan Element
}
