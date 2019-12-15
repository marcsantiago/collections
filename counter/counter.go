package counter

import "github.com/marcsantiago/collections"

type IntSliceCounter struct {
	Map map[int]int
}

func NewIntSliceCounter(arr []int) IntSliceCounter {
	obj := IntSliceCounter{Map: make(map[int]int)}
	for _, a := range arr {
		obj.Map[a]++
	}
	return obj
}

func (o IntSliceCounter) Iterate() <-chan collections.Element {
	ch := make(chan collections.Element)
	go func() {
		for key, value := range o.Map {
			ch <- collections.Element{Key: collections.IntValue(key), Value: collections.IntValue(value)}
		}
		close(ch)
	}()
	return ch
}
