package permute

import (
	"fmt"
	"strings"
	"testing"
)

func TestNext(t *testing.T) {
	p := New(3)
	for i := 0; i < 6; i++ {
		LexNext(p)
	}

	if p[0] != 0 || p[1] != 1 || p[2] != 2 {
		t.Errorf("invalid Next calculator: got %v, expecting 0 1 2 ", p)
	}
}

func ExampleLexicographical() {

	p := New(4)
	i := 0

	x := []string{"A", "B", "C", "D"}
	fmt.Printf("%2v:%v\n", i, strings.Join(x, ""))

	for LexNext(p) {
		x := []string{"A", "B", "C", "D"}
		Permute(p, x)
		i++
		fmt.Printf("%2v:%v\n", i, strings.Join(x, ""))
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

	p := New(3)
	i := 0
	x := []string{"1", "2", "3", "4", "5"}
	fmt.Printf("%v:%v\n", i, strings.Join(Subset(p, x), ""))
	for SubsetLexNext(p, len(x)) && i < 100 {

		fmt.Printf("%v:%v\n", i, strings.Join(Subset(p, x), ""))
		i++
	}
	fmt.Printf("end %v\n", strings.Join(Subset(p, x), ""))

	//Output:
	// 0:123
	// 0:124
	// 1:125
	// 2:134
	// 3:135
	// 4:145
	// 5:234
	// 6:235
	// 7:245
	// 8:345
	// end 123
}
