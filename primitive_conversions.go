package collections

import "strconv"

// IntValue type alias for int
type IntValue int

// Int casts and returns int value
func (i IntValue) Int() int {
	return int(i)
}

// String casts and returns string value
func (i IntValue) String() string {
	return strconv.Itoa(int(i))
}

// IntValues type alias for a slice of IntValue
type IntValues []IntValue

func (i IntValues) Data() []Data {
	d := make([]Data, len(i))
	for j := range i {
		d[j] = i[j]
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

// String casts and returns string value
func (s StringValue) String() string {
	return string(s)
}

// StringValues type alias for a slice of StringValue
type StringValues []StringValue

//func (s StringValue) Data() []Data {
//	d := make([]Data, len(s))
//	for j := range s {
//		d[j] = s[j]
//	}
//	return d
//}
