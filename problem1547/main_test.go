package problem1547

import "testing"

func TestMinCost(t *testing.T) {
	tests := []struct {
		n    int
		cuts []int
		want int
	}{
		{
			n:    7,
			cuts: []int{1, 3, 4, 5},
			want: 16,
		},
		{
			n:    9,
			cuts: []int{5, 6, 1, 4, 2},
			want: 22,
		},
	}

	for _, test := range tests {
		if got := minCost(test.n, test.cuts); got != test.want {
			t.Errorf("Test failed. The input n=%d, cuts=%v is expected to %d, but return %d", test.n, test.cuts, test.want, got)
		}
	}
}
