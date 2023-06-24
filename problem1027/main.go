package problem1027

func longestArithSeqLength(nums []int) int {
	dp := make([]map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make(map[int]int)
	}
	ans := 2
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			k := nums[j] - v
			if prev, ok := dp[i][k]; ok {
				dp[j][k] = prev + 1
				if prev+1 > ans {
					ans = prev + 1
				}
			} else {
				dp[j][k] = 2
			}
		}
	}
	return ans
}
