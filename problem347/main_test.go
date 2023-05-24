package problem347

import "testing"

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		excepted []int
	}{
		{
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			excepted: []int{1, 2},
		},
		{
			nums:     []int{1},
			k:        1,
			excepted: []int{1},
		},
	}

	for _, test := range tests {
		res := topKFrequent(test.nums, test.k)
		if len(test.excepted) != len(res) {
			t.Fatalf("Test failed. The input nums=%v, k=%d, is excepted to %v, but return %v", test.nums, test.k, test.excepted, res)
		}
		for i := 0; i < len(test.excepted); i++ {
			if test.excepted[i] != res[i] {
				t.Errorf("Test failed. The input nums=%v, k=%d, is excepted to %v, but return %v", test.nums, test.k, test.excepted, res)
			}
		}
	}
}
