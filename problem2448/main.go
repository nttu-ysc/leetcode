package problem2448

import "sort"

func minCost(nums []int, cost []int) int64 {
	var ans int64 = 1<<63 - 1
	var n = len(nums)
	var preSumCost int64
	left, right := make([]int64, n), make([]int64, n)

	arr := [][]int64{}
	for i := 0; i < n; i++ {
		arr = append(arr, []int64{int64(nums[i]), int64(cost[i])})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})

	preSumCost = arr[0][1]
	for i := 1; i < n; i++ {
		left[i] = left[i-1] + preSumCost*(abs(arr[i][0]-arr[i-1][0]))
		preSumCost += arr[i][1]
	}

	preSumCost = arr[len(arr)-1][1]
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] + preSumCost*int64(abs(arr[i][0]-arr[i+1][0]))
		preSumCost += arr[i][1]
	}

	for i := 0; i < n; i++ {
		ans = min(ans, left[i]+right[i])
	}

	return ans
}

func min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}
