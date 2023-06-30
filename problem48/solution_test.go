package problem48

import "testing"

func TestRotate(t *testing.T) {
	testCases := []struct {
		matrix, want [][]int
	}{
		{
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			matrix: [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			want:   [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
	}

	for _, tC := range testCases {
		rotate(tC.matrix)
		for i := 0; i < len(tC.want); i++ {
			for j := 0; j < len(tC.want); j++ {
				if tC.matrix[i][j] != tC.want[i][j] {
					t.Errorf("Test failed. Matrix = %v", tC.matrix)
					goto b
				}
			}
		}
	b:
	}
}
