package problem200

import "testing"

func TestNUmIslands(t *testing.T) {
	tests := []struct {
		grid [][]byte
		want int
	}{
		{
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		{
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
	}

	for _, test := range tests {
		if got := numIslands(test.grid); got != test.want {
			t.Errorf("Test failed. The input grid=%v is expected to %d but return %d", test.grid, test.want, got)
		}
	}
}
