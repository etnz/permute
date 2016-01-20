package permute

import "testing"

func BenchmarkPermGenLex(b *testing.B) {
	p := New(20)
	for i := 0; i < b.N; i++ {
		LexNext(p)
	}
}

func BenchmarkPermGenSJT(b *testing.B) {
	h := NewPlainChangeGen(20)
	var sw [2]int
	for i := 0; i < b.N; i++ {
		h.Next(&sw)
	}
}

func BenchmarkPermGenHeap(b *testing.B) {
	h := NewHeap(20)
	var sw [2]int
	for i := 0; i < b.N; i++ {
		h.Next(&sw)
	}
}
func BenchmarkPermGenEven(b *testing.B) {
	h := NewPlainChangeFastGen(20)
	var sw [2]int
	for i := 0; i < b.N; i++ {
		h.Next(&sw)
	}
}
