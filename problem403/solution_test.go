package problem403

import "testing"

func TestCanCRoss(t *testing.T) {
	testCases := []struct {
		stones []int
		want   bool
	}{
		{
			stones: []int{0, 1, 3, 5, 6, 8, 12, 17},
			want:   true,
		},
		{
			stones: []int{0, 1, 2, 3, 4, 8, 9, 11},
			want:   false,
		},
	}

	for _, tC := range testCases {
		if got := canCross(tC.stones); got != tC.want {
			t.Errorf("Test failed. The input stones = %v is expected to %t but return %t", tC.stones, tC.want, got)
		}
	}
}
