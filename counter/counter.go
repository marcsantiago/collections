package counter

import (
	"github.com/marcsantiago/collections"
)

// Counter is a python like abstraction on counting slice data
// passing in a single optionalType for type means zero type reflection is needed
// while collection.Data is an interface, it is expected that all elements are of the same primitive underneath
// e.g []int, []int32, []int64, []float32, []float64, []bool, []string
func Counter(data []collections.Data, optionalType ...collections.Type) collections.CounterMap {
	var hash collections.CounterMap

	specifiedType := collections.UnknownType
	if len(optionalType) > 0 {
		specifiedType = optionalType[0]
	} else {
		// this is the slow case, which uses reflection to try and match the collection type
		specifiedType = collections.DetermineDataType(data[0])
	}

	switch specifiedType {
	//use reflection to try and get the dataType
	case collections.IntType, collections.IntSliceType:
		hash = make(IntMap)
	case collections.Int32Type, collections.Int32SliceType:
		hash = make(IntMap32)
	case collections.Int64Type, collections.Int64SliceType:
		hash = make(IntMap64)
	case collections.Float32Type, collections.Float32SliceType:
		hash = make(FloatMap32)
	case collections.Float64Type, collections.Float64SliceType:
		hash = make(FloatMap64)
	case collections.StringType, collections.StringSliceType:
		hash = make(StringMap)
	case collections.UnknownType:
		// if the type cannot be found then we panic
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

// GenericCounter unlike counter doesn't care what the data types are, it will simply count anything that satisfies the
// collections.Data interface
func GenericCounter(data []collections.Data) collections.CounterMap {
	hash := make(DataMap)
	// sets the key and increments duplications
	for i := 0; i < len(data); i++ {
		hash.Update(data[i])
	}
	return hash
}
