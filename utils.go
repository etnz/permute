package permute

// permutations utils

// this file provides util functions using generics

// Permute permutation p to 'val'
func Permute[E any](p []int, val []E) {
	for _, s := range Decompose(p) {
		i, j := s[0], s[1]
		val[i], val[j] = val[j], val[i]
	}
}

// Subset applies subset 'p' to 'val' and returns it
func Subset[Slice ~[]E, E any](p []int, val Slice) Slice {
	q := make(Slice, len(p))
	for i, pi := range p {
		q[i] = val[pi]
	}
	return q
}

// Transpose applies the transposition 't' to 'p'
func Transpose[Slice ~[]E, E any](t T, p Slice) { p[t[0]], p[t[1]] = p[t[1]], p[t[0]] }

// Inv returns a new permutation 'q' that is the inverse of 'p'
//
// Meaning that q( p(x)) = x
func Inv(p []int) (q []int) {
	q = make([]int, len(p))
	for i, v := range p {
		q[v] = i
	}
	return
}

// New creates a new Permutation Identity of size 'n'
func New(n int) []int {
	x := make([]int, n)
	for i := range x {
		x[i] = i
	}
	return x
}

// Equals returns true if 'p' and 'q' are the same permutations
func Equals(p, q []int) bool {
	// for backward compat, we cannot use slice.Equals
	if len(p) != len(q) {
		return false
	}
	for i := range p {
		if p[i] != q[i] {
			return false
		}
	}
	return true
}
