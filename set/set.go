package set

import (
	"bytes"
	"reflect"
	"strconv"

	"github.com/marcsantiago/collections"
)

type Set map[collections.Data]struct{}

// New creates an initialized set
func New() Set {
	return make(Set)
}

// Add adds a unique item to the set
func (s Set) Add(data collections.Data) {
	s[data] = struct{}{}
}

// Clear removes all elements from the set by removing each element iteratively
// this allows for reuse of the backing map
func (s Set) Clear() {
	for k := range s {
		delete(s, k)
	}
}

// Contains returns true if the set contains the inputted data
func (s Set) Contains(data collections.Data) bool {
	_, ok := s[data]
	return ok
}

// Copy returns a copy of the set
func (s Set) Copy() Set {
	ns := New()
	for k := range s {
		ns.Add(k)
	}
	return ns
}

// Difference returns a set containing the difference between two or more sets
// The returned set contains items that exist only in the first set, and not in both sets.
func (s Set) Difference(sets ...Set) Set {
	ns := s.Copy()
	for _, ss := range sets {
		for data := range ss {
			if ns.Contains(data) {
				delete(ns, data)
			}
		}
	}
	return ns
}

// DifferenceUpdate updates the current set containing the difference between two or more sets
// The returned set contains items that exist only in the first set, and not in both sets.
func (s Set) DifferenceUpdate(sets ...Set) {
	for _, ss := range sets {
		for data := range ss {
			if s.Contains(data) {
				delete(s, data)
			}
		}
	}
}

// Discard removes the specified item
func (s Set) Discard(data collections.Data) {
	delete(s, data)
}

// Intersection method returns a set that contains the similarity between two or more sets
// The returned set contains only items that exist in both sets, or in all sets if the comparison is done with more than two sets.
func (s Set) Intersection(sets ...Set) Set {
	ns := New()
	for item := range s {
		contained := true
		for _, ss := range sets {
			if !ss.Contains(item) {
				contained = false
				break
			}
		}
		if contained {
			ns.Add(item)
		}
	}
	return ns
}

// IntersectionUpdate updates the current set that contains the similarity between two or more sets
// The returned set contains only items that exist in both sets, or in all sets if the comparison is done with more than two sets.
func (s Set) IntersectionUpdate(sets ...Set) {
	for item := range s {
		contained := true
		for _, ss := range sets {
			if !ss.Contains(item) {
				contained = false
				break
			}
		}
		if !contained {
			delete(s, item)
		}
	}
}

// IsDisjoint returns whether two sets have a intersection or not
func (s Set) IsDisjoint(ss Set) bool {
	for item := range s {
		if ss.Contains(item) {
			return false
		}
	}
	return true
}

// IsSubset returns whether another set contains this set or not
func (s Set) IsSubset(ss Set) bool {
	for item := range s {
		if !ss.Contains(item) {
			return false
		}
	}
	return true
}

// IsSuperset returns true if all items in the specified set exists in the original set, otherwise it returns false
func (s Set) IsSuperset(ss Set) bool {
	for item := range ss {
		if !s.Contains(item) {
			return false
		}
	}
	return true
}

// Pop removes a random item from the set
func (s Set) Pop() collections.Data {
	for key := range s {
		delete(s, key)
		return key
	}
	return nil
}

// Remove removes the specified element
func (s Set) Remove(key collections.Data) {
	delete(s, key)
}

// SymmetricDifference returns a set that contains all items from both set, but not the items that are present in both sets
func (s Set) SymmetricDifference(ss Set) Set {
	ns := s.Copy()
	for item := range ss {
		if ns.Contains(item) {
			delete(ns, item)
			continue
		}
		ns.Add(item)
	}
	return ns
}

// SymmetricDifferenceUpdate updates the current set to contains all items from both set, but not the items that are present in both sets
func (s Set) SymmetricDifferenceUpdate(ss Set) Set {
	for item := range ss {
		if s.Contains(item) {
			delete(s, item)
			continue
		}
		s.Add(item)
	}
	return s
}

// Union returns a set that contains all items from the original set, and all items from the specified sets.
func (s Set) Union(sets ...Set) Set {
	ns := s.Copy()
	for _, ss := range sets {
		for item := range ss {
			ns.Add(item)
		}
	}
	return ns
}

// Update updates the current set that contains all items from the original set, and all items from the specified sets.
func (s Set) Update(sets ...Set) {
	for _, ss := range sets {
		for item := range ss {
			s.Add(item)
		}
	}
}

// String creates an set representation of the internal map data
func (s Set) String() string {
	var buf bytes.Buffer
	i, max := 0, len(s)
	for key := range s {
		encode(&buf, key, determineDataType(key))
		if i+1 < max {
			buf.WriteRune(',')
		}
		i++
	}
	return buf.String()
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
