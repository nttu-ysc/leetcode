package problem2024

import "testing"

func TestMaxConsecutiveAnswers(t *testing.T) {
	testCases := []struct {
		answerKey string
		k         int
		want      int
	}{
		{
			answerKey: "TTFF",
			k:         2,
			want:      4,
		},
		{
			answerKey: "TFFT",
			k:         1,
			want:      3,
		},
		{
			answerKey: "TTFTTFTT",
			k:         1,
			want:      5,
		},
	}

	for _, tC := range testCases {
		if got := maxConsecutiveAnswers(tC.answerKey, tC.k); got != tC.want {
			t.Errorf("Test failed. The input answerKey = \"%s\", k = %d is expected to %d but return %d", tC.answerKey, tC.k, tC.want, got)
		}
	}
}
