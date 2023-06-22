package problem714

import "testing"

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		prices []int
		fee    int
		want   int
	}{
		{
			prices: []int{1, 3, 2, 8, 4, 9},
			fee:    2,
			want:   8,
		},
		{
			prices: []int{1, 3, 7, 5, 10, 3},
			fee:    3,
			want:   6,
		},
	}
	for _, tC := range testCases {
		if got := maxProfit(tC.prices, tC.fee); got != tC.want {
			t.Errorf("Test failed. The prices = %v, fee = %d is expected to %d but return %d", tC.prices, tC.fee, tC.want, got)
		}
	}
}
