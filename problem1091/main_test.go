package problem1091

import "testing"

func TestShortestPathBinaryMatrix(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{
				{0, 1},
				{1, 0},
			},
			want: 2,
		},
		{
			grid: [][]int{
				{0, 0, 0},
				{1, 1, 0},
				{1, 1, 0},
			},
			want: 4,
		},
		{
			grid: [][]int{
				{1, 0, 0},
				{1, 1, 0},
				{1, 1, 0},
			},
			want: -1,
		},
	}

	for _, test := range tests {
		if got := shortestPathBinaryMatrix(test.grid); got != test.want {
			t.Errorf("Test failed. the input grid=%v is excepted to %d, but get %d", test.grid, test.want, got)
		}
	}
}
