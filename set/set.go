package set

import "fmt"

//type (
//	Ordered interface {
//		~int | ~int8 | ~int16 | ~int32 | ~int64 |
//			~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
//			~float32 | ~float64 |
//			~string
//	}
//)

// Set implements a generic set with common set functionality, this implementation is not safe for concurrency, guard your
// data with your own mutexes please
type Set[T comparable] struct {
	values map[T]struct{}
}

// Zero returns the "empty" type of "any" T value
func Zero[T any]() (ret T) {
	return
}

// New creates an initialized set
func New[T comparable](values ...T) *Set[T] {
	m := make(map[T]struct{}, len(values))
	for _, v := range values {
		m[v] = struct{}{}
	}
	return &Set[T]{
		values: m,
	}
}

// Add adds a unique item to the set
func (s *Set[T]) Add(values ...T) {
	for _, value := range values {
		s.values[value] = struct{}{}
	}
}

// Clear removes all elements from the set by removing each element iteratively
// this allows for reuse of the backing map
func (s *Set[T]) Clear() {
	for k := range s.values {
		delete(s.values, k)
	}
}

// ClearForced removes all the data by instantiating a new Set
func (s *Set[T]) ClearForced() {
	s.values = make(map[T]struct{})
}

// Contains returns true if the set contains the inputted data
func (s *Set[T]) Contains(value T) bool {
	_, ok := s.values[value]
	return ok
}

// Copy returns a copy of the set
func (s *Set[T]) Copy() *Set[T] {
	return &Set[T]{
		values: s.values,
	}
}

// Difference returns a set containing the difference between two or more sets
// The returned set contains items that exist only in the first set, and not in both sets.
func (s *Set[T]) Difference(sets ...*Set[T]) *Set[T] {
	ns := s.Copy()
	for _, ss := range sets {
		for value := range ss.values {
			if ns.Contains(value) {
				delete(ns.values, value)
			}
		}
	}
	return ns
}

// DifferenceUpdate updates the current set containing the difference between two or more sets
// The returned set contains items that exist only in the first set, and not in both sets.
func (s *Set[T]) DifferenceUpdate(sets ...*Set[T]) {
	for _, ss := range sets {
		for value := range ss.values {
			if s.Contains(value) {
				delete(s.values, value)
			}
		}
	}
}

// Discard removes the specified item
func (s *Set[T]) Discard(value T) {
	delete(s.values, value)
}

func (s *Set[T]) DiscardAny(value T) {
	delete(s.values, value)
}

// Intersection method returns a set that contains the similarity between two or more sets
// The returned set contains only items that exist in both sets, or in all sets if the comparison is done with more than two sets.
func (s *Set[T]) Intersection(sets ...*Set[T]) *Set[T] {
	ns := New[T]()
	for value := range s.values {
		contained := true
		for _, ss := range sets {
			if !ss.Contains(value) {
				contained = false
				break
			}
		}
		if contained {
			ns.Add(value)
		}
	}
	return ns
}

// IntersectionUpdate updates the current set that contains the similarity between two or more sets
// The returned set contains only items that exist in both sets, or in all sets if the comparison is done with more than two sets.
func (s *Set[T]) IntersectionUpdate(sets ...*Set[T]) {
	for item := range s.values {
		contained := true
		for _, ss := range sets {
			if !ss.Contains(item) {
				contained = false
				break
			}
		}
		if !contained {
			delete(s.values, item)
		}
	}
}

// IsDisjoint returns whether two sets have a intersection or not
func (s *Set[T]) IsDisjoint(ss *Set[T]) bool {
	for item := range s.values {
		if ss.Contains(item) {
			return false
		}
	}
	return true
}

// IsSubset returns whether another set contains this set or not
func (s *Set[T]) IsSubset(ss *Set[T]) bool {
	for item := range s.values {
		if !ss.Contains(item) {
			return false
		}
	}
	return true
}

// IsSuperset returns true if all items in the specified set exists in the original set, otherwise it returns false
func (s *Set[T]) IsSuperset(ss *Set[T]) bool {
	for item := range ss.values {
		if !s.Contains(item) {
			return false
		}
	}
	return true
}

// Pop removes a random item from the set
func (s *Set[T]) Pop() (T, bool) {
	for key := range s.values {
		delete(s.values, key)
		return key, true
	}
	return Zero[T](), false
}

// Remove removes the specified element
func (s *Set[T]) Remove(value T) {
	delete(s.values, value)
}

// SymmetricDifference returns a set that contains all items from both set, but not the items that are present in both sets
func (s *Set[T]) SymmetricDifference(ss *Set[T]) *Set[T] {
	ns := s.Copy()
	for item := range ss.values {
		if ns.Contains(item) {
			delete(ns.values, item)
			continue
		}
		ns.Add(item)
	}
	return ns
}

// SymmetricDifferenceUpdate updates the current set to contains all items from both set, but not the items that are present in both sets
func (s *Set[T]) SymmetricDifferenceUpdate(ss *Set[T]) {
	for item := range ss.values {
		if s.Contains(item) {
			delete(s.values, item)
			continue
		}
		s.Add(item)
	}
}

// Union returns a set that contains all items from the original set, and all items from the specified sets.
func (s *Set[T]) Union(sets ...*Set[T]) *Set[T] {
	ns := s.Copy()
	for _, ss := range sets {
		for item := range ss.values {
			ns.Add(item)
		}
	}
	return ns
}

// Update updates the current set that contains all items from the original set, and all items from the specified sets.
func (s *Set[T]) Update(sets ...*Set[T]) {
	for _, ss := range sets {
		for item := range ss.values {
			s.Add(item)
		}
	}
}

// Size returns the number of items in the Set
func (s *Set[T]) Size() int {
	return len(s.values)
}

//String creates a list representation of the set
func (s *Set[T]) String() string {
	values := make([]T, 0, s.Size())
	for value := range s.values {
		values = append(values, value)
	}
	return fmt.Sprint(values)
}
