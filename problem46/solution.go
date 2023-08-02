package problem46

func permute(nums []int) [][]int {
	var ans [][]int
	helper(nums, &ans, make([]int, 0))
	return ans
}

func helper(nums []int, ans *[][]int, tmp []int) {
	if len(nums) == 0 {
		*ans = append(*ans, append([]int(nil), tmp...))
		return
	}
	for i := 0; i < len(nums); i++ {
		newNums := append([]int(nil), nums[:i]...)
		newNums = append(newNums, nums[i+1:]...)
		newTmp := append([]int(nil), tmp...)
		newTmp = append(newTmp, nums[i])
		helper(newNums, ans, newTmp)
	}
}
