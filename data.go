package collections

// Data interface for casting and getting data values used to avoid reflection
type Data interface {
	Int() int
	String() string
}
