package problem1232

import (
	"testing"
)

func TestCheckStraightLine(t *testing.T) {
	tests := []struct {
		coordinates [][]int
		want        bool
	}{
		{
			coordinates: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{4, 5},
				{5, 6},
				{6, 7},
			},
			want: true,
		},
		{
			coordinates: [][]int{
				{1, 1},
				{2, 2},
				{3, 4},
				{4, 5},
				{5, 6},
				{7, 7},
			},
			want: false,
		},
		{
			coordinates: [][]int{
				{0, 0},
				{0, 1},
				{0, -1},
			},
			want: true,
		},
	}

	for _, test := range tests {
		if got := checkStraightLine(test.coordinates); got != test.want {
			t.Errorf("Test failed. The input coordinates = %v is excepted to %t, but return %t", test.coordinates, test.want, got)
		}
	}
}
