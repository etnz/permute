package permute

import (
	"fmt"
	"slices"
	"testing"
)

// TestPermute assess on a simple case that the function Permute is correct.
// Corner case are tested by iterators.
func TestPermute(t *testing.T) {
	val := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	x := []string{"c", "d", "e", "g", "f", "a", "h", "b"}
	p := []int{2, 3, 4, 6, 5, 0, 7, 1}

	Permute(p, val)
	if !slices.Equal(x, val) {
		t.Errorf("Invalid Apply: got %v, expecting %v", val, x)
	}

}


func ExamplePermutations_WithTranspositions() {
	A := []string{"A", "B", "C"}

	// Loop over all permutations of A.
	for t := range Permutations(A) {
		// 't' is the transposition to go from previous value to the current one.
		Transpose(t, A)
		fmt.Println(A)
	}

	//Output:
	// [A B C]
	// [B A C]
	// [C A B]
	// [A C B]
	// [B C A]
	// [C B A]
}

// TestDecompose assess on a simple example that the
// result is correct.
func TestDecompose(t *testing.T) {
	//let p be the permutation :
	p := []int{2, 3, 4, 6, 5, 0, 7, 1}

	x := newPermutation(len(p))

	transpositions := Decompose(p)
	for _, t := range transpositions {
		Transpose(t, x)
	}
	if !slices.Equal(p, x) {
		t.Errorf("Invalid Transposition decomposition, got %v expecting %v", x, p)
	}
}

func BenchmarkDecompose(b *testing.B) {
	p := []int{2, 3, 4, 6, 5, 0, 7, 1, 9, 8}

	for n := 0; n < b.N; n++ {
		Decompose(p)
	}
}
