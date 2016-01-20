package permute

// MinimalChangeGen is the interface implemented by all generation algorithm with the minimal change constraints.
//
// It has a single Next method that updates 'sw' pointer to the swap to be apply to generate the next item.
//
// It returns false when it has generated all the elements.
type MinimalChangeGen interface {
	Next(sw *[2]int) bool
}

// plainChangeGen implements minimal change based on sjt
type plainChangeGen struct{ p []int }

func (gen *plainChangeGen) Next(sw *[2]int) bool { return SteinhausJohnsonTrotter(gen.p, sw) }

// NewPlainChangeGen return a Plain change Order generator based on the Steinhaus-Johnson-Trotter algorithm.
func NewPlainChangeGen(size int) MinimalChangeGen { return &plainChangeGen{New(size)} }

// NewPlainChangeFastGen return a Plain change Order generator based on the Steinhaus-Johnson-Trotter algorithm with Even speed up.
func NewPlainChangeFastGen(n int) MinimalChangeGen {
	dir := make([]int, n)
	//initialise the direction, the biggest number is
	for i := range dir {
		dir[i] = -1
	}
	dir[0] = 0
	return &SteinhausJohnsonTrotterEven{
		D: dir,
		P: New(n),
	}
}

// NewHeapGen return a Plain change Order generator based on the Heap's algorithm.
func NewHeapGen(n int) MinimalChangeGen { return NewHeap(n) }

// minimalChangeGen implements minimal change generator for combination based on revolving door (for combinations)
type minimalChangeGen struct {
	n int
	p []int
}

func (gen *minimalChangeGen) Next(sw *[2]int) bool {
	return SubsetRevolvingDoorNext(gen.p, gen.n, sw)
}

// NewMinimalChangeGen return a Minimal Change Order generator for n,k-combinations based on Revolving door algorithm.
//
// This is a minimal change for n,k-subset generation. In that case, the 'sw' semantic is to replace sw[0] with sw[1]
func NewMinimalChangeGen(n, k int) MinimalChangeGen { return &minimalChangeGen{p: New(k), n: n} }
