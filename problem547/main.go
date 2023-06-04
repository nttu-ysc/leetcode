package problem547

func findCircleNum(isConnected [][]int) int {
	var isVisited []bool = make([]bool, len(isConnected))
	var cnt int

	for i := 0; i < len(isConnected); i++ {
		if !isVisited[i] {
			cnt++
			dfs(i, isConnected, &isVisited)
		}
	}

	return cnt
}

func dfs(node int, isConnected [][]int, isVisited *[]bool) {
	(*isVisited)[node] = true
	for next := range isConnected[node] {
		if isConnected[node][next] == 0 {
			continue
		}

		if !(*isVisited)[next] {
			dfs(next, isConnected, isVisited)
		}
	}
}
