package permute

import "testing"

// Benchmarks the different algorithms.

// for permutations

func BenchmarkHeapPermutations(b *testing.B) {
	x := newPermutation(4)
	for i := 0; i < b.N; i++ {
		for _ = range HeapPermutations(x) {
		}
	}
}
func BenchmarkLexPermutations(b *testing.B) {
	x := newPermutation(4)
	for i := 0; i < b.N; i++ {
		for _ = range LexPermutations(x) {
		}
	}
}
func BenchmarkSteinhausJohnsonTrotterPermutations(b *testing.B) {
	x := newPermutation(4)
	for i := 0; i < b.N; i++ {
		for _ = range SteinhausJohnsonTrotterPermutations(x) {
		}
	}
}
func BenchmarkSteinhausJohnsonTrotterEvenPermutations(b *testing.B) {
	x := newPermutation(4)
	for i := 0; i < b.N; i++ {
		for _ = range SteinhausJohnsonTrotterEvenPermutations(x) {
		}
	}
}

// Heap
// Lexicographical
// SteinhausJohnsonTrotter
// SteinhausJohnsonTrotterEven
//
// for combinations
// Lexicographical
// RevolvingDoor

func BenchmarkLexCombinations(b *testing.B) {
	x := newPermutation(12)
	for i := 0; i < b.N; i++ {
		for _ = range LexCombinations(6, x) {
		}
	}
}
func BenchmarkRevolvingDoorCombinations(b *testing.B) {
	x := newPermutation(12)
	for i := 0; i < b.N; i++ {
		for _ = range RevolvingDoorCombinations(6, x) {
		}
	}
}
