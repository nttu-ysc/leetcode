package problem7

import (
	"strconv"
)

func reverse(x int) int {
	strX := []rune(strconv.Itoa(x))
	var negative bool
	if strX[0] == '-' {
		negative = true
		strX = strX[1:]
	}
	for i := 0; i < len(strX)/2; i++ {
		strX[i], strX[len(strX)-1-i] = strX[len(strX)-1-i], strX[i]
	}
	ans, _ := strconv.Atoi(string(strX))
	if negative {
		ans = -ans
	}
	if ans < -1<<31 || ans > 1<<31-1 {
		return 0
	}

	return ans
}
