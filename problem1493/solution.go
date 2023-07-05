package problem1493

func longestSubarray(nums []int) int {
	var zeroCnt int
	var ans int
	var l, r int
	for ; r < len(nums); r++ {
		if nums[r] == 0 {
			zeroCnt++
		}
		for zeroCnt > 1 {
			ans = max(ans, r-l-1)
			if nums[l] == 0 {
				zeroCnt--
			}
			l++
		}
	}
	return max(ans, r-l-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
