package problem2542

import "testing"

func TestMaxScore(t *testing.T) {
	tests := []struct {
		nums1    []int
		nums2    []int
		k        int
		excepted int64
	}{
		{
			nums1:    []int{1, 3, 3, 2},
			nums2:    []int{2, 1, 3, 4},
			k:        3,
			excepted: 12,
		},
		{
			nums1:    []int{4, 2, 3, 1, 1},
			nums2:    []int{7, 5, 10, 9, 6},
			k:        1,
			excepted: 30,
		},
	}

	for _, test := range tests {
		if res := maxScore(test.nums1, test.nums2, test.k); res != test.excepted {
			t.Errorf("Test failed.")
			// t.Errorf("Test failed. The input nums1=%d, nums2=")
		}
	}
}
