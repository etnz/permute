package permute

// from B. R. Heap in 1963
// algorithm for the non recursive  adapted from
// Sedgewick, Robert. "a talk on Permutation Generation Algorithms"
// assbackward implementation is mine

// Heap is a struct to generate all the permutations in Heap's order
//
// see https://en.wikipedia.org/wiki/Heap%27s_algorithm for details
type Heap struct {
	c []int //  current index set
	n int   // deep position (index of c)
}

//NewHeap creates a new Heap generator to generate all permutations of length n
func NewHeap(n int) *Heap { return &Heap{c: make([]int, n)} }

// Next return false when we have gone back to the identity
//
// sw is updated with the transposition from previous permutation to the next one
func (h *Heap) Next(swap *[2]int) (ok bool) {
	N := len(h.c)
	for h.n < N {
		if h.c[h.n] < h.n {
			s := h.c[h.n]
			if h.n%2 == 0 {
				s = 0
			}
			*swap = [2]int{s, h.n}
			h.c[h.n]++
			h.n = 0
			return true
		} else {
			h.c[h.n] = 0
			h.n++
		}
	}
	*swap = [2]int{0, 0}
	return false
}
