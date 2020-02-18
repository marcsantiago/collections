package chain

import (
	"bytes"
	"reflect"
	"strconv"

	"github.com/marcsantiago/collections"
)

type Maps []collections.Map

// New returns a new chain map
// if no maps are passed i new chain map with an initialized map is return
// else all the maps are merged into indexes
// Lookups search the underlying mappings successively until a key is found. In contrast, writes, updates, and deletions only operate on the first mapping.
func New(maps ...collections.Map) Maps {
	var cMap []collections.Map
	if len(maps) == 0 {
		return append(cMap, collections.NewGenericMap())
	}

	for _, m := range maps {
		cMap = append(cMap, m)
	}
	return cMap
}

// Delete removes the element from the first internal map
func (m Maps) Delete(key collections.Data) {
	m[0].Delete(key)
}

// Get retrieves a data value from the first internal map that contains they targeted key
func (m Maps) Get(key collections.Data) (collections.Data, bool) {
	for _, mm := range m {
		value, ok := mm.Get(key)
		if ok {
			return value, ok
		}
	}
	return nil, false
}

// NewChild returns a new ChainMap containing a new map followed by all of the maps in the current instance. If m is specified, it becomes the new map at the front of the list of mappings; if not specified, an empty dict is used
func (m Maps) NewChild(mm collections.Map) Maps {
	if mm == nil {
		return append([]collections.Map{collections.NewGenericMap()}, m...)
	}
	return append([]collections.Map{mm}, m...)
}

// Parents returning a new ChainMap containing all of the maps in the current instance except the first one. This is useful for skipping the first map in the search.
func (m Maps) Parents() Maps {
	return m[1:]
}

// Set adds a key value pairing to the first internal map
func (m Maps) Set(key collections.Data, value collections.Data) {
	m[0].Set(key, value)
}

// String returns the JSON string representation of the map data
func (m Maps) String() string {
	var buf bytes.Buffer
	buf.WriteString("[")

	chainSize := len(m)
	for k, mm := range m {
		var i int
		max := mm.Len()
		buf.WriteString("{")
		for elem := range mm.Iterate() {
			encode(&buf, elem.Key, determineDataType(elem.Key))
			buf.WriteRune(':')
			encode(&buf, elem.Value, determineDataType(elem.Value))
			if i+1 < max {
				buf.WriteRune(',')
			}
			i++
		}
		buf.WriteString("}")
		if k+i < chainSize {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("]")
	return buf.String()
}

// Items returns the internal map as a set of elements
func (m Maps) Items() []collections.Element {
	var total int
	for _, mm := range m {
		total += mm.Len()
	}

	items := make([]collections.Element, 0, total)
	for _, mm := range m {
		for elem := range mm.Iterate() {
			items = append(items, collections.Element{
				Key:   elem.Key,
				Value: elem.Value,
			})
		}
	}
	return items
}

// Iterate creates a channel to create an iterator for he Go range statement
func (m Maps) Iterate() <-chan collections.Element {
	ch := make(chan collections.Element)
	go func() {
		for _, mm := range m {
			for elem := range mm.Iterate() {
				ch <- collections.Element{Key: elem.Key, Value: elem.Value}
			}
		}
		close(ch)
	}()
	return ch
}

// determineDataType gets the internal data type converts it to a supported collections type
func determineDataType(data collections.Data) collections.Type {
	switch reflect.TypeOf(data).Kind() {
	case reflect.Int:
		return collections.IntType
	case reflect.Int32:
		return collections.Int32Type
	case reflect.Int64:
		return collections.Int64Type
	case reflect.Float32:
		return collections.Float32Type
	case reflect.Float64:
		return collections.Float64Type
	case reflect.String:
		return collections.StringType
	}
	return collections.UnknownType
}

// encode writes data into a bytes buffer, used in the String method
func encode(encoder *bytes.Buffer, data collections.Data, t collections.Type) {
	// rip for optimization
	const size = 0
	b := make([]byte, size)
	switch t {
	case collections.IntType, collections.Int32Type, collections.Int64Type:
		encoder.Write(strconv.AppendInt(b, data.Int64(), 10))
	case collections.Float32Type:
		encoder.Write(strconv.AppendFloat(b, data.Float64(), 'f', -1, 32))
	case collections.Float64Type:
		encoder.Write(strconv.AppendFloat(b, data.Float64(), 'f', -1, 64))
	case collections.StringType:
		encoder.Write([]byte("\"" + data.String() + "\""))
	}
}
