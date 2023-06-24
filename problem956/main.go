package problem956

func tallestBillboard(rods []int) int {
	var n = len(rods)
	var sum int
	dp := make([][]int, n+1)
	for _, v := range rods {
		sum += v
	}
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, sum+1)
		for k := range dp[i] {
			dp[i][k] = -1
		}
	}
	dp[0][0] = 0

	for i := 1; i < n+1; i++ {
		h := rods[i-1]
		for j := 0; j < sum+1-h; j++ {
			if dp[i-1][j] < 0 {
				continue
			}
			// not used
			dp[i][j] = max(dp[i][j], dp[i-1][j])
			// put on the taller one
			dp[i][j+h] = max(dp[i][j+h], dp[i-1][j])
			// put on the shorter one
			dp[i][abs(j-h)] = max(dp[i][abs(j-h)], dp[i-1][j]+min(h, j))
		}
	}
	return dp[n][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
