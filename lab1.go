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

	if(set.Length() > input.Length()) {
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

func main() {
	set1 := NewSet()
	set2 := NewSet()

	for i := 0; i < 10; i++ {
		set1.Add(i)
	}

	for i:= 5; i < 15; i++ {
		set2.Add(i)
	}

	fmt.Println(set1.Length(),set2.Length(),(set1.Intersect(set2)).Length(), *set1, *set2,*(set1.Intersect(set2)))
}