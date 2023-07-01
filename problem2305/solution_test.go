package problem2305

import "testing"

func TestDistributeCookies(t *testing.T) {
	testCases := []struct {
		cookies []int
		k       int
		want    int
	}{
		{
			cookies: []int{8, 15, 10, 20, 8},
			k:       2,
			want:    31,
		},
		{
			cookies: []int{6, 1, 3, 2, 2, 4, 1, 2},
			k:       3,
			want:    7,
		},
	}

	for _, tC := range testCases {
		if got := distributeCookies(tC.cookies, tC.k); got != tC.want {
			t.Errorf("Test failed. The input cookies = %v, k = %d is expected to %d but return %d", tC.cookies, tC.k, tC.want, got)
		}
	}
}
