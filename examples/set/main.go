package main

import (
	"fmt"

	"github.com/marcsantiago/collections"
	"github.com/marcsantiago/collections/set"
)

func main() {
	x := set.New()
	for i := 0; i < 100; i++ {
		x.Add(collections.IntValue(i))
	}

	y := set.New()
	for i := 50; i < 200; i++ {
		y.Add(collections.IntValue(i))
	}

	intersect := x.Intersection(y)
	fmt.Println("X", x)
	fmt.Println("Y", y)
	fmt.Println("intersection", intersect)

}
