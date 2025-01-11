package permute

import (
	"fmt"
	"strings"
	"testing"
)

func TestRevolvingDoorCombinations(t *testing.T) {
	data := []string{"", "A", "AB", "ABC", "ABCD"}

	for _, td := range data {
		for n := 0; n <= len(td); n++ {
			t.Run(fmt.Sprintf("%v_%v", td, n), func(t *testing.T) {
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
				//Prepare a set of all combinations.
				all := make(map[string]struct{})
				count := 0
				for p := range RevCombinations(n, list) {
					all[strings.Join(p, "")] = struct{}{}
					for _, i := range p {
						if _, in := items[i]; !in {
							t.Errorf("Invalid proposed combination %q, some items are not from the initial set.", p)
						}
					}
					if len(p) != n {
						t.Errorf("Invalid proposed combination %q: wrong number of items %v want %v.", p, len(p), n)
					}
					if !uniques(p) {
						t.Errorf("Invalid proposed combination %q: not all items are unique.", p)
					}
					count++
				}

				expectedCount := int(Binomial(len(list), n))
				if count != expectedCount {
					t.Errorf("Invalid iterator: returned %v items, want %v", count, expectedCount)
				}
				if expectedCount != len(all) {
					t.Errorf("Invalid iterator: returned duplicated permutations. got %v unique permutations, want %v", len(all), expectedCount)
				}

			})
		}
	}
}
