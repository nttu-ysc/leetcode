package problem1824

func minSideJumps(obstacles []int) int {
	var dp [3]int
	dp[0], dp[2] = 1, 1
	for i := 0; i < len(obstacles); i++ {
		obj := obstacles[i]
		pre0, pre1, pre2 := dp[0], dp[1], dp[2]
		for i := range dp {
			dp[i] = 1 << 32
		}

		if obj != 1 {
			dp[0] = pre0
		}
		if obj != 2 {
			dp[1] = pre1
		}
		if obj != 3 {
			dp[2] = pre2
		}

		if obj != 1 {
			dp[0] = min(dp[0], min(dp[1], dp[2])+1)
		}
		if obj != 2 {
			dp[1] = min(dp[1], min(dp[0], dp[2])+1)
		}
		if obj != 3 {
			dp[2] = min(dp[2], min(dp[0], dp[1])+1)
		}
	}
	return min(dp[0], min(dp[1], dp[2]))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
