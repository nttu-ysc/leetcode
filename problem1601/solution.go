package problem1601

func maximumRequests(n int, requests [][]int) int {
	count := make([]int, n)
	return backtrack(requests, 0, count, 0)
}

func backtrack(requests [][]int, index int, count []int, num int) int {
	if index == len(requests) {
		for i := range count {
			if count[i] != 0 {
				return 0
			}
		}
		return num
	}
	var ans int
	count[requests[index][0]]++
	count[requests[index][1]]--
	ans = max(ans, backtrack(requests, index+1, count, num+1))
	count[requests[index][0]]--
	count[requests[index][1]]++
	return max(ans, backtrack(requests, index+1, count, num))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
