package problem50

func myPow(x float64, n int) float64 {
	var res = helper(x, abs(n))
	if n < 0 {
		return 1 / res
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func helper(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	var res = helper(x, n>>1)
	res *= res
	if n%2 != 0 {
		res *= x
	}
	return res
}
