package problem2616

import (
	"testing"
)

func TestMinimizeMax(t *testing.T) {
	testCases := []struct {
		nums []int
		p    int
		want int
	}{
		{
			nums: []int{10, 1, 2, 7, 1, 3},
			p:    2,
			want: 1,
		},
		{
			nums: []int{4, 2, 1, 2},
			p:    1,
			want: 0,
		},
	}

	for _, tC := range testCases {
		if got := minimizeMax(tC.nums, tC.p); got != tC.want {
			t.Errorf("Test failed. The input nums = %v, p = %d is expected to %d but return %d", tC.nums, tC.p, tC.want, got)
		}
	}
}
