

# chain
`import "github.com/marcsantiago/collections/chainmap"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
ChainMap class is provided for quickly linking a number of mappings so they can be treated as a single unit. It is often much faster than creating a new dictionary and running multiple update() calls. A go parody of the Python implementation.




## <a name="pkg-index">Index</a>
* [type Maps](#Maps)
  * [func New(maps ...collections.Map) Maps](#New)
  * [func (m Maps) Delete(key collections.Data)](#Maps.Delete)
  * [func (m Maps) Get(key collections.Data) (collections.Data, bool)](#Maps.Get)
  * [func (m Maps) Items() []collections.Element](#Maps.Items)
  * [func (m Maps) Iterate() &lt;-chan collections.Element](#Maps.Iterate)
  * [func (m Maps) NewChild(mm collections.Map) Maps](#Maps.NewChild)
  * [func (m Maps) Parents() Maps](#Maps.Parents)
  * [func (m Maps) Set(key collections.Data, value collections.Data)](#Maps.Set)
  * [func (m Maps) String() string](#Maps.String)


#### <a name="pkg-files">Package files</a>
[chainmap.go](/src/github.com/marcsantiago/collections/chainmap/chainmap.go) [doc.go](/src/github.com/marcsantiago/collections/chainmap/doc.go) 






## <a name="Maps">type</a> [Maps](/src/target/chainmap.go?s=84:111#L9)
``` go
type Maps []collections.Map
```






### <a name="New">func</a> [New](/src/target/chainmap.go?s=420:458#L15)
``` go
func New(maps ...collections.Map) Maps
```
New returns a new chain map
if no maps are passed i new chain map with an initialized map is return
else all the maps are merged into indexes
Lookups search the underlying mappings successively until a key is found. In contrast, writes, updates, and deletions only operate on the first mapping.





### <a name="Maps.Delete">func</a> (Maps) [Delete](/src/target/chainmap.go?s=693:735#L28)
``` go
func (m Maps) Delete(key collections.Data)
```
Delete removes the element from the first internal map




### <a name="Maps.Get">func</a> (Maps) [Get](/src/target/chainmap.go?s=849:913#L33)
``` go
func (m Maps) Get(key collections.Data) (collections.Data, bool)
```
Get retrieves a data value from the first internal map that contains they targeted key




### <a name="Maps.Items">func</a> (Maps) [Items](/src/target/chainmap.go?s=2002:2045#L68)
``` go
func (m Maps) Items() []collections.Element
```
Items returns the internal map as a set of elements




### <a name="Maps.Iterate">func</a> (Maps) [Iterate](/src/target/chainmap.go?s=2415:2465#L87)
``` go
func (m Maps) Iterate() <-chan collections.Element
```
Iterate creates a channel to create an iterator for he Go range statement




### <a name="Maps.NewChild">func</a> (Maps) [NewChild](/src/target/chainmap.go?s=1255:1302#L44)
``` go
func (m Maps) NewChild(mm collections.Map) Maps
```
NewChild returns a new ChainMap containing a new map followed by all of the maps in the current instance. If m is specified, it becomes the new map at the front of the list of mappings; if not specified, an empty dict is used




### <a name="Maps.Parents">func</a> (Maps) [Parents](/src/target/chainmap.go?s=1607:1635#L52)
``` go
func (m Maps) Parents() Maps
```
Parents returning a new ChainMap containing all of the maps in the current instance except the first one. This is useful for skipping the first map in the search.




### <a name="Maps.Set">func</a> (Maps) [Set](/src/target/chainmap.go?s=1713:1776#L57)
``` go
func (m Maps) Set(key collections.Data, value collections.Data)
```
Set adds a key value pairing to the first internal map




### <a name="Maps.String">func</a> (Maps) [String](/src/target/chainmap.go?s=1869:1898#L62)
``` go
func (m Maps) String() string
```
String returns the JSON string representation of the map data








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)