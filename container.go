package collections

import "strconv"

// Container defines a set of rules for types to implement
type Container interface {
	Keys() Iterable
	Values() Iterable
	Type() Type
	String() string
}

type Data interface {
	Int() int
	String() string
}

type IntValue int

func (i IntValue) Int() int {
	return int(i)
}

func (i IntValue) String() string {
	return strconv.Itoa(int(i))
}

type StringValue string

func (s StringValue) Int() int {
	n, _ := strconv.Atoi(string(s))
	return n
}

func (s StringValue) String() string {
	return string(s)
}
