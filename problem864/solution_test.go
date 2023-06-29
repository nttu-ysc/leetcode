package problem864

import "testing"

func TestShortestPathAllKeys(t *testing.T) {
	testCases := []struct {
		grid []string
		want int
	}{
		{
			grid: []string{"@...a", "###.#", "b.A.B"},
			want: 10,
		},
		{
			grid: []string{"@.a..", "###.#", "b.A.B"},
			want: 8,
		},
		{
			grid: []string{"@..aA", "..B#.", "....b"},
			want: 6,
		},
		{
			grid: []string{"@Aa"},
			want: -1,
		},
		{
			grid: []string{"@abcdeABCDEFf"},
			want: -1,
		},
		{
			grid: []string{"Dd#b@", ".fE.e", "##.B.", "#.cA.", "aF.#C"},
			want: 14,
		},
		{
			grid: []string{"..#....##.", "....d.#.D#", "#...#.c...", "..##.#..a.", "...#....##", "#....b....", ".#..#.....", "..........", ".#..##..A.", ".B..C.#..@"},
			want: 19,
		},
	}

	for _, tC := range testCases {
		if got := shortestPathAllKeys(tC.grid); got != tC.want {
			t.Errorf("Test failed. The input grid = %v is expected to %d but return %d", tC.grid, tC.want, got)
		}
	}
}

func BenchmarkShortestPathAllKeys(b *testing.B) {
	grid := []string{"..#....##.", "....d.#.D#", "#...#.c...", "..##.#..a.", "...#....##", "#....b....", ".#..#.....", "..........", ".#..##..A.", ".B..C.#..@"}
	for i := 0; i < b.N; i++ {
		shortestPathAllKeys(grid)
	}
}
