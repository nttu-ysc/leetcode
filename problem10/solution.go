package problem10

func isMatch(s string, p string) bool {
	if len(p) < 1 {
		return len(s) < 1
	}
	if len(p) == 1 {
		return len(s) == 1 && (s[0] == p[0] || p[0] == '.')
	}
	for len(s) > 0 && s[0] == p[0] || p[0] == '.' {
		if isMatch(s, p[2:]) {
			return true
		}
		s = s[1:]
	}
	return isMatch(s, p[2:])
}
