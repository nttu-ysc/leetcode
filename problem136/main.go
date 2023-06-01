package problem136

func singleNumber(nums []int) int {
	var num int = nums[0]
	for i := 1; i < len(nums); i++ {
		num ^= nums[i]
	}

	return num
}
