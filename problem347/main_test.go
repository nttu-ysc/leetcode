package problem347

import "testing"

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2},
		},
		{
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
	}

	for _, test := range tests {
		res := topKFrequent(test.nums, test.k)
		if len(test.expected) != len(res) {
			t.Fatalf("Test failed. The input nums=%v, k=%d, is expected to %v, but return %v", test.nums, test.k, test.expected, res)
		}
		for i := 0; i < len(test.expected); i++ {
			if test.expected[i] != res[i] {
				t.Errorf("Test failed. The input nums=%v, k=%d, is expected to %v, but return %v", test.nums, test.k, test.expected, res)
			}
		}
	}
}
