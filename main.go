// Package provides algorithm to generates all permutations and combinations.
package permute

import "iter"

// T represent a single transposition
type (
	T [2]int
	P []int
	S []int
)

// inversions in 'p' returns the number of inversions.
//
// A pair of indices (i,j) with i < j and p[i] > p[j] is called an inversion
func inversions(p P) (count int) {
	for i, pi := range p {
		for j := 0; j < i; j++ {
			if p[j] > pi {
				count++
			}
		}
	}
	return
}

// parity returns the number of transposition mod 2
//
// In fact this is also the number of inversion mod 2
func parity(p P) (parity int) { return inversions(p) % 2 }

// even returns true if the permutation is even see https://en.wikipedia.org/wiki/Parity_of_a_permutation
func even(p P) bool { return parity(p) == 0 }

// transpose two positions in p
//
// equivalent to Swap() but without generic to keep the algorithm local.
func transpose(s T, p P) { p[s[0]], p[s[1]] = p[s[1]], p[s[0]] }

// newPermutation creates a new Permutation Identity of size 'n'
func newPermutation(n int) P {
	x := make(P, n)
	for i := range x {
		x[i] = i
	}
	return x
}

// newSubset creates a new Subset S Identity of size 'n'
func newSubset(n int) S {
	x := make(S, n)
	for i := range x {
		x[i] = i
	}
	return x
}

// newTransposition creates a new transposition T where the first element is always lower than the second one
func newTransposition(a, b int) T {
	if b < a {
		a, b = b, a
	}
	return T{a, b}
}

// a small part of permutation is related to transposition
// a transposition is a pair of indexes to transpose

// smallestNonFixedIndex used in Transpositions decomposition.
//
// Returns the first index that is different from the identity permutation.
func smallestNonFixedIndex(p P) int {
	for i, pi := range p {
		if i != pi {
			return i
		}
	}
	return -1 //if none is non fix
}

// Decompose 'p' into an equivalent sequences of Transpositions.
func Decompose(p P) (transpositions []T) {

	//we don't know the exact number of swaps
	transpositions = make([]T, 0, len(p))

	// iteratively convert 'current' towards the identity using atomic Transpositions,
	// record them, and return them in reverse order.
	current := make(P, len(p))
	copy(current, p)
	for {
		// for k := smallestNonFixedIndex(current); k >= 0; k = smallestNonFixedIndex(current) {
		// Get the first index that is different from the identity permutation.
		k := smallestNonFixedIndex(current)
		if k < 0 {
			// There are none, so current has reached the identity.
			break
		}
		sk := current[k]

		perm := newTransposition(k, sk)
		transpositions = append(transpositions, perm)
		transpose(perm, current) //inplace permute
	}
	//reverse the array
	ns := len(transpositions) - 1
	for i := 0; i < (ns+1)/2; i++ {
		transpositions[i], transpositions[ns-i] = transpositions[ns-i], transpositions[i]
	}
	return
}

// Permute permutation p to 'val'.
func Permute[Slice ~[]E, E any](p P, val Slice) {
	for _, s := range Decompose(p) {
		i, j := s[0], s[1]
		val[i], val[j] = val[j], val[i]
	}
}

// Subset applies subset 'p' to 'val' and returns it.
func Subset[Slice ~[]E, E any](p S, val Slice) Slice {
	q := make(Slice, len(p))
	for i, pi := range p {
		q[i] = val[pi]
	}
	return q
}

// Transpose applies the transposition 't' to 'p'.
func Transpose[Slice ~[]E, E any](t T, p Slice) {
	if t[0] == t[1] {
		return
	}
	p[t[0]], p[t[1]] = p[t[1]], p[t[0]]
}

// Permutations returns an interator over all permutations of 'list'.
//
// Each permutation is computed from the previous one + one Transposition.
// The iterator returns both values.
func Permutations[Slice ~[]E, E any](list Slice) iter.Seq2[T, Slice] { return HeapPermutations(list) }

// Combinations returns an iterator over all n-combinations of 'list'.
func Combinations[Slice ~[]E, E any](n int, l Slice) iter.Seq[Slice] { return RevCombinations(n, l) }
