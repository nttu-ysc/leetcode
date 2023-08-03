package problem17

import "testing"

func TestLetterCombinations(t *testing.T) {
	testCases := []struct {
		digits string
		want   []string
	}{
		{
			digits: "23",
			want:   []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			digits: "",
			want:   []string{},
		},
		{
			digits: "2",
			want:   []string{"a", "b", "c"},
		},
	}

	for _, tC := range testCases {
		got := letterCombinations2(tC.digits)
		if len(got) != len(tC.want) {
			t.Errorf("Test failed.")
		}

		for i := range got {
			if got[i] != tC.want[i] {
				t.Errorf("Test failed.")
				break
			}
		}
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		letterCombinations("23")
	}
}

func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		letterCombinations2("23")
	}
}
