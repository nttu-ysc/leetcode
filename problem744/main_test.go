package problem744

import "testing"

func TestNextGreatestLetter(t *testing.T) {
	tests := []struct {
		letters []byte
		target  byte
		want    byte
	}{
		{
			letters: []byte{'c', 'f', 'j'},
			target:  'a',
			want:    'c',
		},
		{
			letters: []byte{'c', 'f', 'j'},
			target:  'c',
			want:    'f',
		},
		{
			letters: []byte{'x', 'x', 'y', 'y'},
			target:  'z',
			want:    'x',
		},
	}

	for _, test := range tests {
		if got := nextGreatestLetter(test.letters, test.target); got != test.want {
			t.Errorf("Test failed. The input letters=%v target=%v  is excepted to %v but return %v", test.letters, test.target, test.want, got)
		}
	}
}
