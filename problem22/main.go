package problem22

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	current := make([]byte, n*2)
	rec(&ans, n, 0, 0, current)
	return ans
}

func rec(ans *[]string, n int, left int, right int, current []byte) {
	if left+right == n*2 {
		*ans = append(*ans, string(current))
		return
	}

	if left < n {
		current[left+right] = '('
		rec(ans, n, left+1, right, current)
	}

	if right < left {
		current[left+right] = ')'
		rec(ans, n, left, right+1, current)
	}
}
