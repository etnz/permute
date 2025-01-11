package permute

import "fmt"

func ExamplePermute() {
	// For a given set 'x'.
	x := []string{"a", "b", "c"}
	// The permutation 'p'
	p := []int{2, 1, 0}
	// transform 'x' into: `{x[2], x[1], x[0]}`
	Permute(p, x)
	fmt.Printf("%#v", x)
	//Output: []string{"c", "b", "a"}
}

func ExampleSubset() {
	// For a given set 'x'.
	x := []string{"a", "b", "c"}
	// The combination 'c'
	c := []int{0, 2}
	// transform 'x' into: `{x[2], x[0]}`
	x = Subset(c, x)
	fmt.Printf("%#v", x)
	//Output: []string{"a", "c"}
}

func ExamplePermutations() {
	// For a given set 'x'.
	x := []string{"a", "b", "c"}

	// One can simply loop over all permutations.
	for _, p := range Permutations(x) {
		fmt.Println(p)
	}
	//Output:
	// [a b c]
	// [b a c]
	// [c a b]
	// [a c b]
	// [b c a]
	// [c b a]
}

func ExampleCombinations() {
	// For a given set 'x'.
	x := []string{"a", "b", "c"}

	// One can simply loop over all 2-combinations.
	for v := range Combinations(2, x) {
		fmt.Println(v)
	}
	//Output:
	// [a b]
	// [b c]
	// [a c]
}
