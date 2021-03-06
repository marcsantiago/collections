

# ordereddict
`import "github.com/marcsantiago/collections/ordereddict"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [type OrderedDict](#OrderedDict)
  * [func New() OrderedDict](#New)
  * [func (o OrderedDict) Delete(key collections.Data)](#OrderedDict.Delete)
  * [func (o OrderedDict) Get(key collections.Data) (collections.Data, bool)](#OrderedDict.Get)
  * [func (o OrderedDict) Iterate() &lt;-chan collections.Element](#OrderedDict.Iterate)
  * [func (o OrderedDict) Reverse()](#OrderedDict.Reverse)
  * [func (o *OrderedDict) Set(key collections.Data, value collections.Data)](#OrderedDict.Set)
  * [func (o OrderedDict) String() string](#OrderedDict.String)


#### <a name="pkg-files">Package files</a>
[ordereddict.go](/src/github.com/marcsantiago/collections/ordereddict/ordereddict.go) 






## <a name="OrderedDict">type</a> [OrderedDict](/src/target/ordereddict.go?s=104:266#L11)
``` go
type OrderedDict struct {
    // contains filtered or unexported fields
}
```






### <a name="New">func</a> [New](/src/target/ordereddict.go?s=319:341#L19)
``` go
func New() OrderedDict
```
OrderedDict returns a map initialized structure





### <a name="OrderedDict.Delete">func</a> (OrderedDict) [Delete](/src/target/ordereddict.go?s=780:829#L34)
``` go
func (o OrderedDict) Delete(key collections.Data)
```
Delete removes the element from the internal map and the backing slice that records order




### <a name="OrderedDict.Get">func</a> (OrderedDict) [Get](/src/target/ordereddict.go?s=570:641#L28)
``` go
func (o OrderedDict) Get(key collections.Data) (collections.Data, bool)
```
Get retrieves a data value from the internal map if it exists




### <a name="OrderedDict.Iterate">func</a> (OrderedDict) [Iterate](/src/target/ordereddict.go?s=1611:1668#L66)
``` go
func (o OrderedDict) Iterate() <-chan collections.Element
```
Iterate creates a channel to create an iterator for he Go range statement




### <a name="OrderedDict.Reverse">func</a> (OrderedDict) [Reverse](/src/target/ordereddict.go?s=1908:1938#L78)
``` go
func (o OrderedDict) Reverse()
```
Reverse does an inplace reverse sort of the backing slice




### <a name="OrderedDict.Set">func</a> (\*OrderedDict) [Set](/src/target/ordereddict.go?s=1130:1201#L51)
``` go
func (o *OrderedDict) Set(key collections.Data, value collections.Data)
```
Set updates the internal map and backing slice that records order




### <a name="OrderedDict.String">func</a> (OrderedDict) [String](/src/target/ordereddict.go?s=2142:2178#L86)
``` go
func (o OrderedDict) String() string
```
String creates an OrderedDict representation of the internal map data








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
