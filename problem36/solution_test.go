package problem36

import "testing"

func TestIsValidSudoku(t *testing.T) {
	testCases := []struct {
		board [][]byte
		want  bool
	}{
		{
			board: [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}},
			want:  true,
		},
	}

	for _, tC := range testCases {
		if got := isValidSudoku(tC.board); got != tC.want {
			t.Errorf("Test failed.")
		}
	}
}
