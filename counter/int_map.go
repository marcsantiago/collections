package counter

import (
	"encoding/json"

	"github.com/marcsantiago/collections"
)

type IntMap map[int]int

func (i IntMap) Get(key collections.Data) (collections.Data, bool) {
	val, ok := i[key.Int()]
	return collections.IntValue(val), ok
}

func (i IntMap) Set(value collections.Data) {
	i[value.Int()]++
}

func (i IntMap) Delete(value collections.Data) {
	delete(i, value.Int())
}

func (i IntMap) Items() []collections.Element {
	items := make([]collections.Element, 0, len(i))
	for key, value := range i {
		items = append(items, collections.Element{
			Key:   collections.IntValue(key),
			Value: collections.IntValue(value),
		})
	}
	return items
}

func (i IntMap) Iterate() <-chan collections.Element {
	ch := make(chan collections.Element)
	go func() {
		for key, value := range i {
			ch <- collections.Element{Key: collections.IntValue(key), Value: collections.IntValue(value)}
		}
		close(ch)
	}()
	return ch
}

func (i IntMap) String() string {
	b, _ := json.Marshal(i)
	return string(b)
}
