package permute

// T represent a single transposition
type T [2]int

// NewTransposition creates a new transposition T where the first element is always lower than the second one
func NewTransposition(a, b int) T {
	if b < a {
		a, b = b, a
	}
	return T{a, b}
}

// a small part of permutation is related to transposition
// a transposition is a pair of indexes to transpose

// Transpositions generate the sequence of transpositions (swap) equivalent to 'p':
// apply them all to the identity permutation, leads to 'p'.
func Transpositions(p []int) (swaps []T) {

	//we don't know the exact number of swaps
	swaps = make([]T, 0, len(p))

	// iteratively convert 'current' towards the identity using atomic Transpositions,
	// record them, and return them in reverse order.
	current := make([]int, len(p))
	copy(current, p)
	for {
		// for k := smallestNonFixedIndex(current); k >= 0; k = smallestNonFixedIndex(current) {
		// Get the first index that is different from the identity permutation.
		k := smallestNonFixedIndex(current)
		if k < 0 {
			// There are none, so current has reached the identity.
			break
		}
		// TODO: there is another probably faster algo:
		// to avoid searching for the right value: it's about pushing the current value.
		//sk := indexof(k, current)
		sk := current[k]

		perm := NewTransposition(k, sk)
		swaps = append(swaps, perm)
		swap(perm, current) //inplace permute
	}
	//reverse the array
	ns := len(swaps) - 1
	for i := 0; i < (ns+1)/2; i++ {
		swaps[i], swaps[ns-i] = swaps[ns-i], swaps[i]
	}
	return
}

// indexof compute the index in 'p' of a given value 'p'
func indexof(x int, p []int) int {
	for i, pi := range p {
		if pi == x {
			return i
		}
	}
	return -1
}

// smallestNonFixedIndex used in Transpositions decomposition.
//
// Returns the first index that is different from the identity permutation.
func smallestNonFixedIndex(p []int) int {
	for i, pi := range p {
		if i != pi {
			return i
		}
	}
	return -1 //if none is non fix
}
