package problem209

import "testing"

func TestMinSubArray(t *testing.T) {
	testCases := []struct {
		target int
		nums   []int
		want   int
	}{
		// {
		// 	target: 213,
		// 	nums:   []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12},
		// 	want:   8,
		// },
		{
			target: 7,
			nums:   []int{2, 3, 1, 2, 4, 3},
			want:   2,
		},
		{
			target: 4,
			nums:   []int{1, 4, 4},
			want:   1,
		},
		{
			target: 11,
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			want:   0,
		},
	}

	for _, tC := range testCases {
		if got := minSubArrayLen(tC.target, tC.nums); got != tC.want {
			t.Errorf("Test failed. The input target = %d, nums = %v is expected to %d but return %d", tC.target, tC.nums, tC.want, got)
		}
	}
}
