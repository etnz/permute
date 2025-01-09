//go:build go1.23

package permute

import "fmt"

// func TestPermutations(t *testing.T) {}

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

func ExamplePermutations2() {
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
	fmt.Printf("There are %v combinations.\n", Binomial(len(A), 2))
	for v := range Combinations(2, A) {
		fmt.Println(v)
	}

	//Output:
	// There are 6 combinations.
	// [A B]
	// [B C]
	// [A C]
	// [C D]
	// [B D]
	// [A D]
}
