package problem2049

func countHighestScoreNodes(parents []int) int {
	n := len(parents)
	children := make([][]int, n)
	for k := range children {
		children[k] = make([]int, 0)
	}
	for i := 1; i < n; i++ {
		children[parents[i]] = append(children[parents[i]], i)
	}
	var ans int
	var counts = make([][2]int, n)
	count(0, children, &counts)
	scores := make(map[int]int)
	for i := 0; i < n; i++ {
		l, r, remain := counts[i][0], counts[i][1], (n - 1 - counts[i][0] - counts[i][1])
		if l == 0 {
			l = 1
		}
		if r == 0 {
			r = 1
		}
		if remain == 0 {
			remain = 1
		}
		scores[l*r*remain]++
		ans = max(ans, l*r*remain)
	}
	return scores[ans]
}

func count(node int, children [][]int, counts *[][2]int) {
	if len(children[node]) == 2 {
		count(children[node][0], children, counts)
		count(children[node][1], children, counts)
		l := (*counts)[children[node][0]]
		r := (*counts)[children[node][1]]
		left := l[0] + l[1] + 1
		right := r[0] + r[1] + 1
		(*counts)[node] = [2]int{left, right}
	} else if len(children[node]) == 1 {
		count(children[node][0], children, counts)
		l := (*counts)[children[node][0]]
		left := l[0] + l[1] + 1
		(*counts)[node] = [2]int{left}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
