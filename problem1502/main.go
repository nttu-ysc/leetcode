package problem1502

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	criteria := arr[1] - arr[0]

	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != criteria {
			return false
		}
	}

	return true
}
