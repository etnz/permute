package permute

import (
	"fmt"
)

//the fastest way to compute int64 factorial is by statically storing all the values (there are really few of them)

const upperLimit = 21 // max number of factorial values below the uint64 maximum value: 21! > MaxUint64

var memoization = [upperLimit]uint64{
	1,
	1,
	2,
	6,
	24,
	120,
	720,
	5040,
	40320,
	362880,
	3628800,
	39916800,
	479001600,
	6227020800,
	87178291200,
	1307674368000,
	20922789888000,
	355687428096000,
	6402373705728000,
	121645100408832000,
	2432902008176640000,
}

//Factorial computes n! n beeing in [0,20], simply because there is no uint64 to represent 21!
func Factorial(n int) uint64 {
	if n >= upperLimit {
		panic(fmt.Errorf("uint64 overflow in %v!", n))
	}
	return memoization[n]
}
