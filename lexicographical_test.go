package permute

import (
	"fmt"
	"strings"
	"testing"
)

func TestNext(t *testing.T) {
	p := New(3)
	for i := 0; i < 6; i++ {
		Lexicographical(p)
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

	for Lexicographical(p) {
		x := []string{"A", "B", "C", "D"}
		Strings(p, x)
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
