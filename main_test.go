package permute

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"testing"
)

func TestInv(t *testing.T) {

	val := []string{"a", "b", "c"}

	p := []int{1, 2, 0}
	q := Inv(p)
	Strings(p, val)
	Strings(q, val) // q beeing the inverse of p, I should go back to val
	if !reflect.DeepEqual(val, []string{"a", "b", "c"}) {
		t.Fatalf("invalid permutation, got %v expecting abc", strings.Join(val, ""))
	}
}

func TestInts(t *testing.T) {

	p := []int{1, 2, 0}
	val := New(len(p))

	Ints(p, val)
	if !Equals(p, val) {
		t.Errorf("invalid Apply: got %v, expecting %v", val, p)
	}

}

func TestStrings(t *testing.T) {

	val := []string{"a", "b", "c"}
	p := []int{1, 2, 0}
	Strings(p, val)
	if val[0] != "b" || val[1] != "c" || val[2] != "a" {
		t.Errorf("invalid ApplyString: got %v, expecting b c a ", val)
	}
}
func ExampleStrings() {
	x := []string{"a", "b", "c"}
	Strings([]int{2, 1, 0}, x)
	fmt.Println(x)
	//Output: [c b a]
}

func ExampleSubStrings() {
	x := SubStrings([]int{0, 2}, []string{"a", "b", "c"})
	fmt.Println(x)
	//Output: [a c]

}

func TestApply(t *testing.T) {
	val := sort.StringSlice{"a", "b", "c", "d", "e", "f", "g", "h"}
	x := sort.StringSlice{"c", "d", "e", "g", "f", "a", "h", "b"}
	p := []int{2, 3, 4, 6, 5, 0, 7, 1}

	Apply(p, val)
	if !reflect.DeepEqual(x, val) {
		t.Errorf("Invalid Apply Interface: got %v, expecting %v", val, x)
	}

}

func BenchmarkInts(b *testing.B) {
	p := []int{2, 3, 4, 6, 5, 0, 7, 1}
	nn := New(len(p))
	for n := 0; n < b.N; n++ {
		Ints(p, nn)
	}
}

func BenchmarkInts2(b *testing.B) {
	p := []int{2, 3, 4, 6, 5, 0, 7, 1}
	nn := New(len(p))
	for n := 0; n < b.N; n++ {
		ints(p, nn)
	}
}

//alternative impl to test speed
func ints(p []int, val []int) {
	Apply(p, sort.IntSlice(val))
}

func swaps(s [2]int, p []string) {
	k, pk := s[0], s[1]
	p[k], p[pk] = p[pk], p[k]
}
