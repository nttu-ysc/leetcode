package problem202

import "math"

func isHappy(n int) bool {
	var tmp int
	r := make(map[int]bool)
	r[n] = true

	for n != 1 {
		tmp = 0
		for i := 0; n >= int(math.Pow10(i)); i++ {
			tmp += int(math.Pow(float64((n/int(math.Pow10(i)))%10), 2))
		}
		if _, ok := r[tmp]; ok {
			return false
		}
		r[tmp] = true
		n = tmp
	}

	return true
}
