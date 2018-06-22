package problem3

import (
	"testing"
)

type testPair struct {
	value    uint64
	expected uint64
}

var tests = []testPair{
	{600851475143, 6857},
}

func TestPrimeFactor(t *testing.T) {
	for _, test := range tests {
		max := PrimeMax(test.value)
		if max != test.expected {
			t.Errorf("PrimeMax(%d) = %d, WANT: %d", test.value, max, test.expected)
		}
	}
}
