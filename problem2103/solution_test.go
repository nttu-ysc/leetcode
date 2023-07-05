package problem2103

import "testing"

func TestCountPoints(t *testing.T) {
	testCases := []struct {
		rings string
		want  int
	}{
		{
			rings: "B0B6G0R6R0R6G9",
			want:  1,
		},
		{
			rings: "B0R0G0R9R0B0G0",
			want:  1,
		},
		{
			rings: "G4",
			want:  0,
		},
	}

	for _, tC := range testCases {
		if got := countPoints(tC.rings); got != tC.want {
			t.Errorf("Test failed. The rings = \"%s\" is expected to %d but return %d", tC.rings, tC.want, got)
		}
	}
}
