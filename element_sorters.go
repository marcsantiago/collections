package collections

// Element represents (key, value) with backed by the Data interface
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

// ElementsByKeyStringAsc used to used low to high where value is of type string
type ElementsByValueStringAsc []Element

func (e ElementsByValueStringAsc) Len() int { return len(e) }
func (e ElementsByValueStringAsc) Less(i, j int) bool {
	return e[i].Value.String() < e[j].Value.String()
}
func (e ElementsByValueStringAsc) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

// ElementsByValueStringDesc used to used high to low where value is of type string
type ElementsByValueStringDesc []Element

func (e ElementsByValueStringDesc) Len() int { return len(e) }
func (e ElementsByValueStringDesc) Less(i, j int) bool {
	return e[i].Value.String() > e[j].Value.String()
}
func (e ElementsByValueStringDesc) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

// ElementsByValueIntAsc used to used low to high where value is of type string
type ElementsByValueIntAsc []Element

func (e ElementsByValueIntAsc) Len() int           { return len(e) }
func (e ElementsByValueIntAsc) Less(i, j int) bool { return e[i].Value.Int() < e[j].Value.Int() }
func (e ElementsByValueIntAsc) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

// ElementsByValueIntDesc used to used high to low where value is of type string
type ElementsByValueIntDesc []Element

func (e ElementsByValueIntDesc) Len() int           { return len(e) }
func (e ElementsByValueIntDesc) Less(i, j int) bool { return e[i].Value.Int() > e[j].Value.Int() }
func (e ElementsByValueIntDesc) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
