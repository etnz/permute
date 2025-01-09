// Package permute provide tools to deal with permutations
//
// A permutation is exactly a '[]int', where each values are unique, *and* in the interval [0, len[
//
//	[]int{0, 1, 2} is the identity
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
// # Lexicographical
//
// Generates all permutation in lexicographical order. Not the fastest way to generate permutations.
//
// Permutations in lexicographical order can imply several transposition to apply.
//
// # SteinhausJohnsonTrotter
//
// Generates all permutation so that each permutation in the sequence differs from the previous permutation by swapping two adjacent elements.
//
// This is the slowest method to generate permutations.
//
// But it does not require extra memory (like with Even speed up).
//
// The generated sequence is the fastest to apply: there is only one transposition each time, and of adjacent items.
//
// # SteinhausJohnsonTrotterEven
//
// Add to SteinhausJohnsonTrotter method a little extra memory (O(n)) that greatly speeds up the permutation generation.
//
// # Heap
//
// Generates all permutation so that each permutation in the sequence differs from the previous permutation by swapping two element (not necessarily adjacent).
//
// By far, Heap is the fastest of all to generate permutation. If there is no huge benefits in swapping adjacent items ( random access to collection items), then
// it the best choice.
//
// Here are some benchmarks executed on my computer (relative numbers matter)
//
// goos: linux
// goarch: amd64
// pkg: github.com/etnz/permute
// cpu: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz
// BenchmarkPermGenLex             74396396       16.97 ns/op
// BenchmarkPermGenSJT              6341468       194.7 ns/op
// BenchmarkPermGenHeap           321184149       3.888 ns/op
// BenchmarkPermGenEven            18083234       69.13 ns/op
// BenchmarkFact                   65677332       18.85 ns/op
// BenchmarkFactorial            1000000000       1.166 ns/op
// BenchmarkParity                 68620435       17.41 ns/op
// BenchmarkParity2                12564770       100.4 ns/op
// BenchmarkAsTransposition         9052173       144.7 ns/op
// BenchmarkAsTransposition2        2989990       412.1 ns/op
package permute
