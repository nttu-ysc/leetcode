package problem217

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{1, 2, 3, 1},
			want: true,
		},
		{
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			want: true,
		},
	}

	for _, test := range tests {
		if got := containsDuplicate(test.nums); got != test.want {
			t.Errorf("Test failed. The input nums=%v is excepted to %t but return %t", test.nums, test.want, got)
		}
	}
}
