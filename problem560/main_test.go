package problem560

import "testing"

func TestSubarraySum(t *testing.T) {
	testCases := []struct {
		nums []int
		k    int
		want int
	}{
		{
			nums: []int{1, 1, 1},
			k:    2,
			want: 2,
		},
		{
			nums: []int{1, 2, 3},
			k:    3,
			want: 2,
		},
	}
	for _, tC := range testCases {
		if got := subarraySum(tC.nums, tC.k); got != tC.want {
			t.Errorf("Test failed. The input nums = %v, k = %d is expected to %d but return %d", tC.nums, tC.k, tC.want, got)
		}
	}
}
