package permute

import (
	"fmt"
	"strings"
	"testing"
)

func TestLexNext(t *testing.T) {
	p := newPermutation(3)
	for i := 0; i < 6; i++ {
		lexNext(p)
	}

	if p[0] != 0 || p[1] != 1 || p[2] != 2 {
		t.Errorf("invalid Next calculator: got %v, expecting 0 1 2 ", p)
	}
}

func ExampleLexicographical() {
	x := []string{"A", "B", "C", "D"}
	i := 0
	for v := range LexPermutations(x) {
		fmt.Printf("%2v:%v\n", i, strings.Join(v, ""))
		i++
	}

	//Output:
	//  0:ABCD
	//  1:ABDC
	//  2:ACBD
	//  3:ACDB
	//  4:ADBC
	//  5:ADCB
	//  6:BACD
	//  7:BADC
	//  8:BCAD
	//  9:BCDA
	// 10:BDAC
	// 11:BDCA
	// 12:CABD
	// 13:CADB
	// 14:CBAD
	// 15:CBDA
	// 16:CDAB
	// 17:CDBA
	// 18:DABC
	// 19:DACB
	// 20:DBAC
	// 21:DBCA
	// 22:DCAB
	// 23:DCBA
}
func ExampleSubsetLex() {
	i := 0
	x := []string{"1", "2", "3", "4", "5"}
	for v := range LexCombinations(3, x) {
		fmt.Printf("%v:%v\n", i, strings.Join(v, ""))
		i++
	}

	//Output:
	// 0:123
	// 1:124
	// 2:125
	// 3:134
	// 4:135
	// 5:145
	// 6:234
	// 7:235
	// 8:245
	// 9:345

}

func TestLexPermutations(t *testing.T) {
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
			for p := range LexPermutations(list) {
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

func TestLexCombinations(t *testing.T) {
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
				for p := range LexCombinations(n, list) {
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

// Binomial computes the binomial coefficient: from n choose k
//
// It's an efficient method that avoid useless overflow.
func Binomial(n, k int) int64 {
	var b int64 = 1
	if n-k < k {
		return Binomial(n, n-k)
	}
	for i := 1; i <= k; i++ {
		b *= int64((n - k + i))
		b /= int64(i)

	}
	return b
}
