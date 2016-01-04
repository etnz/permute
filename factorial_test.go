package permute

import (
	"fmt"
	"testing"
)

func ExampleFactorial() {
	//checking my "fact" function against wikipedia
	for i := 0; i < upperLimit; i++ {
		fmt.Printf("%v\t%v\n", i, fact(i))
	}
	// Output:
	// 0	1
	// 1	1
	// 2	2
	// 3	6
	// 4	24
	// 5	120
	// 6	720
	// 7	5040
	// 8	40320
	// 9	362880
	// 10	3628800
	// 11	39916800
	// 12	479001600
	// 13	6227020800
	// 14	87178291200
	// 15	1307674368000
	// 16	20922789888000
	// 17	355687428096000
	// 18	6402373705728000
	// 19	121645100408832000
	// 20	2432902008176640000
}

func TestFactorial(t *testing.T) {
	//test Factorial against fact()
	for i := 0; i < upperLimit; i++ {
		f := fact(i)
		F := Factorial(i)
		if f != F {
			t.Errorf("invalid Factorial, got %v, expecting %v", F, f)

		}
	}
}

func BenchmarkFact(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fact(i % upperLimit)
	}
}
func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(i % upperLimit)
	}
}

// BenchmarkFact	50000000	        35.0 ns/op
// BenchmarkFactorial	300000000	         4.00 ns/op
