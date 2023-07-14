package problem1218

func longestSubsequence(arr []int, difference int) int {
	var ans int
	dp := make(map[int]int)
	for _, i := range arr {
		dp[i] = 1 + dp[i-difference]
		ans = max(ans, dp[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
