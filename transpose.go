package permute

//Swap applies the transposition 's' to 'p'
func Swap(s [2]int, p Interface) { p.Swap(s[0], s[1]) }

//NewSwap creates a new Transposition where the first element is always lower than the second one
//
// A Swap can be used to swap elements of any collection
func NewSwap(a, b int) [2]int {
	if b < a {
		a, b = b, a
	}
	return [2]int{a, b}
}

//SwapInts applies the transposition 's' to 'p'
func SwapInts(s [2]int, p []int) {
	k, pk := s[0], s[1]
	p[k], p[pk] = p[pk], p[k]
}

//SwapStrings applies the transposition 's' to 'p'
func SwapStrings(s [2]int, p []string) {
	k, pk := s[0], s[1]
	p[k], p[pk] = p[pk], p[k]
}

//SwapFloats applies the transposition 's' to 'p'
func SwapFloats(s [2]int, p []float64) {
	k, pk := s[0], s[1]
	p[k], p[pk] = p[pk], p[k]
}

// a small part of permutation is related to transposition
// a transposition is a pair of indexes to transpose

// Transpositions generate the sequence of transpositions (swap) equivalent to 'p'
//
// A single transposition is a pair of indexes to be swapped
func Transpositions(p []int) (swaps [][2]int) {

	//we don't know the exact number of swaps
	swaps = make([][2]int, 0, len(p))

	current := make([]int, len(p))
	copy(current, p)

	for k := smallestNonFixedIndex(current); k >= 0; k = smallestNonFixedIndex(current) {
		sk := indexof(k, current)
		perm := [2]int{k, sk}
		swaps = append(swaps, perm)
		SwapInts(perm, current) //inplace permute
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

//smallestNonFixedIndex used in Transpositions decomposition
func smallestNonFixedIndex(p []int) int {
	for i, pi := range p {
		if i != pi {
			return i
		}
	}
	return -1 //if none is non fix
}
