package problem1732

func largestAltitude(gain []int) int {
	var ans, sum int
	for _, g := range gain {
		sum += g
		ans = max(ans, sum)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
