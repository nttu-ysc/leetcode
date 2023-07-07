package problem2024

// "TTFTTFTT"
// TTFTFTT
func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(helper(answerKey, k, 'T'), helper(answerKey, k, 'F'))
}

func helper(answerKey string, k int, ch byte) int {
	var ans, cnt int
	l, r := 0, 0
	for r < len(answerKey) {
		if answerKey[r] == ch {
			cnt++
		}
		for cnt == k+1 {
			ans = max(ans, r-l)
			if answerKey[l] == ch {
				cnt--
			}
			l++
		}
		r++
	}
	return max(ans, r-l)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
