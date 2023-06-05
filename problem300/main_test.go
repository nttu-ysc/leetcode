package problem300

import "testing"

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			want: 4,
		},
		{
			nums: []int{0, 1, 0, 3, 2, 3},
			want: 4,
		},
		{
			nums: []int{7, 7, 7, 7, 7, 7, 7},
			want: 1,
		},
	}

	for _, test := range tests {
		if got := lengthOfLIS(test.nums); got != test.want {
			t.Errorf("Test failed. The input nums = %v is excepted to %d, but return %d", test.nums, test.want, got)
		}
	}
}
