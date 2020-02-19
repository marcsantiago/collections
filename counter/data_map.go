package counter

import (
	"encoding/json"
	"sort"

	"github.com/marcsantiago/collections"
)

type DataMap map[collections.Data]int

// NewDataMap takes in hash data, counts, and returns a concrete type that implements a CounterMap
func NewDataMap(hash map[collections.Data]int) DataMap {
	var nh DataMap
	if hash != nil {
		nh = make(DataMap)
		for key := range hash {
			nh.Update(key)
		}
	}
	return hash
}

// Get retrieves a data value from the internal map if it exists
func (i DataMap) Get(key collections.Data) (collections.Data, bool) {
	val, ok := i[key]
	return collections.IntValue(val), ok
}

// Len returns the number of stored keys
func (i DataMap) Len() int {
	return len(i)
}

// Update updates the counter for a value or sets the value it if it does not exist
func (i DataMap) Update(key collections.Data) {
	i[key]++
}

// Set replaces a keys counter data with another integer or creates a new key with data
func (i DataMap) Set(key collections.Data, value collections.Data) {
	i[key] = value.Int()
}

// Subtract removes 1 from the counter if the key exists
func (i DataMap) Subtract(key collections.Data) {
	_, ok := i[key]
	if ok {
		i[key]--
	}
}

// Delete removes the element from the internal map
func (i DataMap) Delete(key collections.Data) {
	delete(i, key)
}

// Items returns the internal map as a set of elements
func (i DataMap) Items() []collections.Element {
	items := make([]collections.Element, 0, len(i))
	for key, value := range i {
		items = append(items, collections.Element{
			Key:   key,
			Value: collections.IntValue(value),
		})
	}
	return items
}

// Iterate creates a channel to create an iterator for he Go range statement
func (i DataMap) Iterate() <-chan collections.Element {
	ch := make(chan collections.Element)
	go func() {
		for key, value := range i {
			ch <- collections.Element{Key: key, Value: collections.IntValue(value)}
		}
		close(ch)
	}()
	return ch
}

// MostCommon returns the most common values by value
func (i DataMap) MostCommon(n int) []collections.Element {
	elements := make([]collections.Element, 0, len(i))
	for key, value := range i {
		elements = append(elements, collections.Element{
			Key:   key,
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
func (i DataMap) String() string {
	b, _ := json.Marshal(i)
	return string(b)
}
