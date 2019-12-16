package collections

type CounterMap interface {
	Iterable
	Get(key Data) (Data, bool)
	Set(value Data)
	Delete(value Data)
	Items() []Element
	String() string
}
