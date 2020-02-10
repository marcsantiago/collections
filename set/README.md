

# set
`import "github.com/marcsantiago/collections/set"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package set is a Go parody of the Pythons set implementation and set methods




## <a name="pkg-index">Index</a>
* [type Set](#Set)
  * [func New() Set](#New)
  * [func (s Set) Add(data collections.Data)](#Set.Add)
  * [func (s Set) Clear()](#Set.Clear)
  * [func (s Set) Contains(data collections.Data) bool](#Set.Contains)
  * [func (s Set) Copy() Set](#Set.Copy)
  * [func (s Set) Difference(sets ...Set) Set](#Set.Difference)
  * [func (s Set) DifferenceUpdate(sets ...Set)](#Set.DifferenceUpdate)
  * [func (s Set) Discard(data collections.Data)](#Set.Discard)
  * [func (s Set) Intersection(sets ...Set) Set](#Set.Intersection)
  * [func (s Set) IntersectionUpdate(sets ...Set)](#Set.IntersectionUpdate)
  * [func (s Set) IsDisjoint(ss Set) bool](#Set.IsDisjoint)
  * [func (s Set) IsSubset(ss Set) bool](#Set.IsSubset)
  * [func (s Set) IsSuperset(ss Set) bool](#Set.IsSuperset)
  * [func (s Set) Pop() collections.Data](#Set.Pop)
  * [func (s Set) Remove(key collections.Data)](#Set.Remove)
  * [func (s Set) String() string](#Set.String)
  * [func (s Set) SymmetricDifference(ss Set) Set](#Set.SymmetricDifference)
  * [func (s Set) SymmetricDifferenceUpdate(ss Set)](#Set.SymmetricDifferenceUpdate)
  * [func (s Set) Union(sets ...Set) Set](#Set.Union)
  * [func (s Set) Update(sets ...Set)](#Set.Update)


#### <a name="pkg-files">Package files</a>
[doc.go](/src/github.com/marcsantiago/collections/set/doc.go) [set.go](/src/github.com/marcsantiago/collections/set/set.go) 






## <a name="Set">type</a> [Set](/src/target/set.go?s=96:134#L11)
``` go
type Set map[collections.Data]struct{}
```






### <a name="New">func</a> [New](/src/target/set.go?s=170:184#L14)
``` go
func New() Set
```
New creates an initialized set





### <a name="Set.Add">func</a> (Set) [Add](/src/target/set.go?s=245:284#L19)
``` go
func (s Set) Add(data collections.Data)
```
Add adds a unique item to the set




### <a name="Set.Clear">func</a> (Set) [Clear](/src/target/set.go?s=436:456#L25)
``` go
func (s Set) Clear()
```
Clear removes all elements from the set by removing each element iteratively
this allows for reuse of the backing map




### <a name="Set.Contains">func</a> (Set) [Contains](/src/target/set.go?s=563:612#L32)
``` go
func (s Set) Contains(data collections.Data) bool
```
Contains returns true if the set contains the inputted data




### <a name="Set.Copy">func</a> (Set) [Copy](/src/target/set.go?s=681:704#L38)
``` go
func (s Set) Copy() Set
```
Copy returns a copy of the set




### <a name="Set.Difference">func</a> (Set) [Difference](/src/target/set.go?s=939:979#L48)
``` go
func (s Set) Difference(sets ...Set) Set
```
Difference returns a set containing the difference between two or more sets
The returned set contains items that exist only in the first set, and not in both sets.




### <a name="Set.DifferenceUpdate">func</a> (Set) [DifferenceUpdate](/src/target/set.go?s=1309:1351#L62)
``` go
func (s Set) DifferenceUpdate(sets ...Set)
```
DifferenceUpdate updates the current set containing the difference between two or more sets
The returned set contains items that exist only in the first set, and not in both sets.




### <a name="Set.Discard">func</a> (Set) [Discard](/src/target/set.go?s=1504:1547#L73)
``` go
func (s Set) Discard(data collections.Data)
```
Discard removes the specified item




### <a name="Set.Intersection">func</a> (Set) [Intersection](/src/target/set.go?s=1792:1834#L79)
``` go
func (s Set) Intersection(sets ...Set) Set
```
Intersection method returns a set that contains the similarity between two or more sets
The returned set contains only items that exist in both sets, or in all sets if the comparison is done with more than two sets.




### <a name="Set.IntersectionUpdate">func</a> (Set) [IntersectionUpdate](/src/target/set.go?s=2274:2318#L98)
``` go
func (s Set) IntersectionUpdate(sets ...Set)
```
IntersectionUpdate updates the current set that contains the similarity between two or more sets
The returned set contains only items that exist in both sets, or in all sets if the comparison is done with more than two sets.




### <a name="Set.IsDisjoint">func</a> (Set) [IsDisjoint](/src/target/set.go?s=2573:2609#L114)
``` go
func (s Set) IsDisjoint(ss Set) bool
```
IsDisjoint returns whether two sets have a intersection or not




### <a name="Set.IsSubset">func</a> (Set) [IsSubset](/src/target/set.go?s=2764:2798#L124)
``` go
func (s Set) IsSubset(ss Set) bool
```
IsSubset returns whether another set contains this set or not




### <a name="Set.IsSuperset">func</a> (Set) [IsSuperset](/src/target/set.go?s=3005:3041#L134)
``` go
func (s Set) IsSuperset(ss Set) bool
```
IsSuperset returns true if all items in the specified set exists in the original set, otherwise it returns false




### <a name="Set.Pop">func</a> (Set) [Pop](/src/target/set.go?s=3174:3209#L144)
``` go
func (s Set) Pop() collections.Data
```
Pop removes a random item from the set




### <a name="Set.Remove">func</a> (Set) [Remove](/src/target/set.go?s=3322:3363#L153)
``` go
func (s Set) Remove(key collections.Data)
```
Remove removes the specified element




### <a name="Set.String">func</a> (Set) [String](/src/target/set.go?s=4528:4556#L202)
``` go
func (s Set) String() string
```
String creates an set representation of the internal map data




### <a name="Set.SymmetricDifference">func</a> (Set) [SymmetricDifference](/src/target/set.go?s=3509:3553#L158)
``` go
func (s Set) SymmetricDifference(ss Set) Set
```
SymmetricDifference returns a set that contains all items from both set, but not the items that are present in both sets




### <a name="Set.SymmetricDifferenceUpdate">func</a> (Set) [SymmetricDifferenceUpdate](/src/target/set.go?s=3827:3873#L171)
``` go
func (s Set) SymmetricDifferenceUpdate(ss Set)
```
SymmetricDifferenceUpdate updates the current set to contains all items from both set, but not the items that are present in both sets




### <a name="Set.Union">func</a> (Set) [Union](/src/target/set.go?s=4088:4123#L182)
``` go
func (s Set) Union(sets ...Set) Set
```
Union returns a set that contains all items from the original set, and all items from the specified sets.




### <a name="Set.Update">func</a> (Set) [Update](/src/target/set.go?s=4351:4383#L193)
``` go
func (s Set) Update(sets ...Set)
```
Update updates the current set that contains all items from the original set, and all items from the specified sets.








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
