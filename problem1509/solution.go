package problem1509

import (
	"sort"
)

func minDifference(nums []int) int {
	if len(nums) <= 4 {
		return 0
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	var ans = 1<<32 - 1
	for l, r := 0, len(nums)-4; r < len(nums); l, r = l+1, r+1 {
		ans = min(ans, nums[l]-nums[r])
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
