package main

import "testing"

func TestOddCells(t *testing.T) {
	tests := []struct {
		m        int
		n        int
		indices  [][]int
		excepted int
	}{
		{
			m: 2,
			n: 3,
			indices: [][]int{
				{0, 1},
				{1, 1},
			},
			excepted: 6,
		},
		{
			m: 2,
			n: 2,
			indices: [][]int{
				{1, 1},
				{0, 0},
			},
			excepted: 0,
		},
	}

	for _, test := range tests {
		if res := oddCells(test.m, test.n, test.indices); res != test.excepted {
			t.Errorf("Test failed. The input m = %d, n = %d, indices = %v is excepted to %d, but response %d", test.m, test.n, test.indices, test.excepted, res)
		}
	}
}
