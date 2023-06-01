package problem9

import "strconv"

func isPalindrome(x int) bool {
	y := strconv.Itoa(x)
	l := len(y) / 2
	for i := 0; i < l; i++ {
		if y[i] != y[len(y)-(i+1)] {
			return false
		}
	}

	return true
}
