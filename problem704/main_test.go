package problem704

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			want:   4,
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			want:   -1,
		},
	}

	for _, test := range tests {
		if got := search(test.nums, test.target); got != test.want {
			t.Errorf("Test failed. The input nums = %v, target = %d is expected to %d but return %d", test.nums, test.target, test.want, got)
		}
	}
}
