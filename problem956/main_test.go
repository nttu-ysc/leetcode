package problem956

import "testing"

func TestTallestBillboard(t *testing.T) {
	testCases := []struct {
		rods []int
		want int
	}{
		{
			rods: []int{1, 2, 3, 6},
			want: 6,
		},
		{
			rods: []int{1, 2, 3, 4, 5, 6},
			want: 10,
		},
		{
			rods: []int{1, 2},
			want: 0,
		},
	}

	for _, tC := range testCases {
		if got := tallestBillboard(tC.rods); got != tC.want {
			t.Errorf("Test failed. The input rods = %v is expected to %d but return %d", tC.rods, tC.want, got)
		}
	}

}
