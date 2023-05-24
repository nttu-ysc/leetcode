package problem703

import "testing"

func TestKthLargest(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		add  []struct {
			input    int
			excepted int
		}
	}{
		{
			nums: []int{4, 5, 8, 2},
			k:    3,
			add: []struct {
				input    int
				excepted int
			}{
				{
					input:    3,
					excepted: 4,
				},
				{
					input:    5,
					excepted: 5,
				},
				{
					input:    10,
					excepted: 5,
				},
				{
					input:    9,
					excepted: 8,
				},
				{
					input:    4,
					excepted: 8,
				},
			},
		},
	}

	for _, test := range tests {
		kl := Constructor(test.k, test.nums)
		for _, v := range test.add {
			if res := kl.Add(v.input); res != v.excepted {
				t.Errorf("Test failed. The input nums=%v, k=%d, add(%d) is excepted to %d but response %d", test.nums, test.k, v.input, v.excepted, res)
			}
		}
	}
}
