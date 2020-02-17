package collections

// CounterMap mimics the Python Counter definitions
// this also implements the collections.Map interface to be able to convert into a ChainMap
type CounterMap interface {
	Iterable
	Delete(value Data)
	Get(key Data) (Data, bool)
	Items() []Element
	MostCommon(n int) []Element
	Set(key Data, value Data)
	String() string
	Subtract(value Data)
	Update(value Data)
}

// Map is a genetic map
type Map interface {
	Iterable
	Delete(value Data)
	Get(key Data) (Data, bool)
	Set(key Data, value Data)
	String() string
	Items() []Element
}
