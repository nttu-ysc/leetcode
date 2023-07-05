package problem2054

import (
	"testing"
)

func TestMaxTwoEvents(t *testing.T) {
	testCases := []struct {
		events [][]int
		want   int
	}{
		{
			events: [][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}},
			want:   4,
		},
		{
			events: [][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}},
			want:   5,
		},
		{
			events: [][]int{{1, 5, 3}, {1, 5, 1}, {6, 6, 5}},
			want:   8,
		},
		{
			events: [][]int{{10, 83, 53}, {63, 87, 45}, {97, 100, 32}, {51, 61, 16}},
			want:   85,
		},
	}

	for _, tC := range testCases {
		if got := maxTwoEvents(tC.events); got != tC.want {
			t.Errorf("Test failed. The input events = %v is expected to %d but return %d", tC.events, tC.want, got)
		}
	}
}
