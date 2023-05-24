package main

func oddCells(m int, n int, indices [][]int) int {
	x := make(map[int]int)
	y := make(map[int]int)
	var counter int

	for _, index := range indices {
		y[index[0]]++
		x[index[1]]++
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if (x[c]+y[r])%2 == 1 {
				counter++
			}
		}
	}

	return counter
}
