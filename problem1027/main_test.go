package problem1027

import "testing"

func TestLongestArithSeqLength(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{3, 6, 9, 12},
			want: 4,
		},
		{
			nums: []int{9, 4, 7, 2, 10},
			want: 3,
		},
		{
			nums: []int{20, 1, 15, 3, 10, 5, 8},
			want: 4,
		},
	}

	for _, tC := range testCases {
		if got := longestArithSeqLength(tC.nums); got != tC.want {
			t.Errorf("Test failed. The input nums = %v is expected to %d but return %d", tC.nums, tC.want, got)
		}
	}
}
