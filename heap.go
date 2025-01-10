package permute

import "iter"

// HeapPermutations returns an interator over all permutations of 'list' using the [Heap] algorithm
//
// Each permutation is computed from the previous one + one Transposition.
// The iterator returns both values.
//
// [Heap]: https://en.wikipedia.org/wiki/Heap%27s_algorithm
func HeapPermutations[Slice ~[]E, E any](list Slice) iter.Seq2[T, Slice] {

	// see https://en.wikipedia.org/wiki/Heap%27s_algorithm for details
	//
	// from B. R. Heap in 1963
	//
	// algorithm for the non recursive  adapted from Sedgewick, Robert. "a talk on Permutation Generation Algorithms"
	//
	// Assbackward implementation is ours

	return func(yield func(t T, v Slice) bool) {
		var t T
		// Always return the first element in the list.
		if !yield(t, list) {
			return
		}
		n := 0 // deep position (index of c)
		N := len(list)
		c := make([]int, N) //  current index set

		for n < N {
			if c[n] < n {
				s := c[n]
				if n%2 == 0 {
					s = 0
				}
				t = T{s, n}
				c[n]++
				n = 0
				Transpose(t, list)
				if !yield(t, list) {
					return
				}
			}
			//else
			c[n] = 0
			n++
		}
	}
}
