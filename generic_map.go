package collections

import (
	"encoding/json"
)

// GenericMap is the default implementation of a map using the Data interface
type GenericMap map[Data]Data

func NewGenericMap() GenericMap {
	return make(GenericMap)
}

// Get retrieves a data value from the internal map if it exists
func (i GenericMap) Get(key Data) (Data, bool) {
	val, ok := i[key]
	return val, ok
}

// Len returns the number of stored keys
func (i GenericMap) Len() int {
	return len(i)
}

// Set adds a key value pairing to the internal map
func (i GenericMap) Set(key Data, value Data) {
	i[key] = value
}

// Delete removes the element from the internal map
func (i GenericMap) Delete(key Data) {
	delete(i, key)
}

// Items returns the internal map as a set of elements
func (i GenericMap) Items() []Element {
	items := make([]Element, 0, len(i))
	for key, value := range i {
		items = append(items, Element{
			Key:   key,
			Value: value,
		})
	}
	return items
}

// Iterate creates a channel to create an iterator for he Go range statement
func (i GenericMap) Iterate() <-chan Element {
	ch := make(chan Element)
	go func() {
		for key, value := range i {
			ch <- Element{Key: key, Value: value}
		}
		close(ch)
	}()
	return ch
}

// String returns the JSON string representation of the map data
func (i GenericMap) String() string {
	b, _ := json.Marshal(i)
	return string(b)
}
