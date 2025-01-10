package permute

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleSteinhausJohnsonTrotter() {
	i := 0

	A := []string{"A", "B", "C", "D"}

	for t, v := range SteinhausJohnsonTrotterPermutations(A) {
		fmt.Printf("%2d: (%d,%d) %v\n", i, t[0], t[1], v)
		i++
	}

	//Output:
	//  0: (0,0) [A B C D]
	//  1: (2,3) [A B D C]
	//  2: (1,2) [A D B C]
	//  3: (0,1) [D A B C]
	//  4: (2,3) [D A C B]
	//  5: (0,1) [A D C B]
	//  6: (1,2) [A C D B]
	//  7: (2,3) [A C B D]
	//  8: (0,1) [C A B D]
	//  9: (2,3) [C A D B]
	// 10: (1,2) [C D A B]
	// 11: (0,1) [D C A B]
	// 12: (2,3) [D C B A]
	// 13: (0,1) [C D B A]
	// 14: (1,2) [C B D A]
	// 15: (2,3) [C B A D]
	// 16: (0,1) [B C A D]
	// 17: (2,3) [B C D A]
	// 18: (1,2) [B D C A]
	// 19: (0,1) [D B C A]
	// 20: (2,3) [D B A C]
	// 21: (0,1) [B D A C]
	// 22: (1,2) [B A D C]
	// 23: (2,3) [B A C D]
}

func ExampleSteinhausJohnsonTrotterEven() {
	i := 0

	A := []string{"A", "B", "C", "D"}

	for t, v := range SteinhausJohnsonTrotterEvenPermutations(A) {
		fmt.Printf("%2d: (%d,%d) %v\n", i, t[0], t[1], v)
		i++
	}

	//Output:
	//  0: (0,0) [A B C D]
	//  1: (2,3) [A B D C]
	//  2: (1,2) [A D B C]
	//  3: (0,1) [D A B C]
	//  4: (2,3) [D A C B]
	//  5: (0,1) [A D C B]
	//  6: (1,2) [A C D B]
	//  7: (2,3) [A C B D]
	//  8: (0,1) [C A B D]
	//  9: (2,3) [C A D B]
	// 10: (1,2) [C D A B]
	// 11: (0,1) [D C A B]
	// 12: (2,3) [D C B A]
	// 13: (0,1) [C D B A]
	// 14: (1,2) [C B D A]
	// 15: (2,3) [C B A D]
	// 16: (0,1) [B C A D]
	// 17: (2,3) [B C D A]
	// 18: (1,2) [B D C A]
	// 19: (0,1) [D B C A]
	// 20: (2,3) [D B A C]
	// 21: (0,1) [B D A C]
	// 22: (1,2) [B A D C]
	// 23: (2,3) [B A C D]
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
			for _, p := range SteinhausJohnsonTrotterPermutations(list) {
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
			for _, p := range SteinhausJohnsonTrotterEvenPermutations(list) {
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
