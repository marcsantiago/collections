package collections

type Type int

const (
	UnknownType Type = iota
	IntSliceType
	//Need to implement more types in the counter package
	//Int32SliceType
	//Int64SliceType
	//Float32SliceType
	//Float64SliceType
	//BoolSliceType
	//StringSliceType
)

type Element struct {
	Key   Data
	Value Data
}

// ElementsByKeyStringAsc used to used low to high where key is of type string
type ElementsByKeyStringAsc []Element

func (e ElementsByKeyStringAsc) Len() int           { return len(e) }
func (e ElementsByKeyStringAsc) Less(i, j int) bool { return e[i].Key.String() < e[j].Key.String() }
func (e ElementsByKeyStringAsc) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

// ElementsByKeyStringDesc used to used high to low where key is of type string
type ElementsByKeyStringDesc []Element

func (e ElementsByKeyStringDesc) Len() int           { return len(e) }
func (e ElementsByKeyStringDesc) Less(i, j int) bool { return e[i].Key.String() > e[j].Key.String() }
func (e ElementsByKeyStringDesc) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

// ElementsByKeyIntAsc used to used low to high where key is of type string
type ElementsByKeyIntAsc []Element

func (e ElementsByKeyIntAsc) Len() int           { return len(e) }
func (e ElementsByKeyIntAsc) Less(i, j int) bool { return e[i].Key.Int() < e[j].Key.Int() }
func (e ElementsByKeyIntAsc) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

// ElementsByKeyIntDesc used to used high to low where key is of type string
type ElementsByKeyIntDesc []Element

func (e ElementsByKeyIntDesc) Len() int           { return len(e) }
func (e ElementsByKeyIntDesc) Less(i, j int) bool { return e[i].Key.Int() > e[j].Key.Int() }
func (e ElementsByKeyIntDesc) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
