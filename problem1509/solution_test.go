package problem1509

import "testing"

func TestMinDifference(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{5, 3, 2, 4},
			want: 0,
		},
		{
			nums: []int{1, 5, 0, 10, 14},
			want: 1,
		},
		{
			nums: []int{3, 100, 20},
			want: 0,
		},
		{
			nums: []int{6, 6, 0, 1, 1, 4, 6},
			want: 2,
		},
		{
			nums: []int{82, 81, 95, 75, 20},
			want: 1,
		},
		{
			nums: []int{9, 48, 92, 48, 81, 31},
			want: 17,
		},
	}

	for _, tC := range testCases {
		if got := minDifference(tC.nums); got != tC.want {
			t.Errorf("Test failed. The input nums = %v is expected to %d but return %d", tC.nums, tC.want, got)
		}
	}
}
