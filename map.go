package collections

type CounterMap interface {
	Iterable
	Get(key Data) (Data, bool)
	Update(value Data)
	Subtract(value Data)
	Delete(value Data)
	Items() []Element
	String() string
	MostCommon(n int) []Element
}
