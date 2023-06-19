package problem136

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{2, 2, 1},
			want: 1,
		},
		{
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			nums: []int{1},
			want: 1,
		},
	}

	for _, test := range tests {
		if got := singleNumber(test.nums); got != test.want {
			t.Errorf("Test failed. the input nums=%v is expected to %d but return %d", test.nums, test.want, got)
		}
	}
}
