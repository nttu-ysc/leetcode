package problem1575

func countRoutes(locations []int, start int, finish int, fuel int) int {
	const MOD = 1e9 + 7
	if abs(locations[finish]-locations[start]) > fuel {
		return 0
	}
	dp := make([][]int64, fuel+1)
	for k := range dp {
		dp[k] = make([]int64, len(locations))
	}
	dp[0][start] = 1
	for i := 1; i <= fuel; i++ {
		for j := 0; j < len(locations); j++ {
			if i > fuel-abs(locations[finish]-locations[j]) {
				continue
			}
			for k := 0; k < len(locations); k++ {
				if j == k {
					if k == start {
						dp[i][j] = (dp[i][j] + 1) % MOD
					}
					continue
				}
				var tmp int64
				if cost := abs(locations[j] - locations[k]); i >= cost {
					tmp = dp[i-cost][k] % MOD
				}
				dp[i][j] = (dp[i][j] + tmp) % MOD
			}
		}
	}
	return int(dp[fuel][finish])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
