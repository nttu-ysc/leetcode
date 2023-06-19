package problem2733

import "testing"

func TestFindNonMinOrMax(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{3, 2, 1, 4},
			want: 2,
		},
		{
			nums: []int{1, 2},
			want: -1,
		},
		{
			nums: []int{2, 1, 3},
			want: 2,
		},
	}

	for _, test := range tests {
		if got := findNonMinOrMax(test.nums); got != test.want {
			t.Errorf("Test failed. The input nums = %v is expected to %d but return %d", test.nums, test.want, got)
		}
	}
}
