package problem1318

import "testing"

func TestMinFlips(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		c    int
		want int
	}{
		{
			a:    2,
			b:    6,
			c:    5,
			want: 3,
		},
		{
			a:    4,
			b:    2,
			c:    7,
			want: 1,
		},
		{
			a:    1,
			b:    2,
			c:    3,
			want: 0,
		},
		{
			a:    0,
			b:    0,
			c:    1,
			want: 1,
		},
	}

	for _, test := range tests {
		if got := minFlips(test.a, test.b, test.c); got != test.want {
			t.Errorf("Test failed. The input a=%d b=%d c=%d is expected to %d but return %d", test.a, test.b, test.c, test.want, got)
		}
	}
}
