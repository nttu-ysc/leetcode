package problem2521

import "testing"

func TestDistinctPrimeFactors(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{2, 4, 3, 7, 10, 6},
			want: 4,
		},
		{
			nums: []int{2, 4, 8, 16},
			want: 1,
		},
	}

	for _, tC := range testCases {
		if got := distinctPrimeFactors(tC.nums); got != tC.want {
			t.Errorf("Test failed. The input nums = %v is expected to %d but return %d", tC.nums, tC.want, got)
		}
	}
}
