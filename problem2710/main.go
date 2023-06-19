package problem2710

func removeTrailingZeros(num string) string {
	var ans string
	for k := len(num) - 1; k >= 0; k-- {
		if num[k] != '0' {
			ans = num[0 : k+1]
			break
		}
	}
	return ans
}

func removeTrailingZeros2(num string) string {
	if num[len(num)-1] == '0' {
		return removeTrailingZeros2(num[:len(num)-1])
	}

	return num
}
