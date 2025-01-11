package permute

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleSteinhausJohnsonTrotter() {
	A := []string{"A", "B", "C"}
	for _, v := range SJTPermutations(A) {
		fmt.Println(v)
	}

	//Output:
	// [A B C]
	// [A C B]
	// [C A B]
	// [C B A]
	// [B C A]
	// [B A C]
}

func ExampleSteinhausJohnsonTrotterEven() {
	A := []string{"A", "B", "C"}
	for _, v := range SJTEPermutations(A) {
		fmt.Println(v)
	}

	//Output:
	// [A B C]
	// [A C B]
	// [C A B]
	// [C B A]
	// [B C A]
	// [B A C]
}

func TestSteinhausJohnsonTrotterPermutations(t *testing.T) {
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
			for _, p := range SJTPermutations(list) {
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

func TestSteinhausJohnsonTrotterEvenPermutations(t *testing.T) {
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
			for _, p := range SJTEPermutations(list) {
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
