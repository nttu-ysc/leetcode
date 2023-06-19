package problem75

import "testing"

func TestSortColors(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{2, 0, 2, 1, 1, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			nums: []int{2, 0, 1},
			want: []int{0, 1, 2},
		},
	}

	for _, test := range tests {
		nums := test.nums
		sortColors(test.nums)
		if len(test.nums) != len(nums) {
			t.Errorf("Test failed.")
			break
		}

		for i := 0; i < len(test.nums); i++ {
			if test.nums[i] != test.want[i] {
				t.Errorf("Test failed. The input nums=%v is expected to %v but return %v", test.nums, test.want, nums)
				break
			}
		}
	}
}
