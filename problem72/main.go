package problem72

func minDistance(word1 string, word2 string) int {
	wl1 := len(word1)
	wl2 := len(word2)
	dp := make([][]int, wl1+1)
	for k := range dp {
		dp[k] = make([]int, wl2+1)
		dp[k][0] = k
		if k == 0 {
			for k2 := range dp[k] {
				dp[k][k2] = k2
			}
		}
	}

	var cost int
	for i := 1; i <= wl1; i++ {
		for j := 1; j <= wl2; j++ {
			if word1[i-1] != word2[j-1] {
				cost = 1
			} else {
				cost = 0
			}

			dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+cost)
		}
	}

	return dp[wl1][wl2]
}

func min(nums ...int) int {
	var m = nums[0]
	for _, v := range nums {
		if v < m {
			m = v
		}
	}
	return m
}
