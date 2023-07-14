package problem1218

import "testing"

func TestLongestSubsequence(t *testing.T) {
	testCases := []struct {
		arr        []int
		difference int
		want       int
	}{
		{
			arr:        []int{1, 2, 3, 4},
			difference: 1,
			want:       4,
		},
		{
			arr:        []int{1, 3, 5, 7},
			difference: 1,
			want:       1,
		},
		{
			arr:        []int{1, 5, 7, 8, 5, 3, 4, 2, 1},
			difference: -2,
			want:       4,
		},
	}

	for _, tC := range testCases {
		if got := longestSubsequence(tC.arr, tC.difference); got != tC.want {
			t.Errorf("Test failed. The arr = %v, difference = %d is expected to %d but return %d", tC.arr, tC.difference, tC.want, got)
		}
	}
}
