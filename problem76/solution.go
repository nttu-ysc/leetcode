package problem76

func minWindow(s, t string) string {
	var ans string
	letterCnt := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		letterCnt[t[i]]++
	}
	var left, cnt int
	var minLen = 1<<63 - 1
	for i := 0; i < len(s); i++ {
		if letterCnt[s[i]]--; letterCnt[s[i]] >= 0 {
			cnt++
		}
		for cnt == len(t) {
			if minLen > i-left+1 {
				minLen = i - left + 1
				ans = s[left : left+minLen]
			}
			if letterCnt[s[left]]++; letterCnt[s[left]] > 0 {
				cnt--
			}
			left++
		}
	}
	return ans
}
