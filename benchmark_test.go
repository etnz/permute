package permute

import (
	"fmt"
	"iter"
	"strings"
	"testing"
)

// Benchmarks the different algorithms.

func args(a ...any) string {
	s := make([]string, len(a))
	for i, x := range a {
		s[i] = fmt.Sprint(x)
	}
	return fmt.Sprintf(";%v;", strings.Join(s, ";"))
}

func benchPerm2(b *testing.B, name string, n int, gen func(list P) iter.Seq2[T, P]) {
	b.Run(args(name, n), func(b *testing.B) {
		x := newPermutation(n)
		i := 0
		for i < b.N {
			for _ = range gen(x) {
				if i >= b.N {
					return
				}
				i++
			}
		}
	})
}
func benchPerm(b *testing.B, name string, n int, gen func(list P) iter.Seq[P]) {
	b.Run(args(name, n), func(b *testing.B) {
		x := newPermutation(n)
		i := 0
		for i < b.N {
			for _ = range gen(x) {
				if i >= b.N {
					return
				}
				i++
			}
		}
	})
}

// for permutations
func BenchmarkPermutations(b *testing.B) {

	for n := 0; n < 20; n++ {
		benchPerm2(b, "Heap", n, HeapPermutations)
		benchPerm(b, "Lex", n, LexPermutations)
		benchPerm2(b, "SJT", n, SteinhausJohnsonTrotterPermutations)
		benchPerm2(b, "SJTE", n, SteinhausJohnsonTrotterEvenPermutations)
	}
}

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

func BenchmarkCombinations(b *testing.B) {

	for n := 0; n < 20; n++ {
		for k := 0; k <= n; k++ {

			b.Run(fmt.Sprintf(";Lex;%v;%v;", n, k), func(b *testing.B) {
				x := newPermutation(n)
				i := 0
				for i < b.N {
					for _ = range LexCombinations(k, x) {
						if i >= b.N {
							return
						}
						i++
					}
				}

			})
			b.Run(fmt.Sprintf(";Rev;%v;%v;", n, k), func(b *testing.B) {
				x := newPermutation(n)
				i := 0
				for i < b.N {

					for _ = range RevolvingDoorCombinations(k, x) {
						if i >= b.N {
							return
						}
						i++
					}
				}

			})
		}
	}
}
