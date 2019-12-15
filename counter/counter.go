package counter

import (
	"reflect"

	"github.com/marcsantiago/collections"
)

type Container struct {
}

type IntMap map[int]int

func (i IntMap) Get(key collections.Data) (collections.Data, bool) {
	val, ok := i[key.Int()]
	return collections.IntValue(val), ok
}

func (i IntMap) Set(value collections.Data) {
	i[value.Int()]++
}

func (i IntMap) Delete(value collections.Data) {
	delete(i, value.Int())
}

func (i IntMap) Items() []collections.Element {
	items := make([]collections.Element, 0, len(i))
	for key, value := range i {
		items = append(items, collections.Element{
			Key:   collections.IntValue(key),
			Value: collections.IntValue(value),
		})
	}
	return items
}

func (i IntMap) Iterate() <-chan collections.Element {
	ch := make(chan collections.Element)
	go func() {
		for key, value := range i {
			ch <- collections.Element{Key: collections.IntValue(key), Value: collections.IntValue(value)}
		}
		close(ch)
	}()
	return ch
}

// Counter is a python like abstraction on counting slice data
// passing in a single optionalType for type means 0 type reflection is needed
// while collection.Data is an interface, it is expected that all elements are of the same primitive underneath
// e.g []int, []int32, []int64, []float32, []float64, []bool, []string
func Counter(data []collections.Data, optionalType ...collections.Type) collections.CounterMap {
	var hash collections.CounterMap

	specifiedType := collections.UnknownType
	if len(optionalType) > 0 {
		specifiedType = optionalType[0]
	} else {
		// this is the slow case, which uses reflection to try and match the collection type
		specifiedType = determineDataType(data[0])
	}

	switch specifiedType {
	//use reflection to try and get the dataType
	case collections.IntSliceType:
		hash = make(IntMap)
	case collections.UnknownType:
		panic("slice type could not be determined")
	default:
		// if the type cannot be found then we return panic
		// this is not very Go like, since no error is returned
		// however by design we are trying to emulate how python's Counter object would behave
		panic("slice type could not be determined")
	}
	// sets the key and increments duplications
	for i := 0; i < len(data); i++ {
		hash.Set(data[i])
	}

	return hash
}

// determineDataType gets the slice type by examining the first element in the data slice
func determineDataType(data collections.Data) collections.Type {
	switch reflect.TypeOf(data).Kind() {
	case reflect.Int:
		return collections.IntSliceType
	}
	return collections.UnknownType
}
