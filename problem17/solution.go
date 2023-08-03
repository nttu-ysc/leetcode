package problem17

var tele = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	var res []string
	for i := range digits {
		if len(res) == 0 {
			res = append(res, tele[digits[i]]...)
			continue
		}
		var tmp []string
		for _, r := range res {
			for _, el := range tele[digits[i]] {
				tmp = append(tmp, r+el)
			}
		}
		res = tmp
	}
	return res
}
