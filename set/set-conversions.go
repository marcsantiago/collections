package set

import "github.com/marcsantiago/collections"

// ConvertIntSlice converts a int slice to a set
func ConvertIntSlice(s []int) Set {
	ns := New()
	for _, val := range s {
		ns.Add(collections.IntValue(val))
	}
	return ns
}

// ConvertInt32Slice converts a int32 slice to a set
func ConvertInt32Slice(s []int32) Set {
	ns := New()
	for _, val := range s {
		ns.Add(collections.IntValue32(val))
	}
	return ns
}

// ConvertInt64Slice converts a int64 slice to a set
func ConvertInt64Slice(s []int64) Set {
	ns := New()
	for _, val := range s {
		ns.Add(collections.IntValue64(val))
	}
	return ns
}

// ConvertFloat32Slice converts a float32 slice to a set
func ConvertFloat32Slice(s []float32) Set {
	ns := New()
	for _, val := range s {
		ns.Add(collections.FloatValue32(val))
	}
	return ns
}

// ConvertFloat64Slice converts a float64 slice to a set
func ConvertFloat64Slice(s []float64) Set {
	ns := New()
	for _, val := range s {
		ns.Add(collections.FloatValue64(val))
	}
	return ns
}

// ConvertStringSlice converts a string slice to a set
func ConvertStringSlice(s []string) Set {
	ns := New()
	for _, val := range s {
		ns.Add(collections.StringValue(val))
	}
	return ns
}

// ConvertRuneSlice converts a rune slice to a set
func ConvertRuneSlice(s []rune) Set {
	ns := New()
	for _, val := range s {
		ns.Add(collections.RuneValue(val))
	}
	return ns
}
