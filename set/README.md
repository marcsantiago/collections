# set

Package set is a Go parody of the Pythons set implementation and set methods

## Functions

### func [Zero](/set.go#L21)

`func Zero[T any]() (ret T)`

Zero returns the "empty" type of "any" T value

## Types

### type [Set](/set.go#L16)

`type Set[T comparable] struct { ... }`

Set implements a generic set with common set functionality, this implementation is not safe for concurrency, guard your
data with your own mutexes please

#### func (*Set[T]) [Add](/set.go#L37)

`func (s *Set[T]) Add(values ...T)`

Add adds a unique item to the set

#### func (*Set[T]) [Clear](/set.go#L45)

`func (s *Set[T]) Clear()`

Clear removes all elements from the set by removing each element iteratively
this allows for reuse of the backing map

#### func (*Set[T]) [ClearForced](/set.go#L52)

`func (s *Set[T]) ClearForced()`

ClearForced removes all the data by instantiating a new Set

#### func (*Set[T]) [Contains](/set.go#L57)

`func (s *Set[T]) Contains(value T) bool`

Contains returns true if the set contains the inputted data

#### func (*Set[T]) [Copy](/set.go#L63)

`func (s *Set[T]) Copy() *Set[T]`

Copy returns a copy of the set

#### func (*Set[T]) [Difference](/set.go#L71)

`func (s *Set[T]) Difference(sets ...*Set[T]) *Set[T]`

Difference returns a set containing the difference between two or more sets
The returned set contains items that exist only in the first set, and not in both sets.

#### func (*Set[T]) [DifferenceUpdate](/set.go#L85)

`func (s *Set[T]) DifferenceUpdate(sets ...*Set[T])`

DifferenceUpdate updates the current set containing the difference between two or more sets
The returned set contains items that exist only in the first set, and not in both sets.

#### func (*Set[T]) [Discard](/set.go#L96)

`func (s *Set[T]) Discard(value T)`

Discard removes the specified item

#### func (*Set[T]) [DiscardAny](/set.go#L100)

`func (s *Set[T]) DiscardAny(value T)`

#### func (*Set[T]) [Intersection](/set.go#L106)

`func (s *Set[T]) Intersection(sets ...*Set[T]) *Set[T]`

Intersection method returns a set that contains the similarity between two or more sets
The returned set contains only items that exist in both sets, or in all sets if the comparison is done with more than two sets.

#### func (*Set[T]) [IntersectionUpdate](/set.go#L125)

`func (s *Set[T]) IntersectionUpdate(sets ...*Set[T])`

IntersectionUpdate updates the current set that contains the similarity between two or more sets
The returned set contains only items that exist in both sets, or in all sets if the comparison is done with more than two sets.

#### func (*Set[T]) [IsDisjoint](/set.go#L141)

`func (s *Set[T]) IsDisjoint(ss *Set[T]) bool`

IsDisjoint returns whether two sets have a intersection or not

#### func (*Set[T]) [IsSubset](/set.go#L151)

`func (s *Set[T]) IsSubset(ss *Set[T]) bool`

IsSubset returns whether another set contains this set or not

#### func (*Set[T]) [IsSuperset](/set.go#L161)

`func (s *Set[T]) IsSuperset(ss *Set[T]) bool`

IsSuperset returns true if all items in the specified set exists in the original set, otherwise it returns false

#### func (*Set[T]) [Pop](/set.go#L171)

`func (s *Set[T]) Pop() (T, bool)`

Pop removes a random item from the set

#### func (*Set[T]) [Remove](/set.go#L180)

`func (s *Set[T]) Remove(value T)`

Remove removes the specified element

#### func (*Set[T]) [Size](/set.go#L229)

`func (s *Set[T]) Size() int`

Size returns the number of items in the Set

#### func (*Set[T]) [String](/set.go#L234)

`func (s *Set[T]) String() string`

String creates a list representation of the set

#### func (*Set[T]) [SymmetricDifference](/set.go#L185)

`func (s *Set[T]) SymmetricDifference(ss *Set[T]) *Set[T]`

SymmetricDifference returns a set that contains all items from both set, but not the items that are present in both sets

#### func (*Set[T]) [SymmetricDifferenceUpdate](/set.go#L198)

`func (s *Set[T]) SymmetricDifferenceUpdate(ss *Set[T])`

SymmetricDifferenceUpdate updates the current set to contains all items from both set, but not the items that are present in both sets

#### func (*Set[T]) [Union](/set.go#L209)

`func (s *Set[T]) Union(sets ...*Set[T]) *Set[T]`

Union returns a set that contains all items from the original set, and all items from the specified sets.

#### func (*Set[T]) [Update](/set.go#L220)

`func (s *Set[T]) Update(sets ...*Set[T])`

Update updates the current set that contains all items from the original set, and all items from the specified sets.

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
