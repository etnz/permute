package permute

//Even returns true if the permutation is even see https://en.wikipedia.org/wiki/Parity_of_a_permutation
func Even(p []int) bool { return Parity(p) == 0 }

//Odd returns true if the permutation is odd see https://en.wikipedia.org/wiki/Parity_of_a_permutation
func Odd(p []int) bool { return Parity(p) == 1 }

// LehmerCode converts a permutation 'p' into its lehmer code.
//
// The Lehmer Code vector 'L' is defined so that L[i] is "the number of entries to the right of p[i], which are smaller
//
// L[i] = card(j>i: p[j]<p[i])
//
func LehmerCode(p []int) (L []int) {
	N := len(p)
	L = make([]int, N)
	for i, pi := range p {
		s := 0
		for j := i + 1; j < N; j++ {
			if p[j] < pi {
				s++
			}
		}
		L[i] = s
	}
	return
}

// NewLehmer Coded return a new permutation defined by it's lehmer code
//
// len(p) == len(L)
func NewLehmerCoded(L []int) (p []int) {

	N := len(L)
	nn := New(N) // temporary list of values
	p = make([]int, N)

	for i, li := range L {
		p[i] = nn[li]
		delv(&nn, p[i])
	}
	return
}

//delv remove value 'x' from p
func delv(p *[]int, x int) {
	for i, v := range *p {
		if v == x {
			//this is the tiping point
			*p = append((*p)[0:i], (*p)[i+1:]...)
			return
		}
	}
	return
}

//CountInversions in 'p'
//
//A pair of indices (i,j) with i < j and p[i] > p[j] is called an inversion
func CountInversions(p []int) (count int) {
	for i, pi := range p {
		for j := 0; j < i; j++ {
			if p[j] > pi {
				count++
			}
		}
	}
	return
}

//Parity returns the number of transposition mod 2
//
// In fact this is also the number of inversion mod 2
func Parity(p []int) (parity int) { return CountInversions(p) % 2 }
