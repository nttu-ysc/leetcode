package problem33

import "testing"

func TestSearch(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		{
			nums:   []int{1},
			target: 0,
			want:   -1,
		},
		{
			nums:   []int{1, 3},
			target: 2,
			want:   -1,
		},
		{
			nums:   []int{3, 1},
			target: 1,
			want:   1,
		},
	}

	for _, tC := range testCases {
		if got := search(tC.nums, tC.target); got != tC.want {
			t.Errorf("Test failed. The input nums = %v, target = %d is expected to %d but return %d", tC.nums, tC.target, tC.want, got)
		}
	}
}
