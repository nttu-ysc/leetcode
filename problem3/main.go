package problem3

func lengthOfLongestSubstring(s string) int {
	hashMap := make(map[byte]bool)

	var ans int
	for l, r := 0, 0; r < len(s); r++ {
		for _, ok := hashMap[s[r]]; ok; _, ok = hashMap[s[r]] {
			if len(hashMap) > ans {
				ans = len(hashMap)
			}
			delete(hashMap, s[l])
			l++
		}
		hashMap[s[r]] = true
	}

	if len(hashMap) > ans {
		ans = len(hashMap)
	}

	return ans
}
