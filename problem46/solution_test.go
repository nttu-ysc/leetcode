package problem46

import (
	"testing"
)

func TestPermute(t *testing.T) {
	testCases := []struct {
		nums []int
		want [][]int
	}{
		{
			nums: []int{1, 2, 3},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			nums: []int{0, 1},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			nums: []int{1},
			want: [][]int{{1}},
		},
	}

	for _, tC := range testCases {
		got := permute(tC.nums)

		if len(got) != len(tC.want) {
			t.Errorf("Test failed.")
			continue
		}

		for k := range got {
			if len(got[k]) != len(tC.want[k]) {
				t.Errorf("Test failed.")
				break
			}

			for kk := range got[k] {
				if got[k][kk] != tC.want[k][kk] {
					t.Errorf("Test failed. The input nums = %v is expected to %v but return %v", tC.nums, tC.want, got)
				}
				break
			}
		}
	}
}
