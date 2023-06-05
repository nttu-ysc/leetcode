package problem1232

import "math"

func checkStraightLine(coordinates [][]int) bool {
	p := calcSlope(coordinates[0], coordinates[1])

	for i := 2; i < len(coordinates); i++ {
		if calcSlope(coordinates[0], coordinates[i]) != p {
			return false
		}
	}

	return true
}

func calcSlope(a, b []int) float64 {
	if b[0]-a[0] == 0 {
		return math.Inf(1)
	}
	return (float64(a[1]) - float64(b[1])) / (float64(a[0]) - float64(b[0]))
}
