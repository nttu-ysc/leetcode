package problem77

func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	helper(1, n, k, &ans, make([]int, 0))
	return ans
}

func helper(cur int, n int, k int, ans *[][]int, sel []int) {
	if len(sel) == k {
		*ans = append(*ans, append([]int{}, sel...))
		return
	}
	if cur > n {
		return
	}
	helper(cur+1, n, k, ans, sel)
	sel = append(sel, cur)
	helper(cur+1, n, k, ans, sel)
}
