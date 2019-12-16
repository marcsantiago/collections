package counter

import (
	"reflect"

	"github.com/marcsantiago/collections"
)

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
	case collections.Int32SliceType:
		hash = make(IntMap32)
	case collections.Int64SliceType:
		hash = make(IntMap64)
	case collections.Float32SliceType:
		hash = make(FloatMap32)
	case collections.Float64SliceType:
		hash = make(FloatMap64)
	case collections.StringSliceType:
		hash = make(StringMap)
	case collections.UnknownType:
		// if the type cannot be found then we return panic
		// this is not very Go like, since no error is returned
		// however by design we are trying to emulate how python's Counter object would behave
		panic("slice type could not be determined")
	}
	// sets the key and increments duplications
	for i := 0; i < len(data); i++ {
		hash.Update(data[i])
	}

	return hash
}

// determineDataType gets the slice type by examining the first element in the data slice
func determineDataType(data collections.Data) collections.Type {
	switch reflect.TypeOf(data).Kind() {
	case reflect.Int:
		return collections.IntSliceType
	case reflect.Int32:
		return collections.Int32SliceType
	case reflect.Int64:
		return collections.Int64SliceType
	case reflect.Float32:
		return collections.Float32SliceType
	case reflect.Float64:
		return collections.Float64SliceType
	case reflect.String:
		return collections.StringSliceType
	}
	return collections.UnknownType
}
