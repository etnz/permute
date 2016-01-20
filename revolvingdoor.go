package permute

// SubsetRevolvingDoorNext computes the next combination 'p' from 'n'.
//
// sw[0] will be the element to be replaced
//
// sw[1] will be the element (in the original list) to replace with
//
// return false if all combinations have been generated
func SubsetRevolvingDoorNext(p []int, n int, sw *[2]int) bool {
	j, k := 0, len(p)

	for ; j < k && p[j] == j; j++ {
	}
	if (k-j)%2 == 0 {
		if j == 0 {
			*sw = [2]int{p[0], p[0] - 1}
			p[0]--
		} else {
			iset(p, sw, j-2, j-1, j)
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

				*sw = [2]int{p[k-1], k - 1}
				p[k-1] = k - 1
				return false
			}
			iset(p, sw, j-1, pj, pj+1)
		} else {
			iset(p, sw, j, j, pj)
		}
	}
	return true
}

//iset set into the array if possible (the algorithm is MUCH easier to write
// if we can 'write' out of the bounds of the p)
func iset(p []int, sw *[2]int, i, vi, vj int) {
	//I've got two new values vi, and vj
	// but one is already present (because this is a minimal change alg, there can't be two swaps)
	//so, either vj== p[i] (I'm moving p[i] to j) or the other one

	//j := i + 1
	k := len(p)
	// fmt.Printf("iset(p[%v],p[%v]=%v,%v)", i, i+1, vi, vj)

	switch {
	case i == -1: //obviously i is outside
		*sw = [2]int{p[0], vj}
		p[0] = vj
	case i < k-1: // i and j are still inside
		// only one is deleted
		if p[i] == vj { // vj is not created, then p[j] is deleted
			*sw = [2]int{p[i+1], vi}
		} else { //the other way around
			*sw = [2]int{p[i], vj}
		}
		p[i], p[i+1] = vi, vj
	case i == k-1: // j is outside
		*sw = [2]int{p[i], vj}

	}

	// if i >= 0 && i < len(p) {
	// 	if vj != p[i] {
	// 		*sw = [2]int{p[i], vj}
	// 	}
	// 	p[i] = vi
	// }
	// if j >= 0 && j < len(p) {
	// 	if vi != p[j] {
	// 		*sw = [2]int{p[j], vi}
	// 	}
	// 	p[j] = vj
	// }
}
