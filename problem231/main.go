package problem231

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	var ans int
	for i := 0; ans < n; i++ {
		ans = pow(2, i)
	}
	return ans == n
}

func pow(x, n int) int {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	var ans = 1
	for i := 0; i < n; i++ {
		ans *= x
	}

	return ans
}
