package problem27

import "testing"

func TestRemoveElement(t *testing.T) {
	testCases := []struct {
		nums []int
		val  int
		want int
	}{
		{
			nums: []int{3, 2, 2, 3},
			val:  3,
			want: 2,
		},
		{
			nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:  2,
			want: 5,
		},
	}
	for _, tC := range testCases {
		if got := removeElement(tC.nums, tC.val); got != tC.want {
			t.Errorf("Test failed. The input nums = %v, val = %d is expected to %d but return %d", tC.nums, tC.val, tC.want, got)
		}
	}
}
