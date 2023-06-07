package problem1558

import "testing"

func TestMinOperations(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 5},
			want: 5,
		},
		{
			nums: []int{2, 2},
			want: 3,
		},
		{
			nums: []int{4, 2, 5},
			want: 6,
		},
	}

	for _, test := range tests {
		if got := minOperations(test.nums); got != test.want {
			t.Errorf("Test failed. The input nums=%v is excepted to %d but return %d", test.nums, test.want, got)
		}
	}

}
