package collections

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

type OrderedMap interface {
	Iterable
	Delete(key Data)
	Get(key Data) (Data, bool)
	Set(key Data, value Data)
	String() string
}
