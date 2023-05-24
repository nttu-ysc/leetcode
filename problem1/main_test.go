package problem1

import (
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		excepted []int
	}{
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			excepted: []int{0, 1},
		},
		{
			nums:     []int{3, 2, 4},
			target:   6,
			excepted: []int{1, 2},
		},
		{
			nums:     []int{3, 3},
			target:   6,
			excepted: []int{0, 1},
		},
	}

	for _, test := range tests {
		res := twoSum(test.nums, test.target)
		sort.Slice(res, func(i, j int) bool {
			return res[i] < res[j]
		})

		for i := 0; i < len(res); i++ {
			if res[i] != test.excepted[i] {
				t.Errorf("Test failed. The input nums=%v, target=%d is excepted to %v, but response %v", test.nums, test.target, test.excepted, res)
				goto e
			}
		}
	e:
	}
}
