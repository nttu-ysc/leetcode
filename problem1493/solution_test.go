package problem1493

import "testing"

func TestLongestSubarray(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 1, 0, 1},
			want: 3,
		},
		{
			nums: []int{0, 1, 1, 1, 0, 1, 1, 0, 1},
			want: 5,
		},
		{
			nums: []int{1, 1, 1},
			want: 2,
		},
	}

	for _, tC := range testCases {
		if got := longestSubarray(tC.nums); got != tC.want {
			t.Errorf("Test failed. The input nums = %v is expected to %d but return %d", tC.nums, tC.want, got)
		}
	}
}
