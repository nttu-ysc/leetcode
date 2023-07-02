package problem1863

func subsetXORSum(nums []int) int {
	return backtrack(nums, 0, 0)
}

func backtrack(nums []int, index, num int) int {
	if index == len(nums) {
		return num
	}
	var ans int
	ans += backtrack(nums, index+1, num)
	ans += backtrack(nums, index+1, num^nums[index])
	return ans
}
