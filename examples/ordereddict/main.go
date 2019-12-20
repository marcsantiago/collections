package main

import (
	"fmt"

	"github.com/marcsantiago/collections"
	"github.com/marcsantiago/collections/ordereddict"
)

func main() {
	od := ordereddict.New()
	for i := 1; i <= 30; i++ {
		key, value := i, i*2
		od.Set(collections.IntValue(key), collections.IntValue(value))
	}

	for elem := range od.Iterate() {
		fmt.Println("before changes", "key", elem.Key, "value", elem.Value)
	}

	od.Delete(collections.IntValue(10))
	od.Set(collections.IntValue(999), collections.IntValue(123456))
	od.Set(collections.IntValue(10), collections.IntValue(20))

	for elem := range od.Iterate() {
		fmt.Println("after changes", "key", elem.Key, "value", elem.Value)
	}
	fmt.Println(od)
	od.Reverse()
	fmt.Println(od)
}
