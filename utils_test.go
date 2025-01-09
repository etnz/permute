package permute

import (
	"fmt"
	"slices"
	"strings"
	"testing"
)

func TestInv(t *testing.T) {
	val := []string{"a", "b", "c"}
	want := []string{"a", "b", "c"}

	p := []int{1, 2, 0}
	q := Inv(p)
	Permute(p, val)
	Permute(q, val) // q beeing the inverse of p, I should go back to val
	if !slices.Equal(val, want) {
		t.Fatalf("invalid permutation, got %v expecting abc", strings.Join(val, ""))
	}
}

func TestPermute(t *testing.T) {
	val := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	x := []string{"c", "d", "e", "g", "f", "a", "h", "b"}
	p := []int{2, 3, 4, 6, 5, 0, 7, 1}

	Permute(p, val)
	if !slices.Equal(x, val) {
		t.Errorf("Invalid Apply: got %v, expecting %v", val, x)
	}

}

func ExampleStrings() {
	x := []string{"a", "b", "c"}
	Permute([]int{2, 1, 0}, x)
	fmt.Println(x)
	//Output: [c b a]
}

func ExampleSubStrings() {
	x := Subset([]int{0, 2}, []string{"a", "b", "c"})
	fmt.Println(x)
	//Output: [a c]

}

func swaps(s [2]int, p []string) {
	k, pk := s[0], s[1]
	p[k], p[pk] = p[pk], p[k]
}
