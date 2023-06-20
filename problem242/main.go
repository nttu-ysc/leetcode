package problem242

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hashMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		hashMap[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		hashMap[t[i]]--
		if hashMap[t[i]] < 0 {
			return false
		}
	}
	return true
}
