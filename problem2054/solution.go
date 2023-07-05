package problem2054

import (
	"sort"
)

func maxTwoEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})
	dp := make([]int, len(events))
	dp[0] = events[0][2]
	for i := 1; i < len(events); i++ {
		dp[i] = max(dp[i-1], events[i][2])
	}
	ans := dp[len(dp)-1]
	for i := 1; i < len(events); i++ {
		l, r := 0, i-1
		for l < r {
			mid := (l + r) >> 1
			if events[i][0] > events[mid+1][1] {
				l = mid + 1
			} else {
				r = mid
			}
		}
		if events[i][0] > events[l][1] {
			ans = max(ans, events[i][2]+dp[l])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
