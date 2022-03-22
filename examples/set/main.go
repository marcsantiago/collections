package main

import (
	"fmt"

	"github.com/marcsantiago/collections/set"
)

func main() {
	x := set.New[int]()
	for i := 0; i < 100; i++ {
		x.Add(i)
	}

	y := set.New[int]()
	for i := 50; i < 200; i++ {
		y.Add(i)
	}

	intersect := x.Intersection(y)
	fmt.Println("X", x)
	fmt.Println("Y", y)
	fmt.Println("intersection", intersect)

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	alphaSet := set.New[rune]()
	for ch := range alphabet {
		alphaSet.Add(rune(ch))
	}

	word := "hello world"
	wordSet := set.New[rune]()
	for ch := range word {
		alphaSet.Add(rune(ch))
	}

	if wordSet.IsSubset(alphaSet) {
		fmt.Printf("\"%s\" seems like english", word)
	}
}
