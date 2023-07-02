package problem1824

import "testing"

func TestMinSideJumps(t *testing.T) {
	testCases := []struct {
		obstacles []int
		want      int
	}{
		// {
		// 	obstacles: []int{0, 0, 3, 1, 0, 1, 0, 2, 3, 1, 0},
		// 	want:      2,
		// },
		{
			obstacles: []int{0, 1, 2, 3, 0},
			want:      2,
		},
		{
			obstacles: []int{0, 1, 1, 3, 3, 0},
			want:      0,
		},
		{
			obstacles: []int{0, 2, 1, 0, 3, 0},
			want:      2,
		},
	}

	for _, tC := range testCases {
		if got := minSideJumps(tC.obstacles); got != tC.want {
			t.Errorf("Test failed. The input obstacles = %v is expected to %d but return %d", tC.obstacles, tC.want, got)
		}
	}
}
