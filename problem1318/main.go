package problem1318

func minFlips(a, b, c int) int {
	var res int

	for a > 0 || b > 0 || c > 0 {
		if c&1 == 1 {
			if a&1 != 1 && b&1 != 1 {
				res++
			}
		} else {
			if a&1 != 0 {
				res++
			}
			if b&1 != 0 {
				res++
			}
		}

		a >>= 1
		b >>= 1
		c >>= 1
	}

	return res
}
