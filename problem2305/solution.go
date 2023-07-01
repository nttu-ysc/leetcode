package problem2305

func distributeCookies(cookies []int, k int) int {
	group := make([]int, k)
	return recursion(cookies, k, 0, group)
}

func recursion(cookies []int, k int, cookie int, group []int) int {
	if cookie == len(cookies) {
		dif := 0
		for i := 0; i < k; i++ {
			dif = max(dif, group[i])
		}
		return dif
	}
	var ans = 1<<63 - 1
	for i := 0; i < k; i++ {
		if group[i]+cookies[cookie] > ans {
			continue
		}
		group[i] += cookies[cookie]
		ans = min(ans, recursion(cookies, k, cookie+1, group))
		group[i] -= cookies[cookie]
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
