package counter

import (
	"encoding/json"
	"sort"
	"strconv"

	"github.com/marcsantiago/collections"
)

type FloatMap64 map[float64]int

// NewFloatMap64 takes in hash data, counts, and returns a concrete type that implements a CounterMap
func NewFloatMap64(hash map[float64]int) FloatMap64 {
	var nh FloatMap64
	if hash != nil {
		nh = make(FloatMap64)
		for key := range hash {
			nh.Update(collections.FloatValue64(key))
		}
	}
	// return an empty Map ready for use
	return FloatMap64(hash)
}

// Get retrieves a data value from the internal map if it exists
func (i FloatMap64) Get(key collections.Data) (collections.Data, bool) {
	val, ok := i[key.Float64()]
	return collections.FloatValue64(val), ok
}

// Update updates the counter for a value or sets the value it if it does not exist
func (i FloatMap64) Update(key collections.Data) {
	i[key.Float64()]++
}

// Subtract removes 1 from the counter if the key exists
func (i FloatMap64) Subtract(key collections.Data) {
	_, ok := i[key.Float64()]
	if ok {
		i[key.Float64()]--
	}
}

// Delete removes the element from the internal map
func (i FloatMap64) Delete(key collections.Data) {
	delete(i, key.Float64())
}

// Items returns the internal map as a set of elements
func (i FloatMap64) Items() []collections.Element {
	items := make([]collections.Element, 0, len(i))
	for key, value := range i {
		items = append(items, collections.Element{
			Key:   collections.FloatValue64(key),
			Value: collections.IntValue(value),
		})
	}
	return items
}

// Iterate creates a channel to create an iterator for he Go range statement
func (i FloatMap64) Iterate() <-chan collections.Element {
	ch := make(chan collections.Element)
	go func() {
		for key, value := range i {
			ch <- collections.Element{Key: collections.FloatValue64(key), Value: collections.IntValue(value)}
		}
		close(ch)
	}()
	return ch
}

// MostCommon returns the most common values by value
func (i FloatMap64) MostCommon(n int) []collections.Element {
	elements := make([]collections.Element, 0, len(i))
	for key, value := range i {
		elements = append(elements, collections.Element{
			Key:   collections.FloatValue64(key),
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
// maps with float keys are not valid JSON, so the keys are converted to strings
func (i FloatMap64) String() string {
	m := make(map[string]int)
	for key, value := range i {
		m[strconv.FormatFloat(key, 'f', -1, 64)] = value
	}
	b, _ := json.Marshal(m)
	return string(b)
}
