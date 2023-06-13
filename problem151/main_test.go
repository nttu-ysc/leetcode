package problem151

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			s:    "the sky is blue",
			want: "blue is sky the",
		},
		{
			s:    "  hello world  ",
			want: "world hello",
		},
		{
			s:    "a good   example",
			want: "example good a",
		},
	}

	for _, test := range tests {
		if got := reverseWords(test.s); got != test.want {
			t.Errorf("Test failed. The input s=\"%s\" is excepted to \"%s\" but return \"%s\"", test.s, test.want, got)
		}
	}
}

const benchmarkWord = "a good   example"

func BenchmarkReverseWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseWords(benchmarkWord)
	}
}

func BenchmarkReverseWords2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseWords2(benchmarkWord)
	}
}
