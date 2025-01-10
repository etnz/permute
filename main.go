// Package permute provide tools to deal with permutations
//
// # Permutations
//
// A permutation type P is exactly a '[]int', where each values are unique, *and* in the interval [0, len[
//
//	P{0, 1, 2} is the identity
//
// A permutation transforms
//
//	permutation []int{ 2  ,  1  ,  0  }
//	tranforms     x= {"a" , "b" , "c" }
//	into             {x[2], x[1], x[0]}
//	resulting in     {"c" , "b" , "a" }
//
// In addition to applying a permutation to a collection it offers four different ways to generate all permutations :
//
// ## Lexicographical
//
// Generates all permutation in lexicographical order. Not the fastest way to generate permutations.
//
// Permutations in lexicographical order can imply several transposition to apply.
//
// ## SteinhausJohnsonTrotter
//
// Generates all permutation so that each permutation in the sequence differs from the previous permutation by swapping two adjacent elements.
//
// This is the slowest method to generate permutations.
//
// But it does not require extra memory (like with Even speed up).
//
// The generated sequence is the fastest to apply: there is only one transposition each time, and of adjacent items.
//
// ## SteinhausJohnsonTrotterEven
//
// Add to SteinhausJohnsonTrotter method a little extra memory (O(n)) that greatly speeds up the permutation generation.
//
// ## Heap
//
// Generates all permutation so that each permutation in the sequence differs from the previous permutation by swapping two element (not necessarily adjacent).
//
// By far, Heap is the fastest of all to generate permutation. If there is no huge benefits in swapping adjacent items ( random access to collection items), then
// it the best choice.
//
// Here are some benchmarks executed on my computer (relative numbers matter)
//
//	goos: linux
//	goarch: amd64
//	pkg: github.com/etnz/permute
//	cpu: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz
//	BenchmarkHeapPermutations-2                           4408521    276.2 ns/op
//	BenchmarkLexPermutations-2                             499938    2488 ns/op
//	BenchmarkSteinhausJohnsonTrotterPermutations-2         784820    1561 ns/op
//	BenchmarkSteinhausJohnsonTrotterEvenPermutations-2    2049980    591.7 ns/op
//
// # Combinations
//
// A (n,k)-combination or (n,k)-subset is a type S:
//
// - a `[]int` of length `k`
// - where each values are **unique**  *and* in the interval `[0, n[`
// - values are sorted in ascending order.
//
// For example, a combination:
//
//	combination []int{ 0  ,  2  }
//	transforms    x= {"a" , "b" , "c" }
//	into             {x[0], x[2]}
//	resulting in     {"a" , "c" }
//
// In addition to applying a subset to a collection it offers four different ways to generate all combinations:
//
// ## Lexicographical
//
// Generates all subsets in lexicographical order.
//
// ## RevolvingDoor
//
// Generates all subsets according to the E Knut's Revolving Door algorithm.
//
// Here is a quick benchmark of all the combinations.
//
//	BenchmarkLexCombinations-2                              33280    40189 ns/op
//	BenchmarkRevolvingDoorCombinations-2                    35706    34767 ns/op
//
// ## Transpositions
//
// Transpositions type T is a [2]int that contains two index of a collection to transpose.
//
//	BenchmarkAsTransposition-2                            7855764    164.8 ns/op
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
		// TODO: there is another probably faster algo:
		// to avoid searching for the right value: it's about pushing the current value.
		//sk := indexof(k, current)
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

// indexof compute the index in 'p' of a given value 'p'
func indexof(x int, p P) int {
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
func smallestNonFixedIndex(p P) int {
	for i, pi := range p {
		if i != pi {
			return i
		}
	}
	return -1 //if none is non fix
}

// Permute permutation p to 'val'
func Permute[Slice ~[]E, E any](p P, val Slice) {
	for _, s := range Decompose(p) {
		i, j := s[0], s[1]
		val[i], val[j] = val[j], val[i]
	}
}

// Subset applies subset 'p' to 'val' and returns it
func Subset[Slice ~[]E, E any](p S, val Slice) Slice {
	q := make(Slice, len(p))
	for i, pi := range p {
		q[i] = val[pi]
	}
	return q
}

// Transpose applies the transposition 't' to 'p'
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
func Combinations[Slice ~[]E, E any](n int, list Slice) iter.Seq[Slice] {
	return RevolvingDoorCombinations(n, list)
}
