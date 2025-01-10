package permute

import (
	"fmt"
	"iter"
	"strings"
	"testing"
)

// Benchmarks the different algorithms.

// args creates a name for a bench usually based on name, and variable parameters in
// a way that can be parse back easily.
func args(a ...any) string {
	s := make([]string, len(a))
	for i, x := range a {
		s[i] = fmt.Sprint(x)
	}
	return fmt.Sprintf(";%v;", strings.Join(s, ";"))
}

// benchPerm actually bench a permutation iterator based on iter.Seq
func benchPerm(b *testing.B, name string, n int, gen func(list P) iter.Seq[P]) {
	b.Run(args(name, n), func(b *testing.B) {
		x := newPermutation(n)
		i := 0
		b.ResetTimer()
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

// bencPerm2 actually bench a permutation iterator based on iter.Seq2
func benchPerm2(b *testing.B, name string, n int, gen func(list P) iter.Seq2[T, P]) {
	b.Run(args(name, n), func(b *testing.B) {
		x := newPermutation(n)
		i := 0
		b.ResetTimer()
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

// BenchmarkPermutations bench all permutation algorithms against a
// range of dimensions.
func BenchmarkPermutations(b *testing.B) {
	for n := 0; n < 20; n++ {
		benchPerm2(b, "Heap", n, HeapPermutations)
		benchPerm(b, "Lex", n, LexPermutations)
		benchPerm2(b, "SJT", n, SJTPermutations)
		benchPerm2(b, "SJTE", n, SJTEPermutations)
	}
}

func benchCombination(b *testing.B, name string, n, k int, gen func(int, P) iter.Seq[P]) {
	b.Run(args(name, n, k), func(b *testing.B) {
		x := newPermutation(n)
		i := 0
		b.ResetTimer()
		for i < b.N {
			for _ = range gen(k, x) {
				if i >= b.N {
					return
				}
				i++
			}
		}
	})
}

// BenchmarkCombinations benchmark all combination algorithms against ranges
// of values for 'n' and 'k'
func BenchmarkCombinations(b *testing.B) {
	for n := 0; n < 20; n++ {
		for k := 0; k <= n; k++ {
			benchCombination(b, "Lex", n, k, LexCombinations)
			benchCombination(b, "Rev", n, k, RevolvingDoorCombinations)
		}
	}
}
