package namedtuple

import (
	"sort"
	"strings"

	"github.com/marcsantiago/collections"
)

// NamedTuple mirrors the immutable nature of the Python namedtuple
type NamedTuple struct {
	labels   map[string]int
	elements [][]collections.Data
}

// New creates a new NamedTuple with the labels synced with input order
func New(labels ...string) NamedTuple {
	nt := NamedTuple{
		labels: make(map[string]int),
	}
	for i, l := range labels {
		nt.labels[l] = i
	}
	return nt
}

// Add adds internally, order does matter should mirror the order the labels were inputted
func (n NamedTuple) Add(row []collections.Data) NamedTuple {
	n.elements = append(n.elements, row)
	return n
}

// Get retrieves a value from the internal data
func (n NamedTuple) Get(rowN int, label string) collections.Data {
	i, ok := n.labels[label]
	if !ok {
		return nil
	}
	return n.elements[rowN][i]
}

// Len returns the number of rows
func (n NamedTuple) Len() int {
	return len(n.elements)
}

// Reset empties the labels and internal rows of data, while keeping the internal slice size for potential reuse without
// slice reallocation
func (n *NamedTuple) Reset() {
	n.labels = make(map[string]int)
	n.elements = n.elements[:]
}

// String prints out the data in CSV format
func (n NamedTuple) String() string {
	var buf strings.Builder
	elements := make([]collections.Element, 0, len(n.labels))
	for k, v := range n.labels {
		elements = append(elements, collections.Element{
			Key:   collections.StringValue(k),
			Value: collections.IntValue(v),
		})
	}
	sort.Sort(collections.ElementsByValueIntAsc(elements))
	max := len(elements)
	for i, ele := range elements {
		buf.WriteString(ele.Key.String())
		if i+1 < max {
			buf.WriteRune(',')
		}
	}
	buf.WriteRune('\n')

	for _, row := range n.elements {
		for i, col := range row {
			buf.WriteString(col.String())
			if i+1 < max {
				buf.WriteRune(',')
			}
		}
		buf.WriteRune('\n')
	}
	return buf.String()
}
