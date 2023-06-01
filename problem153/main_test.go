package problem153

import "testing"

func TestFineMin(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{3, 4, 5, 1, 2},
			want: 1,
		},
		{
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			want: 0,
		},
		{
			nums: []int{11, 13, 15, 17},
			want: 11,
		},
	}

	for _, test := range tests {
		if got := findMin(test.nums); got != test.want {
			t.Errorf("Test failed. the input nums=%v is excepted to %d but return %d", test.nums, test.want, got)
		}
	}
}
