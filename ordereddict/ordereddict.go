package ordereddict

import (
	"bytes"

	"github.com/marcsantiago/collections"
)

type OrderedDict struct {
	keys []collections.Data
	hash collections.GenericMap
}

// OrderedDict returns a map initialized structure
func New() OrderedDict {
	return OrderedDict{
		hash: make(collections.GenericMap),
	}
}

// Get retrieves a data value from the internal map if it exists
func (o OrderedDict) Get(key collections.Data) (collections.Data, bool) {
	val, ok := o.hash[key]
	return val, ok
}

// Delete removes the element from the internal map and the backing slice that records order
func (o *OrderedDict) Delete(key collections.Data) {
	if _, ok := o.hash[key]; !ok {
		return
	}
	delete(o.hash, key)

	var i int
	for j, value := range o.keys {
		if value == key {
			o.keys[j] = nil // avoid memory leaks
			i = j
			break
		}
	}
	o.keys = o.keys[:i+copy(o.keys[i:], o.keys[i+1:])]
}

// Len returns the number of stored keys
func (o OrderedDict) Len() int {
	return len(o.keys)
}

// Set updates the internal map and backing slice that records order
func (o *OrderedDict) Set(key collections.Data, value collections.Data) {
	if _, ok := o.hash[key]; !ok {
		// only append to the backing slice if the key doesn't already exist
		o.keys = append(o.keys, key)
	}
	// add or update the key
	o.hash[key] = value
}

// Items returns the internal map as a set of elements
func (o OrderedDict) Items() []collections.Element {
	items := make([]collections.Element, 0, len(o.keys))
	for _, key := range o.keys {
		items = append(items, collections.Element{
			Key:   key,
			Value: o.hash[key],
		})
	}
	return items
}

// Iterate creates a channel to create an iterator for he Go range statement
func (o OrderedDict) Iterate() <-chan collections.Element {
	ch := make(chan collections.Element)
	go func() {
		for _, key := range o.keys {
			ch <- collections.Element{Key: key, Value: o.hash[key]}
		}
		close(ch)
	}()
	return ch
}

// Reverse does an inplace reverse sort of the backing slice
func (o OrderedDict) Reverse() {
	for i := len(o.keys)/2 - 1; i >= 0; i-- {
		opp := len(o.keys) - 1 - i
		o.keys[i], o.keys[opp] = o.keys[opp], o.keys[i]
	}
}

// String creates an OrderedDict representation of the internal map data
func (o OrderedDict) String() string {
	var buf bytes.Buffer
	buf.WriteString("OrderedDict([")
	max := len(o.keys)
	for i, key := range o.keys {
		collections.StringEncoder(&buf, key, collections.DetermineDataType(key))
		buf.WriteRune(':')
		collections.StringEncoder(&buf, o.hash[key], collections.DetermineDataType(o.hash[key]))
		if i+1 < max {
			buf.WriteRune(',')
		}
	}
	buf.WriteString("])")
	return buf.String()
}
