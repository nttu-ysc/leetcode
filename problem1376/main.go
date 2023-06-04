package problem1376

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	var dp []int = make([]int, n)
	var t int
	for i := 0; i < n; i++ {
		var tmp int
		dfs(i, manager, informTime, &tmp, &dp)
		dp[i] = tmp
		t = max(t, tmp)
	}

	return t
}

func dfs(child int, manager []int, informTime []int, time *int, dp *[]int) {
	if manager[child] == -1 {
		return
	}

	*time += informTime[manager[child]]
	if (*dp)[manager[child]] != 0 {
		*time += (*dp)[manager[child]]
	} else {
		dfs(manager[child], manager, informTime, time, dp)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
