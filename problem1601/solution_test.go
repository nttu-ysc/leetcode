package problem1601

import "testing"

func TestMaximumRequests(t *testing.T) {
	testCases := []struct {
		n        int
		requests [][]int
		want     int
	}{
		{
			n:        6,
			requests: [][]int{{0, 1}, {1, 0}, {0, 1}, {1, 2}, {2, 0}, {3, 4}, {5, 2}, {2, 3}},
			want:     5,
		},
		{
			n:        5,
			requests: [][]int{{0, 1}, {1, 0}, {0, 1}, {1, 2}, {2, 0}, {3, 4}},
			want:     5,
		},
		{
			n:        3,
			requests: [][]int{{0, 0}, {1, 2}, {2, 1}},
			want:     3,
		},
		{
			n:        4,
			requests: [][]int{{0, 3}, {3, 1}, {1, 2}, {2, 0}},
			want:     4,
		},
	}

	for _, tC := range testCases {
		if got := maximumRequests(tC.n, tC.requests); got != tC.want {
			t.Errorf("Test failed. The n = %d, requests = %v is expected to %d but return %d", tC.n, tC.requests, tC.want, got)
		}
	}
}
