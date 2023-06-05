package problem1768

func mergeAlternately(word1 string, word2 string) string {
	var merge string
	m := min(len(word1), len(word2))
	for i := 0; i < m; i++ {
		merge += string(word1[i]) + string(word2[i])
	}
	merge += word1[m:] + word2[m:]

	return merge
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
