package problem14

func longestCommonPrefix(strs []string) string {
	var str string = strs[0]
	for i := 1; i < len(strs); i++ {
		if str == "" {
			return str
		}
		var tmp string
		for k := 0; k < len(strs[i]) && k < len(str); k++ {
			if strs[i][k] != str[k] {
				str = tmp
				break
			}
			tmp += string(str[k])
		}
		if len(tmp) < len(str) {
			str = tmp
		}
	}

	return str
}
