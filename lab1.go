package main

import (
	"fmt"
	"set"
)

func main() {
	set1 := set.NewSet()
	set2 := set.NewSet()

	for i := 0; i < 5; i++ {
		set1.Add(i)
	}

	for i:= 3; i < 7; i++ {
		set2.Add(i)
	}

	fmt.Printf("Set1: %v\nLength: %v\n\n", *set1, set1.Length())
	fmt.Printf("Set2: %v\nLength: %v\n\n", *set2, set2.Length())

	intersectResult := set1.Intersect(set2)
	fmt.Printf("Intersection: %v\nLength: %v\n\n", *intersectResult, intersectResult.Length())

	differenceResult := set1.Difference(set2)
	fmt.Printf("Difference: %v\nLength: %v\n\n", *differenceResult, differenceResult.Length())

	unionResult := set1.Union(set2)
	fmt.Printf("Union: %v\nLength: %v\n\n", *unionResult, unionResult.Length())
}
