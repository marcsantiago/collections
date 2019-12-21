

# counter
`import "github.com/marcsantiago/collections/counter"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package counter is a Go parody of the Python Counter object




## <a name="pkg-index">Index</a>
* [func Counter(data []collections.Data, optionalType ...collections.Type) collections.CounterMap](#Counter)
* [type FloatMap32](#FloatMap32)
  * [func NewFloatMap32(hash map[float32]int) FloatMap32](#NewFloatMap32)
  * [func (i FloatMap32) Delete(key collections.Data)](#FloatMap32.Delete)
  * [func (i FloatMap32) Get(key collections.Data) (collections.Data, bool)](#FloatMap32.Get)
  * [func (i FloatMap32) Items() []collections.Element](#FloatMap32.Items)
  * [func (i FloatMap32) Iterate() &lt;-chan collections.Element](#FloatMap32.Iterate)
  * [func (i FloatMap32) MostCommon(n int) []collections.Element](#FloatMap32.MostCommon)
  * [func (i FloatMap32) String() string](#FloatMap32.String)
  * [func (i FloatMap32) Subtract(key collections.Data)](#FloatMap32.Subtract)
  * [func (i FloatMap32) Update(key collections.Data)](#FloatMap32.Update)
* [type FloatMap64](#FloatMap64)
  * [func NewFloatMap64(hash map[float64]int) FloatMap64](#NewFloatMap64)
  * [func (i FloatMap64) Delete(key collections.Data)](#FloatMap64.Delete)
  * [func (i FloatMap64) Get(key collections.Data) (collections.Data, bool)](#FloatMap64.Get)
  * [func (i FloatMap64) Items() []collections.Element](#FloatMap64.Items)
  * [func (i FloatMap64) Iterate() &lt;-chan collections.Element](#FloatMap64.Iterate)
  * [func (i FloatMap64) MostCommon(n int) []collections.Element](#FloatMap64.MostCommon)
  * [func (i FloatMap64) String() string](#FloatMap64.String)
  * [func (i FloatMap64) Subtract(key collections.Data)](#FloatMap64.Subtract)
  * [func (i FloatMap64) Update(key collections.Data)](#FloatMap64.Update)
* [type IntMap](#IntMap)
  * [func NewIntMap(hash map[int]int) IntMap](#NewIntMap)
  * [func (i IntMap) Delete(key collections.Data)](#IntMap.Delete)
  * [func (i IntMap) Get(key collections.Data) (collections.Data, bool)](#IntMap.Get)
  * [func (i IntMap) Items() []collections.Element](#IntMap.Items)
  * [func (i IntMap) Iterate() &lt;-chan collections.Element](#IntMap.Iterate)
  * [func (i IntMap) MostCommon(n int) []collections.Element](#IntMap.MostCommon)
  * [func (i IntMap) String() string](#IntMap.String)
  * [func (i IntMap) Subtract(key collections.Data)](#IntMap.Subtract)
  * [func (i IntMap) Update(key collections.Data)](#IntMap.Update)
* [type IntMap32](#IntMap32)
  * [func NewIntMap32(hash map[int32]int) IntMap32](#NewIntMap32)
  * [func (i IntMap32) Delete(key collections.Data)](#IntMap32.Delete)
  * [func (i IntMap32) Get(key collections.Data) (collections.Data, bool)](#IntMap32.Get)
  * [func (i IntMap32) Items() []collections.Element](#IntMap32.Items)
  * [func (i IntMap32) Iterate() &lt;-chan collections.Element](#IntMap32.Iterate)
  * [func (i IntMap32) MostCommon(n int) []collections.Element](#IntMap32.MostCommon)
  * [func (i IntMap32) String() string](#IntMap32.String)
  * [func (i IntMap32) Subtract(key collections.Data)](#IntMap32.Subtract)
  * [func (i IntMap32) Update(key collections.Data)](#IntMap32.Update)
* [type IntMap64](#IntMap64)
  * [func NewIntMap64(hash map[int64]int) IntMap64](#NewIntMap64)
  * [func (i IntMap64) Delete(key collections.Data)](#IntMap64.Delete)
  * [func (i IntMap64) Get(key collections.Data) (collections.Data, bool)](#IntMap64.Get)
  * [func (i IntMap64) Items() []collections.Element](#IntMap64.Items)
  * [func (i IntMap64) Iterate() &lt;-chan collections.Element](#IntMap64.Iterate)
  * [func (i IntMap64) MostCommon(n int) []collections.Element](#IntMap64.MostCommon)
  * [func (i IntMap64) String() string](#IntMap64.String)
  * [func (i IntMap64) Subtract(key collections.Data)](#IntMap64.Subtract)
  * [func (i IntMap64) Update(key collections.Data)](#IntMap64.Update)
* [type StringMap](#StringMap)
  * [func NewStringMap(hash map[string]int) StringMap](#NewStringMap)
  * [func (i StringMap) Delete(key collections.Data)](#StringMap.Delete)
  * [func (i StringMap) Get(key collections.Data) (collections.Data, bool)](#StringMap.Get)
  * [func (i StringMap) Items() []collections.Element](#StringMap.Items)
  * [func (i StringMap) Iterate() &lt;-chan collections.Element](#StringMap.Iterate)
  * [func (i StringMap) MostCommon(n int) []collections.Element](#StringMap.MostCommon)
  * [func (i StringMap) String() string](#StringMap.String)
  * [func (i StringMap) Subtract(key collections.Data)](#StringMap.Subtract)
  * [func (i StringMap) Update(key collections.Data)](#StringMap.Update)


#### <a name="pkg-files">Package files</a>
[counter.go](/src/github.com/marcsantiago/collections/counter/counter.go) [doc.go](/src/github.com/marcsantiago/collections/counter/doc.go) [float32_map.go](/src/github.com/marcsantiago/collections/counter/float32_map.go) [float64_map.go](/src/github.com/marcsantiago/collections/counter/float64_map.go) [int32_map.go](/src/github.com/marcsantiago/collections/counter/int32_map.go) [int64_map.go](/src/github.com/marcsantiago/collections/counter/int64_map.go) [int_map.go](/src/github.com/marcsantiago/collections/counter/int_map.go) [string_map.go](/src/github.com/marcsantiago/collections/counter/string_map.go) 





## <a name="Counter">func</a> [Counter](/src/target/counter.go?s=408:502#L13)
``` go
func Counter(data []collections.Data, optionalType ...collections.Type) collections.CounterMap
```
Counter is a python like abstraction on counting slice data
passing in a single optionalType for type means zero type reflection is needed
while collection.Data is an interface, it is expected that all elements are of the same primitive underneath
e.g []int, []int32, []int64, []float32, []float64, []bool, []string




## <a name="FloatMap32">type</a> [FloatMap32](/src/target/float32_map.go?s=105:136#L11)
``` go
type FloatMap32 map[float32]int
```






### <a name="NewFloatMap32">func</a> [NewFloatMap32](/src/target/float32_map.go?s=240:291#L14)
``` go
func NewFloatMap32(hash map[float32]int) FloatMap32
```
NewFloatMap32 takes in hash data, counts, and returns a concrete type that implements a CounterMap





### <a name="FloatMap32.Delete">func</a> (FloatMap32) [Delete](/src/target/float32_map.go?s=1093:1141#L46)
``` go
func (i FloatMap32) Delete(key collections.Data)
```
Delete removes the element from the internal map




### <a name="FloatMap32.Get">func</a> (FloatMap32) [Get](/src/target/float32_map.go?s=563:633#L27)
``` go
func (i FloatMap32) Get(key collections.Data) (collections.Data, bool)
```
Get retrieves a data value from the internal map if it exists




### <a name="FloatMap32.Items">func</a> (FloatMap32) [Items](/src/target/float32_map.go?s=1228:1277#L51)
``` go
func (i FloatMap32) Items() []collections.Element
```
Items returns the internal map as a set of elements




### <a name="FloatMap32.Iterate">func</a> (FloatMap32) [Iterate](/src/target/float32_map.go?s=1585:1641#L63)
``` go
func (i FloatMap32) Iterate() <-chan collections.Element
```
Iterate creates a channel to create an iterator for he Go range statement




### <a name="FloatMap32.MostCommon">func</a> (FloatMap32) [MostCommon](/src/target/float32_map.go?s=1915:1974#L75)
``` go
func (i FloatMap32) MostCommon(n int) []collections.Element
```
MostCommon returns the most common values by value




### <a name="FloatMap32.String">func</a> (FloatMap32) [String](/src/target/float32_map.go?s=2473:2508#L93)
``` go
func (i FloatMap32) String() string
```
String returns the JSON string representation of the map data
maps with float keys are not valid JSON, so the keys are converted to strings




### <a name="FloatMap32.Subtract">func</a> (FloatMap32) [Subtract](/src/target/float32_map.go?s=925:975#L38)
``` go
func (i FloatMap32) Subtract(key collections.Data)
```
Subtract removes 1 from the counter if the key exists




### <a name="FloatMap32.Update">func</a> (FloatMap32) [Update](/src/target/float32_map.go?s=794:842#L33)
``` go
func (i FloatMap32) Update(key collections.Data)
```
Update updates the counter for a value or sets the value it if it does not exist




## <a name="FloatMap64">type</a> [FloatMap64](/src/target/float64_map.go?s=105:136#L11)
``` go
type FloatMap64 map[float64]int
```






### <a name="NewFloatMap64">func</a> [NewFloatMap64](/src/target/float64_map.go?s=240:291#L14)
``` go
func NewFloatMap64(hash map[float64]int) FloatMap64
```
NewFloatMap64 takes in hash data, counts, and returns a concrete type that implements a CounterMap





### <a name="FloatMap64.Delete">func</a> (FloatMap64) [Delete](/src/target/float64_map.go?s=1093:1141#L46)
``` go
func (i FloatMap64) Delete(key collections.Data)
```
Delete removes the element from the internal map




### <a name="FloatMap64.Get">func</a> (FloatMap64) [Get](/src/target/float64_map.go?s=563:633#L27)
``` go
func (i FloatMap64) Get(key collections.Data) (collections.Data, bool)
```
Get retrieves a data value from the internal map if it exists




### <a name="FloatMap64.Items">func</a> (FloatMap64) [Items](/src/target/float64_map.go?s=1228:1277#L51)
``` go
func (i FloatMap64) Items() []collections.Element
```
Items returns the internal map as a set of elements




### <a name="FloatMap64.Iterate">func</a> (FloatMap64) [Iterate](/src/target/float64_map.go?s=1585:1641#L63)
``` go
func (i FloatMap64) Iterate() <-chan collections.Element
```
Iterate creates a channel to create an iterator for he Go range statement




### <a name="FloatMap64.MostCommon">func</a> (FloatMap64) [MostCommon](/src/target/float64_map.go?s=1915:1974#L75)
``` go
func (i FloatMap64) MostCommon(n int) []collections.Element
```
MostCommon returns the most common values by value




### <a name="FloatMap64.String">func</a> (FloatMap64) [String](/src/target/float64_map.go?s=2473:2508#L93)
``` go
func (i FloatMap64) String() string
```
String returns the JSON string representation of the map data
maps with float keys are not valid JSON, so the keys are converted to strings




### <a name="FloatMap64.Subtract">func</a> (FloatMap64) [Subtract](/src/target/float64_map.go?s=925:975#L38)
``` go
func (i FloatMap64) Subtract(key collections.Data)
```
Subtract removes 1 from the counter if the key exists




### <a name="FloatMap64.Update">func</a> (FloatMap64) [Update](/src/target/float64_map.go?s=794:842#L33)
``` go
func (i FloatMap64) Update(key collections.Data)
```
Update updates the counter for a value or sets the value it if it does not exist




## <a name="IntMap">type</a> [IntMap](/src/target/int_map.go?s=94:117#L10)
``` go
type IntMap map[int]int
```






### <a name="NewIntMap">func</a> [NewIntMap](/src/target/int_map.go?s=217:256#L13)
``` go
func NewIntMap(hash map[int]int) IntMap
```
NewIntMap takes in hash data, counts, and returns a concrete type that implements a CounterMap





### <a name="IntMap.Delete">func</a> (IntMap) [Delete](/src/target/int_map.go?s=1010:1054#L45)
``` go
func (i IntMap) Delete(key collections.Data)
```
Delete removes the element from the internal map




### <a name="IntMap.Get">func</a> (IntMap) [Get](/src/target/int_map.go?s=512:578#L26)
``` go
func (i IntMap) Get(key collections.Data) (collections.Data, bool)
```
Get retrieves a data value from the internal map if it exists




### <a name="IntMap.Items">func</a> (IntMap) [Items](/src/target/int_map.go?s=1137:1182#L50)
``` go
func (i IntMap) Items() []collections.Element
```
Items returns the internal map as a set of elements




### <a name="IntMap.Iterate">func</a> (IntMap) [Iterate](/src/target/int_map.go?s=1486:1538#L62)
``` go
func (i IntMap) Iterate() <-chan collections.Element
```
Iterate creates a channel to create an iterator for he Go range statement




### <a name="IntMap.MostCommon">func</a> (IntMap) [MostCommon](/src/target/int_map.go?s=1808:1863#L74)
``` go
func (i IntMap) MostCommon(n int) []collections.Element
```
MostCommon returns the most common values by value




### <a name="IntMap.String">func</a> (IntMap) [String](/src/target/int_map.go?s=2277:2308#L91)
``` go
func (i IntMap) String() string
```
String returns the JSON string representation of the map data




### <a name="IntMap.Subtract">func</a> (IntMap) [Subtract](/src/target/int_map.go?s=854:900#L37)
``` go
func (i IntMap) Subtract(key collections.Data)
```
Subtract removes 1 from the counter if the key exists




### <a name="IntMap.Update">func</a> (IntMap) [Update](/src/target/int_map.go?s=731:775#L32)
``` go
func (i IntMap) Update(key collections.Data)
```
Update updates the counter for a value or sets the value it if it does not exist




## <a name="IntMap32">type</a> [IntMap32](/src/target/int32_map.go?s=94:121#L10)
``` go
type IntMap32 map[int32]int
```






### <a name="NewIntMap32">func</a> [NewIntMap32](/src/target/int32_map.go?s=223:268#L13)
``` go
func NewIntMap32(hash map[int32]int) IntMap32
```
NewIntMap32 takes in hash data, counts, and returns a concrete type that implements a CounterMap





### <a name="IntMap32.Delete">func</a> (IntMap32) [Delete](/src/target/int32_map.go?s=1046:1092#L45)
``` go
func (i IntMap32) Delete(key collections.Data)
```
Delete removes the element from the internal map




### <a name="IntMap32.Get">func</a> (IntMap32) [Get](/src/target/int32_map.go?s=532:600#L26)
``` go
func (i IntMap32) Get(key collections.Data) (collections.Data, bool)
```
Get retrieves a data value from the internal map if it exists




### <a name="IntMap32.Items">func</a> (IntMap32) [Items](/src/target/int32_map.go?s=1177:1224#L50)
``` go
func (i IntMap32) Items() []collections.Element
```
Items returns the internal map as a set of elements




### <a name="IntMap32.Iterate">func</a> (IntMap32) [Iterate](/src/target/int32_map.go?s=1530:1584#L62)
``` go
func (i IntMap32) Iterate() <-chan collections.Element
```
Iterate creates a channel to create an iterator for he Go range statement




### <a name="IntMap32.MostCommon">func</a> (IntMap32) [MostCommon](/src/target/int32_map.go?s=1856:1913#L74)
``` go
func (i IntMap32) MostCommon(n int) []collections.Element
```
MostCommon returns the most common values by value




### <a name="IntMap32.String">func</a> (IntMap32) [String](/src/target/int32_map.go?s=2329:2362#L91)
``` go
func (i IntMap32) String() string
```
String returns the JSON string representation of the map data




### <a name="IntMap32.Subtract">func</a> (IntMap32) [Subtract](/src/target/int32_map.go?s=884:932#L37)
``` go
func (i IntMap32) Subtract(key collections.Data)
```
Subtract removes 1 from the counter if the key exists




### <a name="IntMap32.Update">func</a> (IntMap32) [Update](/src/target/int32_map.go?s=757:803#L32)
``` go
func (i IntMap32) Update(key collections.Data)
```
Update updates the counter for a value or sets the value it if it does not exist




## <a name="IntMap64">type</a> [IntMap64](/src/target/int64_map.go?s=94:121#L10)
``` go
type IntMap64 map[int64]int
```






### <a name="NewIntMap64">func</a> [NewIntMap64](/src/target/int64_map.go?s=223:268#L13)
``` go
func NewIntMap64(hash map[int64]int) IntMap64
```
NewIntMap64 takes in hash data, counts, and returns a concrete type that implements a CounterMap





### <a name="IntMap64.Delete">func</a> (IntMap64) [Delete](/src/target/int64_map.go?s=1046:1092#L45)
``` go
func (i IntMap64) Delete(key collections.Data)
```
Delete removes the element from the internal map




### <a name="IntMap64.Get">func</a> (IntMap64) [Get](/src/target/int64_map.go?s=532:600#L26)
``` go
func (i IntMap64) Get(key collections.Data) (collections.Data, bool)
```
Get retrieves a data value from the internal map if it exists




### <a name="IntMap64.Items">func</a> (IntMap64) [Items](/src/target/int64_map.go?s=1177:1224#L50)
``` go
func (i IntMap64) Items() []collections.Element
```
Items returns the internal map as a set of elements




### <a name="IntMap64.Iterate">func</a> (IntMap64) [Iterate](/src/target/int64_map.go?s=1530:1584#L62)
``` go
func (i IntMap64) Iterate() <-chan collections.Element
```
Iterate creates a channel to create an iterator for he Go range statement




### <a name="IntMap64.MostCommon">func</a> (IntMap64) [MostCommon](/src/target/int64_map.go?s=1856:1913#L74)
``` go
func (i IntMap64) MostCommon(n int) []collections.Element
```
MostCommon returns the most common values by value




### <a name="IntMap64.String">func</a> (IntMap64) [String](/src/target/int64_map.go?s=2329:2362#L91)
``` go
func (i IntMap64) String() string
```
String returns the JSON string representation of the map data




### <a name="IntMap64.Subtract">func</a> (IntMap64) [Subtract](/src/target/int64_map.go?s=884:932#L37)
``` go
func (i IntMap64) Subtract(key collections.Data)
```
Subtract removes 1 from the counter if the key exists




### <a name="IntMap64.Update">func</a> (IntMap64) [Update](/src/target/int64_map.go?s=757:803#L32)
``` go
func (i IntMap64) Update(key collections.Data)
```
Update updates the counter for a value or sets the value it if it does not exist




## <a name="StringMap">type</a> [StringMap](/src/target/string_map.go?s=94:123#L10)
``` go
type StringMap map[string]int
```






### <a name="NewStringMap">func</a> [NewStringMap](/src/target/string_map.go?s=226:274#L13)
``` go
func NewStringMap(hash map[string]int) StringMap
```
NewStringMap takes in hash data, counts, and returns a concrete type that implements a CounterMap





### <a name="StringMap.Delete">func</a> (StringMap) [Delete](/src/target/string_map.go?s=1064:1111#L45)
``` go
func (i StringMap) Delete(key collections.Data)
```
Delete removes the element from the internal map




### <a name="StringMap.Get">func</a> (StringMap) [Get](/src/target/string_map.go?s=542:611#L26)
``` go
func (i StringMap) Get(key collections.Data) (collections.Data, bool)
```
Get retrieves a data value from the internal map if it exists




### <a name="StringMap.Items">func</a> (StringMap) [Items](/src/target/string_map.go?s=1197:1245#L50)
``` go
func (i StringMap) Items() []collections.Element
```
Items returns the internal map as a set of elements




### <a name="StringMap.Iterate">func</a> (StringMap) [Iterate](/src/target/string_map.go?s=1552:1607#L62)
``` go
func (i StringMap) Iterate() <-chan collections.Element
```
Iterate creates a channel to create an iterator for he Go range statement




### <a name="StringMap.MostCommon">func</a> (StringMap) [MostCommon](/src/target/string_map.go?s=1880:1938#L74)
``` go
func (i StringMap) MostCommon(n int) []collections.Element
```
MostCommon returns the most common values by value




### <a name="StringMap.String">func</a> (StringMap) [String](/src/target/string_map.go?s=2355:2389#L91)
``` go
func (i StringMap) String() string
```
String returns the JSON string representation of the map data




### <a name="StringMap.Subtract">func</a> (StringMap) [Subtract](/src/target/string_map.go?s=899:948#L37)
``` go
func (i StringMap) Subtract(key collections.Data)
```
Subtract removes 1 from the counter if the key exists




### <a name="StringMap.Update">func</a> (StringMap) [Update](/src/target/string_map.go?s=770:817#L32)
``` go
func (i StringMap) Update(key collections.Data)
```
Update updates the counter for a value or sets the value it if it does not exist








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
