

# collections
`import "github.com/marcsantiago/collections"`
### [![Test Status](https://github.com/marcsantiago/collections/workflows/Go/badge.svg)](go)
#
parody of some of the basic python core features

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [type CounterMap](#CounterMap)
* [type Data](#Data)
* [type Element](#Element)
* [type ElementsByKeyIntAsc](#ElementsByKeyIntAsc)
  * [func (e ElementsByKeyIntAsc) Len() int](#ElementsByKeyIntAsc.Len)
  * [func (e ElementsByKeyIntAsc) Less(i, j int) bool](#ElementsByKeyIntAsc.Less)
  * [func (e ElementsByKeyIntAsc) Swap(i, j int)](#ElementsByKeyIntAsc.Swap)
* [type ElementsByKeyIntDesc](#ElementsByKeyIntDesc)
  * [func (e ElementsByKeyIntDesc) Len() int](#ElementsByKeyIntDesc.Len)
  * [func (e ElementsByKeyIntDesc) Less(i, j int) bool](#ElementsByKeyIntDesc.Less)
  * [func (e ElementsByKeyIntDesc) Swap(i, j int)](#ElementsByKeyIntDesc.Swap)
* [type ElementsByKeyStringAsc](#ElementsByKeyStringAsc)
  * [func (e ElementsByKeyStringAsc) Len() int](#ElementsByKeyStringAsc.Len)
  * [func (e ElementsByKeyStringAsc) Less(i, j int) bool](#ElementsByKeyStringAsc.Less)
  * [func (e ElementsByKeyStringAsc) Swap(i, j int)](#ElementsByKeyStringAsc.Swap)
* [type ElementsByKeyStringDesc](#ElementsByKeyStringDesc)
  * [func (e ElementsByKeyStringDesc) Len() int](#ElementsByKeyStringDesc.Len)
  * [func (e ElementsByKeyStringDesc) Less(i, j int) bool](#ElementsByKeyStringDesc.Less)
  * [func (e ElementsByKeyStringDesc) Swap(i, j int)](#ElementsByKeyStringDesc.Swap)
* [type ElementsByValueIntAsc](#ElementsByValueIntAsc)
  * [func (e ElementsByValueIntAsc) Len() int](#ElementsByValueIntAsc.Len)
  * [func (e ElementsByValueIntAsc) Less(i, j int) bool](#ElementsByValueIntAsc.Less)
  * [func (e ElementsByValueIntAsc) Swap(i, j int)](#ElementsByValueIntAsc.Swap)
* [type ElementsByValueIntDesc](#ElementsByValueIntDesc)
  * [func (e ElementsByValueIntDesc) Len() int](#ElementsByValueIntDesc.Len)
  * [func (e ElementsByValueIntDesc) Less(i, j int) bool](#ElementsByValueIntDesc.Less)
  * [func (e ElementsByValueIntDesc) Swap(i, j int)](#ElementsByValueIntDesc.Swap)
* [type ElementsByValueStringAsc](#ElementsByValueStringAsc)
  * [func (e ElementsByValueStringAsc) Len() int](#ElementsByValueStringAsc.Len)
  * [func (e ElementsByValueStringAsc) Less(i, j int) bool](#ElementsByValueStringAsc.Less)
  * [func (e ElementsByValueStringAsc) Swap(i, j int)](#ElementsByValueStringAsc.Swap)
* [type ElementsByValueStringDesc](#ElementsByValueStringDesc)
  * [func (e ElementsByValueStringDesc) Len() int](#ElementsByValueStringDesc.Len)
  * [func (e ElementsByValueStringDesc) Less(i, j int) bool](#ElementsByValueStringDesc.Less)
  * [func (e ElementsByValueStringDesc) Swap(i, j int)](#ElementsByValueStringDesc.Swap)
* [type FloatValue32](#FloatValue32)
  * [func (i FloatValue32) Float32() float32](#FloatValue32.Float32)
  * [func (i FloatValue32) Float64() float64](#FloatValue32.Float64)
  * [func (i FloatValue32) Int() int](#FloatValue32.Int)
  * [func (i FloatValue32) Int32() int32](#FloatValue32.Int32)
  * [func (i FloatValue32) Int64() int64](#FloatValue32.Int64)
  * [func (i FloatValue32) String() string](#FloatValue32.String)
* [type FloatValue64](#FloatValue64)
  * [func (i FloatValue64) Float32() float32](#FloatValue64.Float32)
  * [func (i FloatValue64) Float64() float64](#FloatValue64.Float64)
  * [func (i FloatValue64) Int() int](#FloatValue64.Int)
  * [func (i FloatValue64) Int32() int32](#FloatValue64.Int32)
  * [func (i FloatValue64) Int64() int64](#FloatValue64.Int64)
  * [func (i FloatValue64) String() string](#FloatValue64.String)
* [type FloatValues32](#FloatValues32)
  * [func (i FloatValues32) Data() []Data](#FloatValues32.Data)
* [type FloatValues64](#FloatValues64)
  * [func (i FloatValues64) Data() []Data](#FloatValues64.Data)
* [type IntValue](#IntValue)
  * [func (i IntValue) Float32() float32](#IntValue.Float32)
  * [func (i IntValue) Float64() float64](#IntValue.Float64)
  * [func (i IntValue) Int() int](#IntValue.Int)
  * [func (i IntValue) Int32() int32](#IntValue.Int32)
  * [func (i IntValue) Int64() int64](#IntValue.Int64)
  * [func (i IntValue) String() string](#IntValue.String)
* [type IntValue32](#IntValue32)
  * [func (i IntValue32) Float32() float32](#IntValue32.Float32)
  * [func (i IntValue32) Float64() float64](#IntValue32.Float64)
  * [func (i IntValue32) Int() int](#IntValue32.Int)
  * [func (i IntValue32) Int32() int32](#IntValue32.Int32)
  * [func (i IntValue32) Int64() int64](#IntValue32.Int64)
  * [func (i IntValue32) String() string](#IntValue32.String)
* [type IntValue64](#IntValue64)
  * [func (i IntValue64) Float32() float32](#IntValue64.Float32)
  * [func (i IntValue64) Float64() float64](#IntValue64.Float64)
  * [func (i IntValue64) Int() int](#IntValue64.Int)
  * [func (i IntValue64) Int32() int32](#IntValue64.Int32)
  * [func (i IntValue64) Int64() int64](#IntValue64.Int64)
  * [func (i IntValue64) String() string](#IntValue64.String)
* [type IntValues](#IntValues)
  * [func (i IntValues) Data() []Data](#IntValues.Data)
* [type IntValues32](#IntValues32)
  * [func (i IntValues32) Data() []Data](#IntValues32.Data)
* [type IntValues64](#IntValues64)
  * [func (i IntValues64) Data() []Data](#IntValues64.Data)
* [type Iterable](#Iterable)
* [type StringValue](#StringValue)
  * [func (s StringValue) Float32() float32](#StringValue.Float32)
  * [func (s StringValue) Float64() float64](#StringValue.Float64)
  * [func (s StringValue) Int() int](#StringValue.Int)
  * [func (s StringValue) Int32() int32](#StringValue.Int32)
  * [func (s StringValue) Int64() int64](#StringValue.Int64)
  * [func (s StringValue) String() string](#StringValue.String)
* [type StringValues](#StringValues)
  * [func (s StringValues) Data() []Data](#StringValues.Data)
* [type Type](#Type)


#### <a name="pkg-files">Package files</a>
[data.go](/src/github.com/marcsantiago/collections/data.go) [element_sorters.go](/src/github.com/marcsantiago/collections/element_sorters.go) [interable.go](/src/github.com/marcsantiago/collections/interable.go) [map.go](/src/github.com/marcsantiago/collections/map.go) [primitive_conversions.go](/src/github.com/marcsantiago/collections/primitive_conversions.go) [types.go](/src/github.com/marcsantiago/collections/types.go) 






## <a name="CounterMap">type</a> [CounterMap](/src/target/map.go?s=73:267#L4)
``` go
type CounterMap interface {
    Iterable
    Delete(value Data)
    Get(key Data) (Data, bool)
    Items() []Element
    MostCommon(n int) []Element
    String() string
    Subtract(value Data)
    Update(value Data)
}
```
CounterMap mimics the Python Counter definitions










## <a name="Data">type</a> [Data](/src/target/data.go?s=100:219#L4)
``` go
type Data interface {
    Int() int
    Int32() int32
    Int64() int64
    Float32() float32
    Float64() float64
    String() string
}
```
Data interface for casting and getting data values used to avoid reflection










## <a name="Element">type</a> [Element](/src/target/element_sorters.go?s=90:137#L4)
``` go
type Element struct {
    Key   Data
    Value Data
}
```
Element represents (key, value) with backed by the Data interface










## <a name="ElementsByKeyIntAsc">type</a> [ElementsByKeyIntAsc](/src/target/element_sorters.go?s=960:994#L24)
``` go
type ElementsByKeyIntAsc []Element
```
ElementsByKeyIntAsc used to used low to high where key is of type string










### <a name="ElementsByKeyIntAsc.Len">func</a> (ElementsByKeyIntAsc) [Len](/src/target/element_sorters.go?s=996:1034#L26)
``` go
func (e ElementsByKeyIntAsc) Len() int
```



### <a name="ElementsByKeyIntAsc.Less">func</a> (ElementsByKeyIntAsc) [Less](/src/target/element_sorters.go?s=1063:1111#L27)
``` go
func (e ElementsByKeyIntAsc) Less(i, j int) bool
```



### <a name="ElementsByKeyIntAsc.Swap">func</a> (ElementsByKeyIntAsc) [Swap](/src/target/element_sorters.go?s=1155:1198#L28)
``` go
func (e ElementsByKeyIntAsc) Swap(i, j int)
```



## <a name="ElementsByKeyIntDesc">type</a> [ElementsByKeyIntDesc](/src/target/element_sorters.go?s=1310:1345#L31)
``` go
type ElementsByKeyIntDesc []Element
```
ElementsByKeyIntDesc used to used high to low where key is of type string










### <a name="ElementsByKeyIntDesc.Len">func</a> (ElementsByKeyIntDesc) [Len](/src/target/element_sorters.go?s=1347:1386#L33)
``` go
func (e ElementsByKeyIntDesc) Len() int
```



### <a name="ElementsByKeyIntDesc.Less">func</a> (ElementsByKeyIntDesc) [Less](/src/target/element_sorters.go?s=1415:1464#L34)
``` go
func (e ElementsByKeyIntDesc) Less(i, j int) bool
```



### <a name="ElementsByKeyIntDesc.Swap">func</a> (ElementsByKeyIntDesc) [Swap](/src/target/element_sorters.go?s=1508:1552#L35)
``` go
func (e ElementsByKeyIntDesc) Swap(i, j int)
```



## <a name="ElementsByKeyStringAsc">type</a> [ElementsByKeyStringAsc](/src/target/element_sorters.go?s=218:255#L10)
``` go
type ElementsByKeyStringAsc []Element
```
ElementsByKeyStringAsc used to used low to high where key is of type string










### <a name="ElementsByKeyStringAsc.Len">func</a> (ElementsByKeyStringAsc) [Len](/src/target/element_sorters.go?s=257:298#L12)
``` go
func (e ElementsByKeyStringAsc) Len() int
```



### <a name="ElementsByKeyStringAsc.Less">func</a> (ElementsByKeyStringAsc) [Less](/src/target/element_sorters.go?s=327:378#L13)
``` go
func (e ElementsByKeyStringAsc) Less(i, j int) bool
```



### <a name="ElementsByKeyStringAsc.Swap">func</a> (ElementsByKeyStringAsc) [Swap](/src/target/element_sorters.go?s=428:474#L14)
``` go
func (e ElementsByKeyStringAsc) Swap(i, j int)
```



## <a name="ElementsByKeyStringDesc">type</a> [ElementsByKeyStringDesc](/src/target/element_sorters.go?s=589:627#L17)
``` go
type ElementsByKeyStringDesc []Element
```
ElementsByKeyStringDesc used to used high to low where key is of type string










### <a name="ElementsByKeyStringDesc.Len">func</a> (ElementsByKeyStringDesc) [Len](/src/target/element_sorters.go?s=629:671#L19)
``` go
func (e ElementsByKeyStringDesc) Len() int
```



### <a name="ElementsByKeyStringDesc.Less">func</a> (ElementsByKeyStringDesc) [Less](/src/target/element_sorters.go?s=700:752#L20)
``` go
func (e ElementsByKeyStringDesc) Less(i, j int) bool
```



### <a name="ElementsByKeyStringDesc.Swap">func</a> (ElementsByKeyStringDesc) [Swap](/src/target/element_sorters.go?s=802:849#L21)
``` go
func (e ElementsByKeyStringDesc) Swap(i, j int)
```



## <a name="ElementsByValueIntAsc">type</a> [ElementsByValueIntAsc](/src/target/element_sorters.go?s=2414:2450#L56)
``` go
type ElementsByValueIntAsc []Element
```
ElementsByValueIntAsc used to used low to high where value is of type string










### <a name="ElementsByValueIntAsc.Len">func</a> (ElementsByValueIntAsc) [Len](/src/target/element_sorters.go?s=2452:2492#L58)
``` go
func (e ElementsByValueIntAsc) Len() int
```



### <a name="ElementsByValueIntAsc.Less">func</a> (ElementsByValueIntAsc) [Less](/src/target/element_sorters.go?s=2521:2571#L59)
``` go
func (e ElementsByValueIntAsc) Less(i, j int) bool
```



### <a name="ElementsByValueIntAsc.Swap">func</a> (ElementsByValueIntAsc) [Swap](/src/target/element_sorters.go?s=2619:2664#L60)
``` go
func (e ElementsByValueIntAsc) Swap(i, j int)
```



## <a name="ElementsByValueIntDesc">type</a> [ElementsByValueIntDesc](/src/target/element_sorters.go?s=2780:2817#L63)
``` go
type ElementsByValueIntDesc []Element
```
ElementsByValueIntDesc used to used high to low where value is of type string










### <a name="ElementsByValueIntDesc.Len">func</a> (ElementsByValueIntDesc) [Len](/src/target/element_sorters.go?s=2819:2860#L65)
``` go
func (e ElementsByValueIntDesc) Len() int
```



### <a name="ElementsByValueIntDesc.Less">func</a> (ElementsByValueIntDesc) [Less](/src/target/element_sorters.go?s=2889:2940#L66)
``` go
func (e ElementsByValueIntDesc) Less(i, j int) bool
```



### <a name="ElementsByValueIntDesc.Swap">func</a> (ElementsByValueIntDesc) [Swap](/src/target/element_sorters.go?s=2988:3034#L67)
``` go
func (e ElementsByValueIntDesc) Swap(i, j int)
```



## <a name="ElementsByValueStringAsc">type</a> [ElementsByValueStringAsc](/src/target/element_sorters.go?s=1668:1707#L38)
``` go
type ElementsByValueStringAsc []Element
```
ElementsByKeyStringAsc used to used low to high where value is of type string










### <a name="ElementsByValueStringAsc.Len">func</a> (ElementsByValueStringAsc) [Len](/src/target/element_sorters.go?s=1709:1752#L40)
``` go
func (e ElementsByValueStringAsc) Len() int
```



### <a name="ElementsByValueStringAsc.Less">func</a> (ElementsByValueStringAsc) [Less](/src/target/element_sorters.go?s=1771:1824#L41)
``` go
func (e ElementsByValueStringAsc) Less(i, j int) bool
```



### <a name="ElementsByValueStringAsc.Swap">func</a> (ElementsByValueStringAsc) [Swap](/src/target/element_sorters.go?s=1879:1927#L44)
``` go
func (e ElementsByValueStringAsc) Swap(i, j int)
```



## <a name="ElementsByValueStringDesc">type</a> [ElementsByValueStringDesc](/src/target/element_sorters.go?s=2041:2081#L47)
``` go
type ElementsByValueStringDesc []Element
```
ElementsByValueStringDesc used to used high to low where value is of type string










### <a name="ElementsByValueStringDesc.Len">func</a> (ElementsByValueStringDesc) [Len](/src/target/element_sorters.go?s=2083:2127#L49)
``` go
func (e ElementsByValueStringDesc) Len() int
```



### <a name="ElementsByValueStringDesc.Less">func</a> (ElementsByValueStringDesc) [Less](/src/target/element_sorters.go?s=2146:2200#L50)
``` go
func (e ElementsByValueStringDesc) Less(i, j int) bool
```



### <a name="ElementsByValueStringDesc.Swap">func</a> (ElementsByValueStringDesc) [Swap](/src/target/element_sorters.go?s=2255:2304#L53)
``` go
func (e ElementsByValueStringDesc) Swap(i, j int)
```



## <a name="FloatValue32">type</a> [FloatValue32](/src/target/primitive_conversions.go?s=2608:2633#L138)
``` go
type FloatValue32 float32
```
FloatValue32 type alias for int










### <a name="FloatValue32.Float32">func</a> (FloatValue32) [Float32](/src/target/primitive_conversions.go?s=2959:2998#L156)
``` go
func (i FloatValue32) Float32() float32
```
Float32 casts and returns Float32 value




### <a name="FloatValue32.Float64">func</a> (FloatValue32) [Float64](/src/target/primitive_conversions.go?s=3066:3105#L161)
``` go
func (i FloatValue32) Float64() float64
```
Float64 casts and returns Float64 value




### <a name="FloatValue32.Int">func</a> (FloatValue32) [Int](/src/target/primitive_conversions.go?s=2670:2701#L141)
``` go
func (i FloatValue32) Int() int
```
Int casts and returns int value




### <a name="FloatValue32.Int32">func</a> (FloatValue32) [Int32](/src/target/primitive_conversions.go?s=2761:2796#L146)
``` go
func (i FloatValue32) Int32() int32
```
Int32 casts and returns Int32 value




### <a name="FloatValue32.Int64">func</a> (FloatValue32) [Int64](/src/target/primitive_conversions.go?s=2858:2893#L151)
``` go
func (i FloatValue32) Int64() int64
```
Int64 casts and returns Int64 value




### <a name="FloatValue32.String">func</a> (FloatValue32) [String](/src/target/primitive_conversions.go?s=3171:3208#L166)
``` go
func (i FloatValue32) String() string
```
String casts and returns string value




## <a name="FloatValue64">type</a> [FloatValue64](/src/target/primitive_conversions.go?s=3514:3539#L182)
``` go
type FloatValue64 float64
```
FloatValue64 type alias for int










### <a name="FloatValue64.Float32">func</a> (FloatValue64) [Float32](/src/target/primitive_conversions.go?s=3865:3904#L200)
``` go
func (i FloatValue64) Float32() float32
```
Float32 casts and returns Float32 value




### <a name="FloatValue64.Float64">func</a> (FloatValue64) [Float64](/src/target/primitive_conversions.go?s=3972:4011#L205)
``` go
func (i FloatValue64) Float64() float64
```
Float64 casts and returns Float64 value




### <a name="FloatValue64.Int">func</a> (FloatValue64) [Int](/src/target/primitive_conversions.go?s=3576:3607#L185)
``` go
func (i FloatValue64) Int() int
```
Int casts and returns int value




### <a name="FloatValue64.Int32">func</a> (FloatValue64) [Int32](/src/target/primitive_conversions.go?s=3667:3702#L190)
``` go
func (i FloatValue64) Int32() int32
```
Int32 casts and returns Int32 value




### <a name="FloatValue64.Int64">func</a> (FloatValue64) [Int64](/src/target/primitive_conversions.go?s=3764:3799#L195)
``` go
func (i FloatValue64) Int64() int64
```
Int64 casts and returns Int64 value




### <a name="FloatValue64.String">func</a> (FloatValue64) [String](/src/target/primitive_conversions.go?s=4077:4114#L210)
``` go
func (i FloatValue64) String() string
```
String casts and returns string value




## <a name="FloatValues32">type</a> [FloatValues32](/src/target/primitive_conversions.go?s=3319:3347#L171)
``` go
type FloatValues32 []float32
```
FloatValues32 type alias for a slice of IntValue










### <a name="FloatValues32.Data">func</a> (FloatValues32) [Data](/src/target/primitive_conversions.go?s=3349:3385#L173)
``` go
func (i FloatValues32) Data() []Data
```



## <a name="FloatValues64">type</a> [FloatValues64](/src/target/primitive_conversions.go?s=4225:4253#L215)
``` go
type FloatValues64 []float32
```
FloatValues64 type alias for a slice of IntValue










### <a name="FloatValues64.Data">func</a> (FloatValues64) [Data](/src/target/primitive_conversions.go?s=4255:4291#L217)
``` go
func (i FloatValues64) Data() []Data
```



## <a name="IntValue">type</a> [IntValue](/src/target/primitive_conversions.go?s=70:87#L6)
``` go
type IntValue int
```
IntValue type alias for int










### <a name="IntValue.Float32">func</a> (IntValue) [Float32](/src/target/primitive_conversions.go?s=401:436#L24)
``` go
func (i IntValue) Float32() float32
```
Float32 casts and returns Float32 value




### <a name="IntValue.Float64">func</a> (IntValue) [Float64](/src/target/primitive_conversions.go?s=504:539#L29)
``` go
func (i IntValue) Float64() float64
```
Float64 casts and returns Float64 value




### <a name="IntValue.Int">func</a> (IntValue) [Int](/src/target/primitive_conversions.go?s=124:151#L9)
``` go
func (i IntValue) Int() int
```
Int casts and returns int value




### <a name="IntValue.Int32">func</a> (IntValue) [Int32](/src/target/primitive_conversions.go?s=211:242#L14)
``` go
func (i IntValue) Int32() int32
```
Int32 casts and returns Int32 value




### <a name="IntValue.Int64">func</a> (IntValue) [Int64](/src/target/primitive_conversions.go?s=304:335#L19)
``` go
func (i IntValue) Int64() int64
```
Int64 casts and returns Int64 value




### <a name="IntValue.String">func</a> (IntValue) [String](/src/target/primitive_conversions.go?s=605:638#L34)
``` go
func (i IntValue) String() string
```
String casts and returns string value




## <a name="IntValue32">type</a> [IntValue32](/src/target/primitive_conversions.go?s=898:919#L50)
``` go
type IntValue32 int32
```
IntValue32 type alias for int










### <a name="IntValue32.Float32">func</a> (IntValue32) [Float32](/src/target/primitive_conversions.go?s=1239:1276#L68)
``` go
func (i IntValue32) Float32() float32
```
Float32 casts and returns Float32 value




### <a name="IntValue32.Float64">func</a> (IntValue32) [Float64](/src/target/primitive_conversions.go?s=1344:1381#L73)
``` go
func (i IntValue32) Float64() float64
```
Float64 casts and returns Float64 value




### <a name="IntValue32.Int">func</a> (IntValue32) [Int](/src/target/primitive_conversions.go?s=956:985#L53)
``` go
func (i IntValue32) Int() int
```
Int casts and returns int value




### <a name="IntValue32.Int32">func</a> (IntValue32) [Int32](/src/target/primitive_conversions.go?s=1045:1078#L58)
``` go
func (i IntValue32) Int32() int32
```
Int32 casts and returns Int32 value




### <a name="IntValue32.Int64">func</a> (IntValue32) [Int64](/src/target/primitive_conversions.go?s=1140:1173#L63)
``` go
func (i IntValue32) Int64() int64
```
Int64 casts and returns Int64 value




### <a name="IntValue32.String">func</a> (IntValue32) [String](/src/target/primitive_conversions.go?s=1447:1482#L78)
``` go
func (i IntValue32) String() string
```
String casts and returns string value




## <a name="IntValue64">type</a> [IntValue64](/src/target/primitive_conversions.go?s=1752:1773#L94)
``` go
type IntValue64 int64
```
IntValue64 type alias for int










### <a name="IntValue64.Float32">func</a> (IntValue64) [Float32](/src/target/primitive_conversions.go?s=2093:2130#L112)
``` go
func (i IntValue64) Float32() float32
```
Float32 casts and returns Float32 value




### <a name="IntValue64.Float64">func</a> (IntValue64) [Float64](/src/target/primitive_conversions.go?s=2198:2235#L117)
``` go
func (i IntValue64) Float64() float64
```
Float64 casts and returns Float64 value




### <a name="IntValue64.Int">func</a> (IntValue64) [Int](/src/target/primitive_conversions.go?s=1810:1839#L97)
``` go
func (i IntValue64) Int() int
```
Int casts and returns int value




### <a name="IntValue64.Int32">func</a> (IntValue64) [Int32](/src/target/primitive_conversions.go?s=1899:1932#L102)
``` go
func (i IntValue64) Int32() int32
```
Int32 casts and returns Int32 value




### <a name="IntValue64.Int64">func</a> (IntValue64) [Int64](/src/target/primitive_conversions.go?s=1994:2027#L107)
``` go
func (i IntValue64) Int64() int64
```
Int64 casts and returns Int64 value




### <a name="IntValue64.String">func</a> (IntValue64) [String](/src/target/primitive_conversions.go?s=2301:2336#L122)
``` go
func (i IntValue64) String() string
```
String casts and returns string value




## <a name="IntValues">type</a> [IntValues](/src/target/primitive_conversions.go?s=721:741#L39)
``` go
type IntValues []int
```
IntValues type alias for a slice of IntValue










### <a name="IntValues.Data">func</a> (IntValues) [Data](/src/target/primitive_conversions.go?s=743:775#L41)
``` go
func (i IntValues) Data() []Data
```



## <a name="IntValues32">type</a> [IntValues32](/src/target/primitive_conversions.go?s=1567:1591#L83)
``` go
type IntValues32 []int32
```
IntValues32 type alias for a slice of IntValue










### <a name="IntValues32.Data">func</a> (IntValues32) [Data](/src/target/primitive_conversions.go?s=1593:1627#L85)
``` go
func (i IntValues32) Data() []Data
```



## <a name="IntValues64">type</a> [IntValues64](/src/target/primitive_conversions.go?s=2421:2445#L127)
``` go
type IntValues64 []int64
```
IntValues64 type alias for a slice of IntValue










### <a name="IntValues64.Data">func</a> (IntValues64) [Data](/src/target/primitive_conversions.go?s=2447:2481#L129)
``` go
func (i IntValues64) Data() []Data
```



## <a name="Iterable">type</a> [Iterable](/src/target/interable.go?s=67:120#L4)
``` go
type Iterable interface {
    Iterate() <-chan Element
}
```
Iterable is anything that can be ranged on










## <a name="StringValue">type</a> [StringValue](/src/target/primitive_conversions.go?s=4422:4445#L226)
``` go
type StringValue string
```
StringValue type alias for string










### <a name="StringValue.Float32">func</a> (StringValue) [Float32](/src/target/primitive_conversions.go?s=4862:4900#L247)
``` go
func (s StringValue) Float32() float32
```
Float32 casts and returns Float32 value




### <a name="StringValue.Float64">func</a> (StringValue) [Float64](/src/target/primitive_conversions.go?s=5011:5049#L253)
``` go
func (s StringValue) Float64() float64
```
Float64 casts and returns Float64 value




### <a name="StringValue.Int">func</a> (StringValue) [Int](/src/target/primitive_conversions.go?s=4482:4512#L229)
``` go
func (s StringValue) Int() int
```
Int casts and returns int value




### <a name="StringValue.Int32">func</a> (StringValue) [Int32](/src/target/primitive_conversions.go?s=4600:4634#L235)
``` go
func (s StringValue) Int32() int32
```
Int32 casts and returns Int32 value




### <a name="StringValue.Int64">func</a> (StringValue) [Int64](/src/target/primitive_conversions.go?s=4729:4763#L241)
``` go
func (s StringValue) Int64() int64
```
Int64 casts and returns Int64 value




### <a name="StringValue.String">func</a> (StringValue) [String](/src/target/primitive_conversions.go?s=5149:5185#L259)
``` go
func (s StringValue) String() string
```
String casts and returns string value




## <a name="StringValues">type</a> [StringValues](/src/target/primitive_conversions.go?s=5263:5289#L264)
``` go
type StringValues []string
```
StringValues type alias for a slice of StringValue










### <a name="StringValues.Data">func</a> (StringValues) [Data](/src/target/primitive_conversions.go?s=5291:5326#L266)
``` go
func (s StringValues) Data() []Data
```



## <a name="Type">type</a> [Type](/src/target/types.go?s=94:107#L4)
``` go
type Type int
```
Type is a type alias for an int used for the  "ENUM" definition below


``` go
const (
    UnknownType Type = iota
    IntType
    Int32Type
    Int64Type
    Float32Type
    Float64Type
    StringType

    IntSliceType
    Int32SliceType
    Int64SliceType
    Float32SliceType
    Float64SliceType
    StringSliceType
)
```
Collection supported types














- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
