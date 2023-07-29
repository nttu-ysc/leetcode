package problem808

func soupServings(n int) float64 {
	if n > 5000 {
		return 1
	}
	if n == 0 {
		return 0.5
	}
	var dp = make(map[[2]int]float64)
	return helper(n, n, &dp)
}

var operations = [][2]int{
	{100, 0},
	{75, 25},
	{50, 50},
	{25, 75},
}

func helper(a, b int, dp *map[[2]int]float64) float64 {
	var p float64

	for _, operation := range operations {
		remainA, remainB := a-operation[0], b-operation[1]
		if remainA <= 0 && remainB > 0 {
			p += 0.25
			continue
		}

		if remainA <= 0 && remainB <= 0 {
			p += 0.25 * 0.5
			continue
		}

		if remainA > 0 && remainB > 0 {
			if pp, ok := (*dp)[[2]int{remainA, remainB}]; !ok {
				p += 0.25 * helper(remainA, remainB, dp)
			} else {
				p += 0.25 * pp
			}
			continue
		}
	}

	(*dp)[[2]int{a, b}] = p
	return p
}
