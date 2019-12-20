package ordereddict

import (
	"bytes"
	"reflect"
	"strconv"

	"github.com/marcsantiago/collections"
)

type OrderedDict struct {
	keys      []collections.Data
	hash      map[collections.Data]collections.Data
	keyType   collections.Type
	valueType collections.Type
}

func New() OrderedDict {
	return OrderedDict{
		hash:      make(map[collections.Data]collections.Data),
		keyType:   collections.UnknownType,
		valueType: collections.UnknownType,
	}
}

func (o OrderedDict) Get(key collections.Data) (collections.Data, bool) {
	val, ok := o.hash[key]
	return val, ok
}

func (o OrderedDict) Delete(key collections.Data) {
	if _, ok := o.hash[key]; !ok || len(o.hash) == 0 {
		return
	}
	delete(o.hash, key)

	var i int
	for j, value := range o.keys {
		if value == key {
			i = j
			break
		}
	}
	o.keys = o.keys[:i+copy(o.keys[i:], o.keys[i+1:])]
}

func (o *OrderedDict) Set(key collections.Data, value collections.Data) {
	if _, ok := o.hash[key]; ok {
		return
	}

	// this allows for proper sudo JSON printing in the String method
	if o.keyType == collections.UnknownType || o.valueType == collections.UnknownType {
		o.keyType, o.valueType = determineDataType(key), determineDataType(value)
	}

	o.hash[key] = value
	o.keys = append(o.keys, key)
}

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

func (o OrderedDict) Reverse() {
	for i := len(o.keys)/2 - 1; i >= 0; i-- {
		opp := len(o.keys) - 1 - i
		o.keys[i], o.keys[opp] = o.keys[opp], o.keys[i]
	}
}

func (o OrderedDict) String() string {
	var buf bytes.Buffer
	buf.WriteString("OrderedDict([")
	max := len(o.keys)
	for i, key := range o.keys {
		encode(&buf, key, o.keyType)
		buf.WriteRune(':')
		encode(&buf, o.hash[key], o.valueType)
		if i+1 < max {
			buf.WriteRune(',')
		}
	}
	buf.WriteString("])")
	return buf.String()
}

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

func encode(encoder *bytes.Buffer, data collections.Data, t collections.Type) {
	// rip for optimization
	const size = 10
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