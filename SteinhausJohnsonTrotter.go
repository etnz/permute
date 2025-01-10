package permute

import "iter"

// SteinhausJohnsonTrotter computes the next permutation according to the Steinhaus-Johnson-Trotter algorithm
//
// it provides a way to generate ALL the permutations by swapping two adjacent values from one to another
//
// Each call to SteinhausJohnsonTrotter updates, in place 'p', and 't'
//
// It returns 'false' if it has gone back to the identity permutation.
//
// 'p' is the current permutation, whereas 't' is the current transposition to go from the previous one to the current one.
// While looping over all permutations of a collection, it is cheaper to apply just the transposition !
func SteinhausJohnsonTrotter(p []int, t *T) bool {
	if len(p) < 2 {
		return false
	}
	s, identity := steinhausJohnsonTrotter(p)
	t[0], t[1] = s[0], s[1]
	return !identity
}

// recursive version
func steinhausJohnsonTrotter(p []int) (sw T, identity bool) {

	// as I understand it:
	// the algorithm is to ALWAYS swap the max number:e.g.
	//      123[4]
	//      12[4]3
	//      1[4]23
	//      [4]123
	//
	// now we have consumed "123" and we move back using the 'next' one (132) (132  is recusively computed)
	//
	// [4]132
	// 1[4]32
	// 13[4]2
	// 132[4]
	//
	// so: for a given permutation, we need to find out if the remaining one (123 in the first eg, 132 in the second one)
	// is left to right or right to left (i.e the oddity of its rank)
	// then we just apply
	//
	// and finally if we are at the end of the position, we need to compute the next sub-permutation (here 123 -> 132 ) that is hopefully the same operation
	N := len(p)
	//compute the current position of the last value (N-1) (4 in the example)
	s := indexof(N-1, p)

	if N == 2 {
		//recursion end here
		//always swap the two values( 0,1) or (1,0)
		// p will be identity if 1's position is currently 0
		sw = T{0, 1}
		transpose(sw, p)
		identity = s == 0
		return
	}

	// build the current subquery
	sub := make([]int, 0, N)
	for i, v := range p {
		if i != s {
			sub = append(sub, v)
		}
	}

	//two very different cases whether this is a even sub permutation or not
	if even(sub) {

		if s == 0 { // this is the boundary of it
			sw, identity = steinhausJohnsonTrotter(sub)
			// unfortunately, the sub permutation is fully on the right so the tranposition need to be updated
			sw[0]++
			sw[1]++
		} else {
			// not on the boundaries
			sw = T{s - 1, s}
		}
	} else { // case ODD
		if s == N-1 { // this is the boundary of it
			sw, identity = steinhausJohnsonTrotter(sub)
			// fortunately, the sub permutation is fully on the right so the tranposition need not to be updated
		} else {
			// not on the boundaries
			sw = T{s, s + 1}
		}
	}
	transpose(sw, p)
	return
}

// SteinhausJohnsonTrotterPermutations returns an interator over all permutations of 'list' using the
// Steinhaus-Johnson-Trotter algorithm.
func SteinhausJohnsonTrotterPermutations[Slice ~[]E, E any](list Slice) iter.Seq2[T, Slice] {
	return func(yield func(t T, v Slice) bool) {
		var t T
		if !yield(t, list) {
			return
		}
		p := newPermutation(len(list))

		for SteinhausJohnsonTrotter(p, &t) {
			Transpose(t, list)
			if !yield(t, list) {
				return
			}
		}
	}
}

// SteinhausJohnsonTrotterEven implements a minimal change generator based on Even speed up
type SteinhausJohnsonTrotterEven struct {
	P, D []int //permutation and direction marker
}

func NewSteinhausJohnsonTrotterEven(n int) *SteinhausJohnsonTrotterEven {
	if n == 0 {
		return &SteinhausJohnsonTrotterEven{}
	}
	dir := make([]int, n)
	//initialise the direction, the biggest number is
	for i := range dir {
		dir[i] = -1
	}
	dir[0] = 0
	return &SteinhausJohnsonTrotterEven{
		D: dir,
		P: newPermutation(n),
	}
}

// Next return false when we have gone back to the identity
//
// sw is updated with the transposition from previous permutation to the next one
func (s *SteinhausJohnsonTrotterEven) Next(sw *T) bool {
	N := len(s.P)
	if N < 2 {
		return false
	}
	last := true
	for i := range s.P {
		if s.D[i] != 0 {
			last = false
		}
	}
	if last {
		//reset to start again:
		for i := range s.D {
			s.D[i] = -1
		}
		s.D[0] = 0
		*sw = newTransposition(0, 1)
		return false
	}
	//position of the max
	// value of the max
	// direction of the max
	maxi, max, maxd := -1, -1, 0
	for i, d := range s.D {
		if d != 0 && (maxi < 0 || s.P[i] > max) {
			//this is a max
			maxi, max, maxd = i, s.P[i], d
		}
	}
	// I've got the max I swap in that direction
	i := maxi + maxd
	*sw = newTransposition(maxi, i)
	transpose(*sw, s.P)
	//and the same goes for te direction
	transpose(*sw, s.D)

	// shall I set this new position to zero ?
	//if element to reach the first or last position within the permutation, or if the next element in the same direction is larger than the chosen element, the direction of the chosen element is set to zero
	if i == 0 || i == N-1 || s.P[i+maxd] > max {
		s.D[i] = 0
	}

	//After each step, all elements greater than the chosen element have their directions set to positive or negative, according to whether they are between the chosen element and the start or the end of the permutation respectively.
	for i, pi := range s.P {
		if pi > max && i < maxi {
			s.D[i] = 1
		}
		if pi > max && i > maxi {
			s.D[i] = -1
		}

	}

	return !last
}

// SteinhausJohnsonTrotterEvenPermutations returns an interator over all permutations of 'list' using the
// Steinhaus-Johnson-Trotter algorithm with the Even speed up.
func SteinhausJohnsonTrotterEvenPermutations[Slice ~[]E, E any](list Slice) iter.Seq2[T, Slice] {
	return func(yield func(t T, v Slice) bool) {
		var t T
		if !yield(t, list) {
			return
		}
		s := NewSteinhausJohnsonTrotterEven(len(list))
		for s.Next(&t) {
			Transpose(t, list)
			if !yield(t, list) {
				return
			}
		}
	}
}
