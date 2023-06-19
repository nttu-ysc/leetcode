package problem1351

import "testing"

func TestCountNegatives(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{
				{4, 3, 2, -1},
				{3, 2, 1, -1},
				{1, 1, -1, -2},
				{-1, -1, -2, -3},
			},
			want: 8,
		},
		{
			grid: [][]int{
				{3, 2},
				{1, 0},
			},
			want: 0,
		},
	}

	for _, test := range tests {
		if got := countNegatives(test.grid); got != test.want {
			t.Errorf("Test failed. The input grid=%v is expected to %d, but return %d", test.grid, test.want, got)
		}
	}

}
