package problem58

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		// {
		// 	s:    "Hello World",
		// 	want: 5,
		// },
		// {
		// 	s:    "   fly me   to   the moon  ",
		// 	want: 4,
		// },
		// {
		// 	s:    "luffy is still joyboy",
		// 	want: 6,
		// },
		{
			s:    "a",
			want: 1,
		},
	}

	for _, test := range tests {
		if got := lengthOfLastWord(test.s); got != test.want {
			t.Errorf("Test failed. the input s=\"%s\" is expected to %d but return %d", test.s, test.want, got)
		}
	}
}
