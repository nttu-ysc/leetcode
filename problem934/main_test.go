package main

import "testing"

func TestShortestBridge(t *testing.T) {
	tests := []struct {
		grid     [][]int
		expected int
	}{
		{
			grid: [][]int{
				{0, 1},
				{1, 0},
			},
			expected: 1,
		},
		{
			grid: [][]int{
				{0, 1, 0},
				{0, 0, 0},
				{0, 0, 1},
			},
			expected: 2,
		},
		{
			grid: [][]int{
				{1, 1, 1, 1, 1},
				{1, 0, 0, 0, 1},
				{1, 0, 1, 0, 1},
				{1, 0, 0, 0, 1},
				{1, 1, 1, 1, 1},
			},
			expected: 1,
		},
	}
	for _, test := range tests {
		if res := shortestBridge(test.grid); res != test.expected {
			t.Errorf("Test failed. The input %v is expected to %d, but response %d", test.grid, test.expected, res)
		}
	}
}
