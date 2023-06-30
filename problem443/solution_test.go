package problem443

import "testing"

func TestCompress(t *testing.T) {
	testCases := []struct {
		chars []byte
		want  int
	}{
		{
			chars: []byte{'a', 'a', 'a', 'a', 'b', 'a'},
			want:  4,
		},
		{
			chars: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			want:  6,
		},
		{
			chars: []byte{'a'},
			want:  1,
		},
		{
			chars: []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			want:  4,
		},
	}

	for _, tC := range testCases {
		if got := compress(tC.chars); got != tC.want {
			t.Errorf("Test failed. The chars = %v is expected to %d but return %d", tC.chars, tC.want, got)
		}
	}
}

var chars = []byte{'a', 'a', 'a', 'a', 'b', 'a', 's', 's', 'a'}

func BenchmarkCompress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		compress(chars)
	}
}
