package problem121

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
	}

	for _, test := range tests {
		if got := maxProfit(test.prices); got != test.want {
			t.Errorf("Test failed. The input prices = %v is expected to %d but return %d", test.prices, test.want, got)
		}
	}
}
