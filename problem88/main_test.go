package problem88

import "testing"

func TestMerge(t *testing.T) {
	tests := []struct {
		nums1, nums2 []int
		m, n         int
		want         []int
	}{
		{
			nums1: []int{1, 2, 3, 0, 0, 0},
			nums2: []int{2, 5, 6},
			m:     3,
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1: []int{1},
			nums2: []int{},
			m:     1,
			n:     0,
			want:  []int{1},
		},
		{
			nums1: []int{0},
			nums2: []int{1},
			m:     0,
			n:     1,
			want:  []int{1},
		},
	}

	for _, test := range tests {
		nums1 := test.nums1
		merge(test.nums1, test.m, test.nums2, test.n)
		if len(test.nums1) != len(test.want) {
			t.Errorf("Test failed.")
			break
		}

		for i := 0; i < len(test.nums1); i++ {
			if test.nums1[i] != test.want[i] {
				t.Errorf("Test failed. The nums1=%v, m=%d, nums2=%v, n=%d, is expected to %v but nums1=%v", nums1, test.m, test.nums2, test.n, test.want, test.nums1)
			}
		}
	}
}
