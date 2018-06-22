package problem3

func max(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

func PrimeMax(n uint64) uint64 {
	factor := n
	for i := uint64(2); i*i <= n; i++ {
		if n%i == 0 {
			factor = max(i, n/i)
			break
		}
	}

	if factor == n {
		return n
	}
	return PrimeMax(factor)
}
