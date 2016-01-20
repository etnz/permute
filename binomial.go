package permute

// Binomial computes the binomial coefficient: from n choose k
//
// It's an efficient method that avoid useless overflow.
func Binomial(n, k int) int64 {
	var b int64 = 1
	if n-k < k {
		return Binomial(n, n-k)
	}
	for i := 1; i <= k; i++ {
		b *= int64((n - k + i))
		b /= int64(i)

	}
	return b
}
