package permute

import "testing"

func BenchmarkPermGenLex(b *testing.B) {
	p := New(20)
	for i := 0; i < b.N; i++ {
		Lexicographical(p)
	}
}

func BenchmarkPermGenSJT(b *testing.B) {
	p := New(20)
	var sw [2]int
	for i := 0; i < b.N; i++ {
		SteinhausJohnsonTrotter(p, &sw)
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
	h := NewSteinhausJohnsonTrotterEven(20)
	var sw [2]int
	for i := 0; i < b.N; i++ {
		h.Next(&sw)
	}
}
