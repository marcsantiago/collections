package collections

// Container defines a set of rules for types to implement
type Container interface {
	Keys() Iterable
	Values() Iterable
	Type() Type
	String() string
}
