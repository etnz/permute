package permute

import "testing"

func TestAsTranspositions(t *testing.T) {
	//let p be the permutation :
	p := []int{2, 3, 4, 6, 5, 0, 7, 1}

	x := New(len(p))

	swaps := Transpositions(p)
	for _, perm := range swaps {
		SwapInts(perm, x)
	}
	if !Equals(p, x) {
		t.Errorf("Invalid Transposition decomposition, got %v expecting %v", x, p)
	}

}

func BenchmarkAsTransposition(b *testing.B) {
	p := []int{2, 3, 4, 6, 5, 0, 7, 1}

	for n := 0; n < b.N; n++ {
		Transpositions(p)
	}
}

func BenchmarkAsTransposition2(b *testing.B) {
	p := []int{2, 3, 4, 6, 5, 0, 7, 1}

	for n := 0; n < b.N; n++ {
		Transpositions2(p)
	}
}

//alternative implementation: using prepend, but disqualified by benchmarks
// we keep it to prevent anyone from "reinventing" the wheel
func Transpositions2(p []int) (swaps [][2]int) {

	//we don't know the exact number of swaps
	swaps = make([][2]int, 0, len(p))

	current := make([]int, len(p))
	copy(current, p)

	for k := smallestNonFixedIndex(current); k >= 0; k = smallestNonFixedIndex(current) {
		sk := indexof(k, current)
		perm := [2]int{k, sk}
		swaps = append([][2]int{perm}, swaps...)
		SwapInts(perm, current) //inplace permute
	}
	return
}
