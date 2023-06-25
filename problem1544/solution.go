package problem1544

func makeGood(s string) string {
	const distance = 'a' - 'A'
	if len(s) < 2 {
		return s
	}
	for i := 0; i < len(s)-1; i++ {
		if abs(int(s[i+1])-int(s[i])) == distance {
			s = s[:i] + s[i+2:]
			return makeGood(s)
		}
	}
	return s
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
