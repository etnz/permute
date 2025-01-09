package permute

// SubsetRevolvingDoorNext computes the next combination 'p' from 'n'.
//
// sw[0] will be the element to be replaced
//
// sw[1] will be the element (in the original list) to replace with
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
