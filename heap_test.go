package permute

import "fmt"

func ExampleHeapPermutations() {
	A := []string{"A", "B", "C", "D"}

	for t, v := range HeapPermutations(A) {
		fmt.Println(t, v)
	}

	//Output:
	// [0 0] [A B C D]
	// [0 1] [B A C D]
	// [0 2] [C A B D]
	// [0 1] [A C B D]
	// [0 2] [B C A D]
	// [0 1] [C B A D]
	// [0 3] [D B A C]
	// [0 1] [B D A C]
	// [0 2] [A D B C]
	// [0 1] [D A B C]
	// [0 2] [B A D C]
	// [0 1] [A B D C]
	// [1 3] [A C D B]
	// [0 1] [C A D B]
	// [0 2] [D A C B]
	// [0 1] [A D C B]
	// [0 2] [C D A B]
	// [0 1] [D C A B]
	// [2 3] [D C B A]
	// [0 1] [C D B A]
	// [0 2] [B D C A]
	// [0 1] [D B C A]
	// [0 2] [C B D A]
	// [0 1] [B C D A]
}
