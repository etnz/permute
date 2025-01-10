package permute

import "iter"

// SubsetRevolvingDoorNext computes the next combination 'p' from 'n'.
//
// return false if all combinations have been generated
func SubsetRevolvingDoorNext(p []int, n int) bool {
	j, k := 0, len(p)

	for ; j < k && p[j] == j; j++ {
	}
	if (k-j)%2 == 0 {
		if j == 0 {
			p[0]--
		} else {
			iset(p, j-2, j-1, j)
		}

	} else {
		pj, pj1 := n+1, n+1
		switch j {
		case k:
		case k - 1:
			pj = p[j]
		default:
			pj = p[j]
			pj1 = p[j+1]
		}

		if pj1 != pj+1 {
			if pj+1 >= n { // detects the end of the revolving door

				p[k-1] = k - 1
				return false
			}
			iset(p, j-1, pj, pj+1)
		} else {
			iset(p, j, j, pj)
		}
	}
	return true
}

// iset set into the array if possible (the algorithm is MUCH easier to write
// if we can 'write' out of the bounds of the p)
func iset(p []int, i, vi, vj int) {
	//I've got two new values vi, and vj
	// but one is already present (because this is a minimal change alg, there can't be two swaps)
	//so, either vj== p[i] (I'm moving p[i] to j) or the other one

	k := len(p)
	switch {
	case i == -1: //obviously i is outside
		p[0] = vj
	case i < k-1: // i and i+1 are still inside
		p[i], p[i+1] = vi, vj
		//case i == k-1: // j is outside
	}
}

// RevolvingDoorCombinations returns an iterator over all n-combinations of 'list' according to the
// Revolving door algorithm.
//
// ref Knuth, Donald Ervin. The Art of Computer Programming, volume 4, fascicle 3; generating all combinations and partitions, sec. 7.2.1.3, algorithm R, Revolving-door combinations,  p. 9.
func RevolvingDoorCombinations[Slice ~[]E, E any](n int, list Slice) iter.Seq[Slice] {
	return func(yield func(v Slice) bool) {
		s := newSubset(n)
		if !yield(Subset(s, list)) {
			return
		}
		for SubsetRevolvingDoorNext(s, len(list)) {
			if !yield(Subset(s, list)) {
				return
			}
		}
	}
}
