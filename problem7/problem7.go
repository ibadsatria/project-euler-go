package problem7

import (
	"math/big"
)

// prime 1001st
// prime 6th is 13

func prime(n int) bool {
	if big.NewInt(int64(n)).ProbablyPrime(0) {
		return true
	}
	return false
}

// PrimeNth returns prime nth
func PrimeNth(nth int) int {
	var primeArr = []int{2, 3, 5, 7}
	num := 11
	i := 4

	for {
		if prime(num) {
			primeArr = append(primeArr, num)
			i++
		}

		if i == nth {
			length := len(primeArr)
			return primeArr[length-1]
		}
		num++
	}
}
