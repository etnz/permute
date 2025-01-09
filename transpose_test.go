//go:build go1.18

package permute

import (
	"testing"
)

func TestAsTranspositions(t *testing.T) {
	//let p be the permutation :
	p := []int{2, 3, 4, 6, 5, 0, 7, 1}

	x := New(len(p))

	swaps := Decompose(p)
	for _, perm := range swaps {
		swap(perm, x)
	}
	if !Equals(p, x) {
		t.Errorf("Invalid Transposition decomposition, got %v expecting %v", x, p)
	}

}

func BenchmarkAsTransposition(b *testing.B) {
	p := []int{2, 3, 4, 6, 5, 0, 7, 1, 9, 8}

	for n := 0; n < b.N; n++ {
		Decompose(p)
	}
}

func BenchmarkAsTransposition2(b *testing.B) {
	p := []int{2, 3, 4, 6, 5, 0, 7, 1, 9, 8}

	for n := 0; n < b.N; n++ {
		Transpositions2(p)
	}
}

func BenchmarkAsTransposition3(b *testing.B) {
	p := []int{2, 3, 4, 6, 5, 0, 7, 1, 9, 8}
	for n := 0; n < b.N; n++ {
		Transpositions3(p)
	}
}

// alternative implementation: using prepend, but disqualified by benchmarks
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
		swap(perm, current) //inplace permute
	}
	return
}

// alternative implementation: using prepend, but disqualified by benchmarks
// we keep it to prevent anyone from "reinventing" the wheel
func Transpositions3(p []int) (swaps []T) {

	//we don't know the exact number of swaps
	swaps = make([]T, 0, len(p))

	// iteratively convert 'current' towards the identity using atomic Transpositions,
	// record them, and return them in reverse order.
	current := make([]int, len(p))
	copy(current, p)
	for {
		// for k := smallestNonFixedIndex(current); k >= 0; k = smallestNonFixedIndex(current) {
		// Get the first index that is different from the identity permutation.
		k := smallestNonFixedIndex(current)
		if k < 0 {
			// There are none, so current has reached the identity.
			break
		}
		sk := indexof(k, current)
		perm := NewTransposition(k, sk)
		swaps = append(swaps, perm)
		swap(perm, current) //inplace permute
	}
	//reverse the array
	ns := len(swaps) - 1
	for i := 0; i < (ns+1)/2; i++ {
		swaps[i], swaps[ns-i] = swaps[ns-i], swaps[i]
	}
	return
}
