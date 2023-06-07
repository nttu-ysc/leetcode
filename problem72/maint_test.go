package problem72

import "testing"

func TestMinDistance(t *testing.T) {
	tests := []struct {
		word1 string
		word2 string
		want  int
	}{
		{
			word1: "doge",
			word2: "dog",
			want:  1,
		},
		{
			word1: "horse",
			word2: "ros",
			want:  3,
		},
		{
			word1: "intention",
			word2: "execution",
			want:  5,
		},
	}

	for _, test := range tests {
		if got := minDistance(test.word1, test.word2); got != test.want {
			t.Errorf("Test failed. The input word1=\"%s\" word2=\"%s\" is excepted to %d, but return %d", test.word1, test.word2, test.want, got)
		}
	}

}
