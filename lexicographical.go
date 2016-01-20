package permute

import "sort"

// LexNext finds the next permutation in lexicographical order.
//
// return false if it has gone back to the identity permutation.
//
// inspired from Narayana Pandita in https://en.wikipedia.org/wiki/Permutation
func LexNext(p []int) bool {

	// the principle for lehmer code is to convert an factoradic number into its permutation:
	//
	// each factoradic digit stand for the selection of an index, among all possible.
	//
	// Excerpt from wikipedia:
	//
	// The process may become clearer with a longer example. For example, here is how the digits in the factoradic 4041000! (equal to 298210) pick out the digits in (4,0,6,2,1,3,5), the 2982nd permutation of the numbers 0 through 6.
	//
	//                                  4041000! → (4,0,6,2,1,3,5)
	// factoradic:  4          0                        4        1          0          0        0!
	//              |          |                        |        |          |          |        |
	//     (0,1,2,3,4,5,6) -> (0,1,2,3,5,6) -> (1,2,3,5,6) -> (1,2,3,5) -> (1,3,5) -> (3,5) -> (5)
	//              |          |                        |        |          |          |        |
	// permutation:(4,         0,                       6,       2,         1,         3,       5)
	//
	//
	// One property is that the last index is always 0 (there is no choice left)
	//
	// If I want to compute the next one (the goal here), I need to compute the next factoradic number by applying two simple rules:
	//
	// if the current digit is the last of its position (0! 1! 2! 3! etc), then I set it to 0, and increment the next digit.
	//
	// else simply increment it
	//
	// In the example it would be :
	//
	// digit0 = 0 last possible one, move to next
	//
	// digit1 = 3 it's the index 0 of (3,5), so increment the index (1) and apply digit1 := 5
	//
	//
	// We can find out that a digit can be incremented iif it is *not* the max of the previous digits values.
	//
	// In the example previous digits values are 5,3 , and 3 is not the max, it can be simply incremented

	// the first number that is not the "max" of the previous values, is the one that can be "incremented"

	i := len(p) - 2 // loop from last but one digit
	// the last one is always 0, and can never be incremented
	for i >= 0 && p[i] > p[i+1] {
		i--
	}
	// this test seems odd, but it works by recursion:
	// recursive hypothethis: p[i+1] is the max of the remaining digits
	// therefore  p[i]> p[i+1] =>  p[i] is also the max of the remaining digits
	//

	// Therefore now 'i' is either -1 meaning that we have reached the end, the next is the first in fact
	if i < 0 {
		sort.Ints(p) //we start again
		return false
	}

	// we keep the actual index at i
	val := p[i]     // this is the one that will be "incremented" ( next value in possible ones)
	vals := p[i+1:] // subslice of remaining values (one is greater that val, rembember!)
	sort.Ints(vals) // sort them, 1- to search among them, but also because this is how they will be added

	pos := sort.SearchInts(vals, val) // find where I would insert val (this is the next one)
	p[i], vals[pos] = vals[pos], val  // permute val, with the next one (found above)

	// and now copy vals to the rest of p
	copy(p[i+1:], vals)
	return true
}

// SubsetLexNext updates 'p' to be the next lexicographical combination in the list of all combinations in lexicographical order
//
// returns false if we have generated all the elements
func SubsetLexNext(p []int, n int) bool {
	k := len(p)
	i := k - 1
	for ; i >= 0 && p[i] == n-k+i; i-- {
	}

	if i < 0 {
		for j := 0; j < len(p); j++ {
			p[j] = j //identity
		}
		return false
	}
	pi := p[i]
	for j := i; j < k; j++ {
		p[j] = pi + 1 + j - i
	}
	return true
}
