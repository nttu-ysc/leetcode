package problem547

import "testing"

func TestFineCircleNum(t *testing.T) {
	tests := []struct {
		isConnected [][]int
		want        int
	}{
		// {
		// 	isConnected: [][]int{
		// 		{1, 1, 0},
		// 		{1, 1, 0},
		// 		{0, 0, 1},
		// 	},
		// 	want: 2,
		// },
		// {
		// 	isConnected: [][]int{
		// 		{1, 0, 0},
		// 		{0, 1, 0},
		// 		{0, 0, 1},
		// 	},
		// 	want: 3,
		// },
		{
			isConnected: [][]int{
				{1, 0, 0, 1},
				{0, 1, 1, 0},
				{0, 1, 1, 1},
				{1, 0, 1, 1},
			},
			want: 1,
		},
	}

	for _, test := range tests {
		if got := findCircleNum(test.isConnected); got != test.want {
			t.Errorf("Test failed. The input isConnected=%v is excepted to %d, but return %d", test.isConnected, test.want, got)
		}
	}
}
