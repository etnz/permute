package permute

import (
	"fmt"
	"testing"
)

func BenchmarkParity(b *testing.B) {
	p := []int{3, 0, 4, 1, 5, 2}
	for n := 0; n < b.N; n++ {
		Parity(p)
	}

}

func BenchmarkParity2(b *testing.B) {
	p := []int{3, 0, 4, 1, 5, 2}
	for n := 0; n < b.N; n++ {
		Parity2(p)
	}

}

//deprecated by benchmarking, kept (in test) to not reproduce the error
func Parity2(p []int) int {
	return len(Transpositions(p)) % 2
}

func TestDelv(t *testing.T) {
	p := []int{6, 4, 1, 3, 5, 7}

	delv(&p, 3)
	if !Equals(p, []int{6, 4, 1, 5, 7}) {
		t.Errorf("invalid deletion %v expecting 6, 4, 1, 5, 7", p)
	}
	delv(&p, 7)
	if !Equals(p, []int{6, 4, 1, 5}) {
		t.Errorf("invalid deletion %v expecting 6, 4, 1, 5", p)
	}
	delv(&p, 6)
	if !Equals(p, []int{4, 1, 5}) {
		t.Errorf("invalid deletion %v expecting 4, 1, 5", p)
	}
	delv(&p, 1)
	if !Equals(p, []int{4, 5}) {
		t.Errorf("invalid deletion %v expecting 4, 5", p)
	}

}

func ExampleLehmerCode() {
	inv := LehmerCode([]int{3, 0, 4, 1, 5, 2})
	fmt.Printf("%v", inv)
	//Output: [3 0 2 0 1 0]

}

func ExampleNewLehmerCoded() {

	//the goal permutation
	p := []int{5, 2, 7, 0, 3, 8, 6, 1, 4}
	fmt.Printf("working with   %v\n", p)

	inv := LehmerCode(p)
	fmt.Printf("Lehmer Code is %v\n", inv)

	q := NewLehmerCoded(inv)
	if !Equals(p, q) {
		panic("oops")
	}

	//example from https://en.wikipedia.org/wiki/Factorial_number_system
	w := NewLehmerCoded([]int{4, 0, 4, 1, 0, 0, 0})
	fmt.Printf("working with   4041000!    (example from https://en.wikipedia.org/wiki/Factorial_number_system)\n")
	fmt.Printf("permutation is %v\n", w)
	//Output:
	// 	working with   [5 2 7 0 3 8 6 1 4]
	// Lehmer Code is [5 2 5 0 1 3 2 0 0]
	// working with   4041000!    (example from https://en.wikipedia.org/wiki/Factorial_number_system)
	// permutation is [4 0 6 2 1 3 5]
}
