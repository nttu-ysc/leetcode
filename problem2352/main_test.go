package problem2352

import "testing"

func TestEqualPairs(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{
				{3, 2, 1},
				{2, 7, 6},
				{2, 7, 7},
			},
			want: 1,
		},
		{
			grid: [][]int{
				{3, 1, 2, 2},
				{1, 4, 4, 5},
				{2, 4, 2, 2},
				{2, 4, 2, 2},
			},
			want: 3,
		},
	}

	for _, test := range tests {
		if got := equalPairs(test.grid); got != test.want {
			t.Errorf("Test failed. The grid = %v is expected to %d but return %d", test.grid, test.want, got)
		}
	}
}
