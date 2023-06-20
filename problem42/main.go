package problem42

func trap(height []int) int {
	var n = len(height)
	var ans int
	l := make([]int, n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			l[i] = height[i]
			continue
		}
		l[i] = max(l[i-1], height[i])
	}
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			r[i] = height[i]
			continue
		}
		r[i] = max(r[i+1], height[i])
	}
	for i := 0; i < n; i++ {
		ans += max(min(l[i], r[i])-height[i], 0)
	}
	return ans
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
