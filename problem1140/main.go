package problem1140

import "math"

var n int
var suffix_sum []int
var dp [200][200]int

func stoneGameII(piles []int) int {
	n = len(piles)
	dp = [200][200]int{}
	suffix_sum = make([]int, n)
	suffix_sum[n-1] = piles[n-1]
	for i := n - 2; i >= 0; i-- {
		suffix_sum[i] = piles[i] + suffix_sum[i+1]
	}

	ans := dfs(0, 1)
	return ans
}

func dfs(i, m int) int {
	if i+2*m >= n {
		dp[i][m] = suffix_sum[i]
		return dp[i][m]
	}
	tmp := math.MaxInt
	for x := 1; x <= 2*m; x++ {
		b := max(x, m)
		if dp[i+x][b] != 0 {
			tmp = min(dp[i+x][b], tmp)
		} else {
			tmp = min(dfs(i+x, b), tmp)
		}
	}

	dp[i][m] = suffix_sum[i] - tmp
	return dp[i][m]
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
