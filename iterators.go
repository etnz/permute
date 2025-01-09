//go:build go1.23

package permute

import (
	"iter"
)

// this file handles iterators (introduced in 1.23) and generics (introduced in 1.18)

// Permutations returns an interator over all permutations of 'list'.
//
// Each permutation is computed from the previous one + one Transposition.
// The iterator returns both values.
func Permutations[Slice ~[]E, E any](list Slice) iter.Seq2[T, Slice] {
	return func(yield func(t T, v Slice) bool) {
		h := NewHeap(len(list))
		var t T
		if !yield(t, list) {
			return
		}
		for h.Next(&t) {
			Transpose(t, list)
			if !yield(t, list) {
				return
			}
		}
	}
}

// Combinations returns an iterator over all n-combinations of 'list'.
func Combinations[Slice ~[]E, E any](n int, list Slice) iter.Seq[Slice] {
	return func(yield func(v Slice) bool) {
		p := New(n)
		if !yield(Subset(p, list)) {
			return
		}
		for SubsetRevolvingDoorNext(p, len(list)) {

			if !yield(Subset(p, list)) {
				return
			}
		}
	}
}
