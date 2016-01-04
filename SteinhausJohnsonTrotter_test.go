package permute

import "fmt"

func ExampleSteinhausJohnsonTrotter() {
	i := 0

	A := []string{"A", "B", "C", "D"}
	fmt.Printf("%2d:       %v\n", i, A)

	p := New(len(A))
	var sw [2]int
	for SteinhausJohnsonTrotter(p, &sw) {
		i++
		SwapStrings(sw, A)
		fmt.Printf("%2d: (%d,%d) %v\n", i, sw[0], sw[1], A)
	}

	//Output:
	//  0:       [A B C D]
	//  1: (2,3) [A B D C]
	//  2: (1,2) [A D B C]
	//  3: (0,1) [D A B C]
	//  4: (2,3) [D A C B]
	//  5: (0,1) [A D C B]
	//  6: (1,2) [A C D B]
	//  7: (2,3) [A C B D]
	//  8: (0,1) [C A B D]
	//  9: (2,3) [C A D B]
	// 10: (1,2) [C D A B]
	// 11: (0,1) [D C A B]
	// 12: (2,3) [D C B A]
	// 13: (0,1) [C D B A]
	// 14: (1,2) [C B D A]
	// 15: (2,3) [C B A D]
	// 16: (0,1) [B C A D]
	// 17: (2,3) [B C D A]
	// 18: (1,2) [B D C A]
	// 19: (0,1) [D B C A]
	// 20: (2,3) [D B A C]
	// 21: (0,1) [B D A C]
	// 22: (1,2) [B A D C]
	// 23: (2,3) [B A C D]
}

func ExampleSteinhausJohnsonTrotterEven() {
	i := 0

	A := []string{"A", "B", "C", "D"}

	p := NewSteinhausJohnsonTrotterEven(len(A))
	fmt.Printf("%2d:       %v\n", i, A)

	var sw [2]int
	for p.Next(&sw) {
		i++
		SwapStrings(sw, A)
		fmt.Printf("%2d: (%d,%d) %v\n", i, sw[0], sw[1], A)
	}

	//Output:
	//  0:       [A B C D]
	//  1: (2,3) [A B D C]
	//  2: (1,2) [A D B C]
	//  3: (0,1) [D A B C]
	//  4: (2,3) [D A C B]
	//  5: (0,1) [A D C B]
	//  6: (1,2) [A C D B]
	//  7: (2,3) [A C B D]
	//  8: (0,1) [C A B D]
	//  9: (2,3) [C A D B]
	// 10: (1,2) [C D A B]
	// 11: (0,1) [D C A B]
	// 12: (2,3) [D C B A]
	// 13: (0,1) [C D B A]
	// 14: (1,2) [C B D A]
	// 15: (2,3) [C B A D]
	// 16: (0,1) [B C A D]
	// 17: (2,3) [B C D A]
	// 18: (1,2) [B D C A]
	// 19: (0,1) [D B C A]
	// 20: (2,3) [D B A C]
	// 21: (0,1) [B D A C]
	// 22: (1,2) [B A D C]
	// 23: (2,3) [B A C D]
}
