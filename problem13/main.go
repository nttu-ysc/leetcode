package problem13

func romanToInt(s string) int {
	symbols := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var ans, prev int
	for i := len(s) - 1; i >= 0; i-- {
		tmp := symbols[s[i]]
		if tmp >= prev {
			ans += tmp
		} else {
			ans -= tmp
		}
		prev = tmp
	}

	return ans
}
