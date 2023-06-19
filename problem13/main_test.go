package problem13

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{
			s:    "III",
			want: 3,
		},
		{
			s:    "LVIII",
			want: 58,
		},
		{
			s:    "MCMXCIV",
			want: 1994,
		},
	}

	for _, test := range tests {
		if got := romanToInt(test.s); got != test.want {
			t.Errorf("Test failed. The s=\"%s\" is excepted to %d but return %d", test.s, test.want, got)
		}
	}

}
