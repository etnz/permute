// package permute provide tools to deal with permutations
//
// A permutation is exactly a '[]int', where each values are unique, *and* in the interval [0, len[
//
//
//     []int{0, 1, 2} is the identity
//
// A permutation transforms
//
//     permutation []int{ 2  ,  1  ,  0  }
//     tranforms     x= {"a" , "b" , "c" }
//     into             {x[2], x[1], x[0]}
//     resulting in     {"c" , "b" , "a" }
//
//
// In addition to applying a permutation to a collection it offers four different ways to generate all permutations :
//
// Lexicographical
//
// Generates all permutation in lexicographical order. Not the fastest way to generate permutations.
//
// Permutations in lexicographical order can imply several transposition to apply.
//
// SteinhausJohnsonTrotter
//
// Generates all permutation so that each permutation in the sequence differs from the previous permutation by swapping two adjacent elements.
//
// This is the slowest method to generate permutations.
//
// But it does not require extra memory (like with Even speed up).
//
// The generated sequence is the fastest to apply: there is only one transposition each time, and of adjacent items.
//
// SteinhausJohnsonTrotterEven
//
// Add to SteinhausJohnsonTrotter method a little extra memory (O(n)) that greatly speeds up the permutation generation.
//
// Heap
//
// Generates all permutation so that each permutation in the sequence differs from the previous permutation by swapping two element (not necessarily adjacent).
//
// By far, Heap is the fastest of all to generate permutation. If there is no huge benefits in swapping adjacent items ( random access to collection items), then
// it the best choice.
//
// Here are some benchmarks executed on my computer (relative numbers matter)
//
//       BenchmarkPermGenHeap	100000000	        12 ns/op
//       BenchmarkPermGenEven	20000000	       117 ns/op
//       BenchmarkPermGenLex	10000000	       176 ns/op
//       BenchmarkPermGenSJT	 2000000	       787 ns/op
//
package permute

//Inv returns a new permutation 'q' that is the inverse of 'p'
//
// Meaning that q( p(x)) = x
func Inv(p []int) (q []int) {

	q = make([]int, len(p))
	for i, v := range p {
		q[v] = i
	}
	return
}

//New creates a new Permutation Identity of size 'n'
func New(n int) []int {
	x := make([]int, n)
	for i := range x {
		x[i] = i
	}
	return x
}

// Equals returns true if 'p' and 'q' are the same permutations
func Equals(p, q []int) bool {
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

// A type, typically a collection, that satisfies Interface can be permuted by 'Apply' function.
//
// All types implementing sort.Interface also implements this one
type Interface interface {
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

//Apply permutation p to 'val'
//
func Apply(p []int, val Interface) {
	for _, s := range Transpositions(p) {
		val.Swap(s[0], s[1])
	}
}

// Strings applies permutation 'p' to 'val'
func Strings(p []int, val []string) {
	q := make([]string, len(val))
	copy(q, val)
	for i, pi := range p {
		val[i] = q[pi]
	}
}

// Floats applies permutation 'p' to 'val'
func Floats(p []int, val []float64) {
	q := make([]float64, len(val))
	copy(q, val)
	for i, pi := range p {
		val[i] = q[pi]
	}
}

// Ints applies permutation 'p' to 'val'
func Ints(p []int, val []int) {
	q := make([]int, len(val))
	copy(q, val)
	for i, pi := range p {
		val[i] = q[pi]
	}
}
