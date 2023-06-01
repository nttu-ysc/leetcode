package problem81

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   bool
	}{
		{
			nums:   []int{2, 5, 6, 0, 0, 1, 2},
			target: 0,
			want:   true,
		},
		{
			nums:   []int{2, 5, 6, 0, 0, 1, 2},
			target: 3,
			want:   false,
		},
		{
			nums:   []int{1, 0, 1, 1, 1},
			target: 0,
			want:   true,
		},
		{
			nums:   []int{1},
			target: 1,
			want:   true,
		},
	}

	for _, test := range tests {
		if got := search(test.nums, test.target); got != test.want {
			t.Errorf("Test failed. The input nums=%v, target=%d  is excepted to %t but return %t", test.nums, test.target, test.want, got)
		}
	}
}
