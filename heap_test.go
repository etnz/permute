package permute

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleHeapPermutations() {
	A := []string{"A", "B", "C", "D"}

	for t, v := range HeapPermutations(A) {
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

func TestHeapPermutations(t *testing.T) {
	data := []string{"", "A", "AB", "ABC", "ABCD"}

	for _, td := range data {
		t.Run(td, func(t *testing.T) {
			// Convert test data into a list of single char.
			list := strings.Split(td, "")
			// Prepare a set of all items in the above list (they must be unique).
			items := make(map[string]struct{})
			for _, i := range list {
				items[i] = struct{}{}
			}
			if len(items) != len(list) {
				t.Fatalf("invalid test data: Letters must be unique, got %v unique items vs %v items in the source", len(items), len(list))
			}
			//Prepare a set of all permutations.
			all := make(map[string]struct{})
			count := 0
			for _, p := range HeapPermutations(list) {
				all[strings.Join(p, "")] = struct{}{}
				for _, i := range p {
					if _, in := items[i]; !in {
						t.Errorf("Invalid proposed permutation %q, some items are not from the initial set.", p)
					}
				}
				if len(p) != len(list) {
					t.Errorf("Invalid proposed permutation %q: wrong number of items %v want %v.", p, len(p), len(list))
				}
				if !uniques(p) {
					t.Errorf("Invalid proposed permutation %q: not all items are unique.", p)
				}
				count++
			}

			expectedCount := fact(len(list))
			if count != expectedCount {
				t.Errorf("Invalid iterator: returned %v items, want %v", count, expectedCount)
			}
			if expectedCount != len(all) {
				t.Errorf("Invalid iterator: returned duplicated permutations. got %v unique permutations, want %v", len(all), expectedCount)
			}

		})
	}
}

func fact(n int) int {
	if n > 0 {
		return n * fact(n-1)
	}
	return 1
}

func uniques[E comparable](list []E) bool {
	all := make(map[E]struct{})
	for _, v := range list {
		all[v] = struct{}{}
	}
	return len(all) == len(list)

}
