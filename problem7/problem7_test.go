package problem7

import (
	"testing"
)

type pairTest struct {
	value    int
	expected int
}

var tests = []pairTest{
	{6, 13},
	{7, 17},
	{8, 19},
	{9, 23},
	{10, 29},
	{11, 31},
	{10001, 104743},
}

func TestPrimeNth(t *testing.T) {
	for _, test := range tests {
		prime := PrimeNth(test.value)
		if prime != test.expected {
			t.Errorf("PrimeNth(%d) = %d, WANT: %d", test.value, prime, test.expected)
		}
	}
}

func TestPrime(t *testing.T) {
	testMap := map[int]bool{
		1:  false,
		2:  true,
		3:  true,
		12: false,
		13: true,
		22: false,
		23: true,
	}

	for k, v := range testMap {
		res := prime(k)
		if res != v {
			t.Errorf("prime(%d) = %v, WANT %v", k, res, v)
		}
	}
}
