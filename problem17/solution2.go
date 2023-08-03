package problem17

func letterCombinations2(digits string) []string {
	var ans []string
	if len(digits) == 0 {
		return ans
	}
	helper(0, digits, &ans, "")
	return ans
}

func helper(cur int, digits string, ans *[]string, tmp string) {
	if len(tmp) == len(digits) {
		*ans = append(*ans, tmp)
		return
	}
	for _, el := range tele[digits[cur]] {
		helper(cur+1, digits, ans, tmp+el)
	}
}
