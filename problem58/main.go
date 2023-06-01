package problem58

func lengthOfLastWord(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			var cnt int
			for j := i; j >= 0 && s[j] != ' '; j-- {
				cnt++
			}
			return cnt
		}
	}

	return 0
}
