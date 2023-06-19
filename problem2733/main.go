package problem2733

func findNonMinOrMax(nums []int) int {
	if len(nums) < 3 {
		return -1
	}
	var max, min int
	if nums[0] > nums[1] {
		max = nums[0]
		min = nums[1]
	} else {
		max = nums[1]
		min = nums[0]
	}

	if nums[2] < min {
		return min
	}
	if nums[2] > max {
		return max
	}
	return nums[2]
}
