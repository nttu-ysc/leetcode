package problem2710

import "testing"

func TestRemoveTrailingZeros(t *testing.T) {
	tests := []struct {
		num  string
		want string
	}{
		{
			num:  "51230100",
			want: "512301",
		},
		{
			num:  "123",
			want: "123",
		},
	}

	for _, test := range tests {
		if got := removeTrailingZeros(test.num); got != test.want {
			t.Errorf("Test failed. The input num =\"%s\" is expected to \"%s\" but return \"%s\"", test.num, test.want, got)
		}
	}
}

func BenchmarkRemoveTrailingZeros(b *testing.B) {
	const num = "51230100"
	for i := 0; i < b.N; i++ {
		removeTrailingZeros(num)
	}
}

func BenchmarkRemoveTrailingZeros2(b *testing.B) {
	const num = "51230100"
	for i := 0; i < b.N; i++ {
		removeTrailingZeros2(num)
	}
}
