package problem1863

import "testing"

func TestSubsetXORSumn(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 3},
			want: 6,
		},
		{
			nums: []int{5, 1, 6},
			want: 28,
		},
		{
			nums: []int{3, 4, 5, 6, 7, 8},
			want: 480,
		},
	}

	for _, tC := range testCases {
		if got := subsetXORSum(tC.nums); got != tC.want {
			t.Errorf("Test failed. The input nums = %v is expected to %d but return %d", tC.nums, tC.want, got)
		}
	}
}
