package problem4

import "testing"

func TestFindMidianSortedArrays(t *testing.T) {
	testCases := []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2,
		},
		{
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.5,
		},
	}

	for _, tC := range testCases {
		if got := findMidianSortedArrays(tC.nums1, tC.nums2); got != tC.want {
			t.Errorf("Test failed. The input nums1 = %v, nums2 = %v is expected to %f but return %f", tC.nums1, tC.nums2, tC.want, got)
		}
	}
}
