package collections

// CounterMap mimics the Python Counter definitions
type CounterMap interface {
	Iterable
	Delete(value Data)
	Get(key Data) (Data, bool)
	Items() []Element
	MostCommon(n int) []Element
	String() string
	Subtract(value Data)
	Update(value Data)
}
