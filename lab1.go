package main

import (
	"fmt"
)

type storage map[interface{}]bool

func NewSet() *storage {
	st := make(storage)
	return &st
}

func (set *storage)Length() int {
	return len(*set)
}

func (set *storage)Add(item interface{}) bool {
	_, entry := (*set)[item]

	if(entry) {
		return false
	}

	(*set)[item] = true

	return true
}

func (set *storage)Has(item interface{}) bool {
	_, entry := (*set)[item]

	if(entry) {
		return true
	} else {
		return false
	}
}

func (set *storage)Intersect(input *storage) *storage {
	intersection := NewSet()

	if(set.Length() <= input.Length()) {
		for elem := range *set {
			if(input.Has(elem)) {
				intersection.Add(elem)
			}
		}
	} else {
		for elem := range *input {
			if(set.Has(elem)) {
				intersection.Add(elem)
			}
		}
	}

	return intersection
}

func (set *storage)Difference(input *storage) *storage {
	difference := NewSet()

	for elem := range *set {
		if(!input.Has(elem)) {
			difference.Add(elem)
		}
	}

	return difference
}

func (set *storage)Union(input *storage) *storage {
	union := NewSet()

	for elem := range *set {
		union.Add(elem)
	}

	for elem := range *input {
		if(!union.Has(elem)) {
			union.Add(elem)
		}
	}

	return union
}

func main() {
	set1 := NewSet()
	set2 := NewSet()

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
