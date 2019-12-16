package counter

import (
	"encoding/json"
	"sort"
	"strconv"

	"github.com/marcsantiago/collections"
)

type FloatMap32 map[float32]int

// Get retrieves a data value from the internal map if it exists
func (i FloatMap32) Get(key collections.Data) (collections.Data, bool) {
	val, ok := i[key.Float32()]
	return collections.FloatValue32(val), ok
}

// Update updates the counter for a value or sets it if it does not exist
func (i FloatMap32) Update(key collections.Data) {
	i[key.Float32()]++
}

func (i FloatMap32) Subtract(key collections.Data) {
	_, ok := i[key.Float32()]
	if ok {
		i[key.Float32()]--
	}
}

// Delete removes a value from the map by it's key
func (i FloatMap32) Delete(value collections.Data) {
	delete(i, value.Float32())
}

// Items returns the internal map as a set of elements
func (i FloatMap32) Items() []collections.Element {
	items := make([]collections.Element, 0, len(i))
	for key, value := range i {
		items = append(items, collections.Element{
			Key:   collections.FloatValue32(key),
			Value: collections.IntValue(value),
		})
	}
	return items
}

// Iterate creates a channel to create an iterator for he Go range statement
func (i FloatMap32) Iterate() <-chan collections.Element {
	ch := make(chan collections.Element)
	go func() {
		for key, value := range i {
			ch <- collections.Element{Key: collections.FloatValue32(key), Value: collections.IntValue(value)}
		}
		close(ch)
	}()
	return ch
}

// MostCommon returns the most common values by value
func (i FloatMap32) MostCommon(n int) []collections.Element {
	elements := make([]collections.Element, 0, len(i))
	for key, value := range i {
		elements = append(elements, collections.Element{
			Key:   collections.FloatValue32(key),
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
// maps with float keys do not marshal well so the keys are converted to strings
func (i FloatMap32) String() string {
	m := make(map[string]int)
	for key, value := range i {
		m[strconv.FormatFloat(float64(key), 'f', -1, 32)] = value
	}
	b, _ := json.Marshal(m)
	return string(b)
}
