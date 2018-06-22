package problem5

import "testing"

type testPair struct {
	value    int
	expected int
}

var tests = []testPair{
	{5, 60},
	{10, 2520},
	{20, 232792560},
}

func TestMinFactor(t *testing.T) {
	for _, test := range tests {
		res := MinFactor(test.value)
		if res != test.expected {
			t.Errorf("MinFactor(%v) = %d, WANT: %d", test.value, res, test.expected)
		}
	}
}
