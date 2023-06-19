package problem837

import "testing"

func TestNew21Game(t *testing.T) {
	tests := []struct {
		n        int
		k        int
		maxPts   int
		expected float64
	}{
		{
			n:        10,
			k:        1,
			maxPts:   10,
			expected: 1.0,
		},
		{
			n:        6,
			k:        1,
			maxPts:   10,
			expected: 0.6,
		},
		{
			n:        21,
			k:        17,
			maxPts:   10,
			expected: 0.73278,
		},
		{
			n:        98,
			k:        33,
			maxPts:   68,
			expected: 0.99898,
		},
	}

	for _, test := range tests {
		if res := new21Game(test.n, test.k, test.maxPts); res != test.expected {
			t.Errorf("Test failed. The input n=%d, k=%d, maxPts=%d, expected %f  but return %f", test.n, test.k, test.maxPts, test.expected, res)
		}
	}
}
