package problem837

import "testing"

func TestNew21Game(t *testing.T) {
	tests := []struct {
		n        int
		k        int
		maxPts   int
		excepted float64
	}{
		{
			n:        10,
			k:        1,
			maxPts:   10,
			excepted: 1.0,
		},
		{
			n:        6,
			k:        1,
			maxPts:   10,
			excepted: 0.6,
		},
		{
			n:        21,
			k:        17,
			maxPts:   10,
			excepted: 0.73278,
		},
		{
			n:        98,
			k:        33,
			maxPts:   68,
			excepted: 0.99898,
		},
	}

	for _, test := range tests {
		if res := new21Game(test.n, test.k, test.maxPts); res != test.excepted {
			t.Errorf("Test failed. The input n=%d, k=%d, maxPts=%d, excepted %f  but return %f", test.n, test.k, test.maxPts, test.excepted, res)
		}
	}
}
