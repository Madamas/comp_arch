package set

import "testing"

func benchmarkUnion(b *testing.B) {
	set1 := NewSet()
	set2 := NewSet()

	for i := 0; i < 100; i++ {
		set1.Add(i)
	}

	for i:= 75; i < 150; i++ {
		set2.Add(i)
	}

	for n := 0; n < b.N; n++ {
		set1.Union(set2)
	}
}

func benchmarkIntersect(b *testing.B) {
	set1 := NewSet()
	set2 := NewSet()

	for i := 0; i < 100; i++ {
		set1.Add(i)
	}

	for i:= 75; i < 150; i++ {
		set2.Add(i)
	}

	for n := 0; n < b.N; n++ {
		set1.Intersect(set2)
	}
}

func benchmarkDifference(b *testing.B) {
	set1 := NewSet()
	set2 := NewSet()

	for i := 0; i < 100; i++ {
		set1.Add(i)
	}

	for i:= 75; i < 150; i++ {
		set2.Add(i)
	}

	for n := 0; n < b.N; n++ {
		set1.Difference(set2)
	}
}

func benchmarkSet(b *testing.B) {
	set1 := NewSet()
	set2 := NewSet()

	for i := 0; i < 100; i++ {
		set1.Add(i)
	}

	for i:= 75; i < 150; i++ {
		set2.Add(i)
	}
	
	for n := 0; n < b.N; n++ {
		setA := set1.Difference(set2)
		setB := set2.Difference(set1)
		union := setA.Union(setB)
		instersect := set1.Intersect(union)
	}
}