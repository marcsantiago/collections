package counter

import (
	"encoding/json"
	"sort"

	"github.com/marcsantiago/collections"
)

type IntMap map[int]int

// NewIntMap takes in hash data, counts, and returns a concrete type that implements a CounterMap
func NewIntMap(hash map[int]int) IntMap {
	var nh IntMap
	if hash != nil {
		nh = make(IntMap)
		for key := range hash {
			nh.Update(collections.IntValue(key))
		}
	}
	// return an empty Map ready for use
	return IntMap(hash)
}

// Get retrieves a data value from the internal map if it exists
func (i IntMap) Get(key collections.Data) (collections.Data, bool) {
	val, ok := i[key.Int()]
	return collections.IntValue(val), ok
}

// Update updates the counter for a value or sets the value it if it does not exist
func (i IntMap) Update(key collections.Data) {
	i[key.Int()]++
}

// Subtract removes 1 from the counter if the key exists
func (i IntMap) Subtract(key collections.Data) {
	_, ok := i[key.Int()]
	if ok {
		i[key.Int()]--
	}
}

// Delete removes the element from the internal map
func (i IntMap) Delete(key collections.Data) {
	delete(i, key.Int())
}

// Items returns the internal map as a set of elements
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

// Iterate creates a channel to create an iterator for he Go range statement
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

// MostCommon returns the most common values by value
func (i IntMap) MostCommon(n int) []collections.Element {
	elements := make([]collections.Element, 0, len(i))
	for key, value := range i {
		elements = append(elements, collections.Element{
			Key:   collections.IntValue(key),
			Value: collections.IntValue(value),
		})
	}

	if n <= 0 || n >= len(i) {
		n = len(i) - 1
	}
	sort.Sort(collections.ElementsByValueIntDesc(elements))
	return elements[:n]
}

// String returns the JSON string representation of the map data
func (i IntMap) String() string {
	b, _ := json.Marshal(i)
	return string(b)
}
