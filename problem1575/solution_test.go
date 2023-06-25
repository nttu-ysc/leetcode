package problem1575

import "testing"

func TestCountRoutes(t *testing.T) {
	testCases := []struct {
		locations []int
		start     int
		finish    int
		fuel      int
		want      int
	}{
		{
			locations: []int{2, 3, 6, 8, 4},
			start:     1,
			finish:    3,
			fuel:      5,
			want:      4,
		},
		{
			locations: []int{4, 3, 1},
			start:     1,
			finish:    0,
			fuel:      6,
			want:      5,
		},
		{
			locations: []int{5, 2, 1},
			start:     0,
			finish:    2,
			fuel:      3,
			want:      0,
		},
		{
			locations: []int{1, 2, 3},
			start:     0,
			finish:    2,
			fuel:      40,
			want:      615088286,
		},
	}

	for _, tC := range testCases {
		if got := countRoutes(tC.locations, tC.start, tC.finish, tC.fuel); got != tC.want {
			t.Errorf("Test failed. The input locations = %v, start = %d, finish = %d, fuel = %d is expected to %d but return %d", tC.locations, tC.start, tC.finish, tC.fuel, tC.want, got)
		}
	}
}
