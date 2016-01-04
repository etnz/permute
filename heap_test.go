package permute

import "fmt"

func ExampleHeap_Next() {
	A := []string{"A", "B", "C", "D"}

	h := NewHeap(len(A))
	i := 0
	var s [2]int
	for h.Next(&s) {
		swaps(s, A)
		fmt.Printf("%2d: (%v, %v)  %v\n", i, s[0], s[1], A)
		i++
	}

	//Output:
	//  0: (0, 1)  [B A C D]
	//  1: (0, 2)  [C A B D]
	//  2: (0, 1)  [A C B D]
	//  3: (0, 2)  [B C A D]
	//  4: (0, 1)  [C B A D]
	//  5: (0, 3)  [D B A C]
	//  6: (0, 1)  [B D A C]
	//  7: (0, 2)  [A D B C]
	//  8: (0, 1)  [D A B C]
	//  9: (0, 2)  [B A D C]
	// 10: (0, 1)  [A B D C]
	// 11: (1, 3)  [A C D B]
	// 12: (0, 1)  [C A D B]
	// 13: (0, 2)  [D A C B]
	// 14: (0, 1)  [A D C B]
	// 15: (0, 2)  [C D A B]
	// 16: (0, 1)  [D C A B]
	// 17: (2, 3)  [D C B A]
	// 18: (0, 1)  [C D B A]
	// 19: (0, 2)  [B D C A]
	// 20: (0, 1)  [D B C A]
	// 21: (0, 2)  [C B D A]
	// 22: (0, 1)  [B C D A]
}
