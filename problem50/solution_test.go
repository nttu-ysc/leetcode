package problem50

import "testing"

func TestMyPow(t *testing.T) {
	TestCases := []struct {
		x    float64
		n    int
		want float64
	}{
		{
			x:    2,
			n:    10,
			want: 1024,
		},
		{
			x:    2.1,
			n:    3,
			want: 9.261,
		},
		{
			x:    2,
			n:    -2,
			want: 0.25,
		},
		{
			x:    1,
			n:    -2147483648,
			want: 1,
		},
	}

	for _, tC := range TestCases {
		if got := myPow(tC.x, tC.n); got != tC.want {
			t.Errorf("Test failed. The input x = %f, n = %d is expected to %f but return %f", tC.x, tC.n, tC.want, got)
		}
	}
}
