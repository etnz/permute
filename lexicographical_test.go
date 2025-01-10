package permute

import (
	"fmt"
	"strings"
	"testing"
)

func TestLexNext(t *testing.T) {
	p := newPermutation(3)
	for i := 0; i < 6; i++ {
		LexNext(p)
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
