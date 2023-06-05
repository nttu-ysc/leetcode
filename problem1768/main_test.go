package problem1768

import "testing"

func TestMergeAlternately(t *testing.T) {
	tests := []struct {
		word1 string
		word2 string
		want  string
	}{
		{
			word1: "abc",
			word2: "pqr",
			want:  "apbqcr",
		},
		{
			word1: "ab",
			word2: "pqrs",
			want:  "apbqrs",
		},
		{
			word1: "abcd",
			word2: "pq",
			want:  "apbqcd",
		},
	}

	for _, test := range tests {
		if got := mergeAlternately(test.word1, test.word2); got != test.want {
			t.Errorf("Test failed. The input word1=\"%s\" word2=\"%s\" is excepted to \"%s\", but return \"%s\"", test.word1, test.word2, test.want, got)
		}
	}
}
