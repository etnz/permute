package permute

import (
	"github.com/etnz/stringset"

	"fmt"
)

func ExampleNewMinimalChangeGen() {
	n, k := 5, 3
	gen := NewMinimalChangeGen(n, k)
	var sw [2]int
	x := []string{"A", "B", "C", "D", "E"}
	y := map[string]interface{}{"A": nil, "B": nil, "C": nil}

	fmt.Printf("%v\n", stringset.Sort(y))
	for gen.Next(&sw) {
		//delete sw0 and insert sw1
		delete(y, x[sw[0]])
		y[x[sw[1]]] = nil
		fmt.Printf("%v\n", stringset.Sort(y))
	}
	//Output:
	// [A B C]
	// [A C D]
	// [B C D]
	// [A B D]
	// [A D E]
	// [B D E]
	// [C D E]
	// [A C E]
	// [B C E]
	// [A B E]

}

func ExampleRevolvingDoorNext() {
	printRevolving(2, 1)
	printRevolving(3, 1)
	printRevolving(3, 2)
	printRevolving(4, 1)
	printRevolving(4, 2)
	printRevolving(4, 3)
	printRevolving(5, 1)
	printRevolving(5, 2)
	printRevolving(5, 3)
	printRevolving(5, 4)

	//Output:
	// Combinations(2,1)=[0],[1],
	// Combinations(3,1)=[0],[1],[2],
	// Combinations(3,2)=[0 1],[1 2],[0 2],
	// Combinations(4,1)=[0],[1],[2],[3],
	// Combinations(4,2)=[0 1],[1 2],[0 2],[2 3],[1 3],[0 3],
	// Combinations(4,3)=[0 1 2],[0 2 3],[1 2 3],[0 1 3],
	// Combinations(5,1)=[0],[1],[2],[3],[4],
	// Combinations(5,2)=[0 1],[1 2],[0 2],[2 3],[1 3],[0 3],[3 4],[2 4],[1 4],[0 4],
	// Combinations(5,3)=[0 1 2],[0 2 3],[1 2 3],[0 1 3],[0 3 4],[1 3 4],[2 3 4],[0 2 4],[1 2 4],[0 1 4],
	// Combinations(5,4)=[0 1 2 3],[0 1 3 4],[1 2 3 4],[0 2 3 4],[0 1 2 4],

}

func printRevolving(n, k int) {
	p := New(k) // the first permutation
	var sw [2]int
	fmt.Printf("Combinations(%v,%v)=", n, k)
	for i := int64(0); i < Binomial(n, k); i++ {
		fmt.Printf("%v,", p)
		SubsetRevolvingDoorNext(p, n, &sw)
	}
	fmt.Println()
}
