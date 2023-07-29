package problem808

import "testing"

func TestSoupServings(t *testing.T) {
	testCases := []struct {
		n    int
		want float64
	}{
		{
			n:    50,
			want: 0.625,
		},
		{
			n:    100,
			want: 0.71875,
		},
		{
			n:    0,
			want: 0.5,
		},
		{
			n:    850,
			want: 0.96612,
		},
		{
			n:    660295675,
			want: 1,
		},
	}

	for _, tC := range testCases {
		if got := soupServings(tC.n); got < tC.want-(10e-5) || got > tC.want+(10e-5) {
			t.Errorf("Test failed. The input n = %d is expected to %f but return %f", tC.n, tC.want, got)
		}
	}
}
