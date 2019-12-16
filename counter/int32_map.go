package counter

import (
	"encoding/json"
	"sort"

	"github.com/marcsantiago/collections"
)

type IntMap32 map[int32]int

// Get retrieves a data value from the internal map if it exists
func (i IntMap32) Get(key collections.Data) (collections.Data, bool) {
	val, ok := i[key.Int32()]
	return collections.IntValue32(val), ok
}

// Update updates the counter for a value or sets it if it does not exist
func (i IntMap32) Update(key collections.Data) {
	i[key.Int32()]++
}

func (i IntMap32) Subtract(key collections.Data) {
	_, ok := i[key.Int32()]
	if ok {
		i[key.Int32()]--
	}
}

// Delete removes a value from the map by it's key
func (i IntMap32) Delete(value collections.Data) {
	delete(i, value.Int32())
}

// Items returns the internal map as a set of elements
func (i IntMap32) Items() []collections.Element {
	items := make([]collections.Element, 0, len(i))
	for key, value := range i {
		items = append(items, collections.Element{
			Key:   collections.IntValue32(key),
			Value: collections.IntValue(value),
		})
	}
	return items
}

// Iterate creates a channel to create an iterator for he Go range statement
func (i IntMap32) Iterate() <-chan collections.Element {
	ch := make(chan collections.Element)
	go func() {
		for key, value := range i {
			ch <- collections.Element{Key: collections.IntValue32(key), Value: collections.IntValue(value)}
		}
		close(ch)
	}()
	return ch
}

// MostCommon returns the most common values by value
func (i IntMap32) MostCommon(n int) []collections.Element {
	elements := make([]collections.Element, 0, len(i))
	for key, value := range i {
		elements = append(elements, collections.Element{
			Key:   collections.IntValue32(key),
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
func (i IntMap32) String() string {
	b, _ := json.Marshal(i)
	return string(b)
}
