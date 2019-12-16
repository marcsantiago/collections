package collections

// Data interface for casting and getting data values used to avoid reflection
type Data interface {
	Int() int
	Int32() int32
	Int64() int64
	Float32() float32
	Float64() float64
	String() string
}
