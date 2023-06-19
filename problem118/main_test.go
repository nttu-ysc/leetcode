package problem118

import "testing"

func TestGenerate(t *testing.T) {
	tests := []struct {
		numRows int
		want    [][]int
	}{
		{
			numRows: 5,
			want: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
		{
			numRows: 1,
			want: [][]int{
				{1},
			},
		},
	}

	for _, test := range tests {
		got := generate(test.numRows)
		for i := 0; i < len(test.want); i++ {
			if len(test.want) != len(got) {
				t.Errorf("Test failed.")
				break
			}
			for j := 0; j < len(test.want[i]); j++ {
				if len(test.want[i]) != len(got[i]) {
					t.Errorf("Test filed.")
					break
				}
				if got[i][j] != test.want[i][j] {
					t.Errorf("Test failed. The numRows = %d is excepted to %v but return %v", test.numRows, test.want, got)
					break
				}
			}
		}
	}
}
