package problem403

// [0,1,3,5,6,8,12,17]

func canCross(stones []int) bool {
	dp := make(map[int]bool)
	res := dfs(stones, 0, 0, &dp)
	return res
}

func dfs(stones []int, pos int, jump int, dp *map[int]bool) bool {
	key := pos | jump<<11
	if pos >= len(stones)-1 {
		return true
	}
	if v, ok := (*dp)[key]; ok {
		return v
	}
	for i := pos + 1; i < len(stones); i++ {
		dist := stones[i] - stones[pos]
		if dist < jump-1 {
			continue
		}
		if dist > jump+1 {
			(*dp)[key] = false
			return (*dp)[key]
		}
		if dfs(stones, i, dist, dp) {
			(*dp)[key] = true
			return (*dp)[key]
		}
	}
	(*dp)[key] = false
	return (*dp)[key]
}
