package permute

import (
	"fmt"
	"slices"
	"testing"
)

func TestPermute(t *testing.T) {
	val := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	x := []string{"c", "d", "e", "g", "f", "a", "h", "b"}
	p := []int{2, 3, 4, 6, 5, 0, 7, 1}

	Permute(p, val)
	if !slices.Equal(x, val) {
		t.Errorf("Invalid Apply: got %v, expecting %v", val, x)
	}

}

func ExamplePermutations() {
	A := []string{"A", "B", "C", "D"}

	// Loop over all permutations of A.
	for _, v := range Permutations(A) {
		fmt.Println(v)
	}

	//Output:
	// [A B C D]
	// [B A C D]
	// [C A B D]
	// [A C B D]
	// [B C A D]
	// [C B A D]
	// [D B A C]
	// [B D A C]
	// [A D B C]
	// [D A B C]
	// [B A D C]
	// [A B D C]
	// [A C D B]
	// [C A D B]
	// [D A C B]
	// [A D C B]
	// [C D A B]
	// [D C A B]
	// [D C B A]
	// [C D B A]
	// [B D C A]
	// [D B C A]
	// [C B D A]
	// [B C D A]
}

func ExamplePermutations_WithTranspositions() {
	A := []string{"A", "B", "C", "D"}

	// Loop over all permutations of A.
	for t, v := range Permutations(A) {
		fmt.Println(t, v)
	}

	//Output:
	// [0 0] [A B C D]
	// [0 1] [B A C D]
	// [0 2] [C A B D]
	// [0 1] [A C B D]
	// [0 2] [B C A D]
	// [0 1] [C B A D]
	// [0 3] [D B A C]
	// [0 1] [B D A C]
	// [0 2] [A D B C]
	// [0 1] [D A B C]
	// [0 2] [B A D C]
	// [0 1] [A B D C]
	// [1 3] [A C D B]
	// [0 1] [C A D B]
	// [0 2] [D A C B]
	// [0 1] [A D C B]
	// [0 2] [C D A B]
	// [0 1] [D C A B]
	// [2 3] [D C B A]
	// [0 1] [C D B A]
	// [0 2] [B D C A]
	// [0 1] [D B C A]
	// [0 2] [C B D A]
	// [0 1] [B C D A]
}

func ExampleCombinations() {
	A := []string{"A", "B", "C", "D"}

	// Loop over all 2-Combinations of A.
	for v := range Combinations(2, A) {
		fmt.Println(v)
	}

	//Output:
	// [A B]
	// [B C]
	// [A C]
	// [C D]
	// [B D]
	// [A D]
}

func TestAsTranspositions(t *testing.T) {
	//let p be the permutation :
	p := []int{2, 3, 4, 6, 5, 0, 7, 1}

	x := newPermutation(len(p))

	swaps := Decompose(p)
	for _, perm := range swaps {
		transpose(perm, x)
	}
	if !slices.Equal(p, x) {
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
		transpose(perm, current) //inplace permute
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
		perm := newTransposition(k, sk)
		swaps = append(swaps, perm)
		transpose(perm, current) //inplace permute
	}
	//reverse the array
	ns := len(swaps) - 1
	for i := 0; i < (ns+1)/2; i++ {
		swaps[i], swaps[ns-i] = swaps[ns-i], swaps[i]
	}
	return
}
