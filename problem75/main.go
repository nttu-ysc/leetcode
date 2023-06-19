package problem75

func sortColors(nums []int) {
	for k := 1; k < len(nums); k++ {
		if nums[k] < nums[k-1] {
			nums[k], nums[k-1] = nums[k-1], nums[k]
			sortColors(nums)
		}
	}
}
