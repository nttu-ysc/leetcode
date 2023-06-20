package problem26

func removeDuplicates(nums []int) int {
	var prev = nums[0]
	for i := 1; i < len(nums); {
		if nums[i] == prev {
			nums = append(nums[:i], nums[i+1:]...)
			continue
		} else {
			prev = nums[i]
			i++
		}
	}
	return len(nums)
}

func removeDuplicates2(nums []int) int {
	tmp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		tmp[nums[i]]++
	}
	return len(tmp)
}
