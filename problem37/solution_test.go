package problem37

import "testing"

func TestSolveSudoku(t *testing.T) {
	testCases := []struct {
		board [][]byte
		want  [][]byte
	}{
		{
			board: [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}},
			want:  [][]byte{{'5', '3', '4', '6', '7', '8', '9', '1', '2'}, {'6', '7', '2', '1', '9', '5', '3', '4', '8'}, {'1', '9', '8', '3', '4', '2', '5', '6', '7'}, {'8', '5', '9', '7', '6', '1', '4', '2', '3'}, {'4', '2', '6', '8', '5', '3', '7', '9', '1'}, {'7', '1', '3', '9', '2', '4', '8', '5', '6'}, {'9', '6', '1', '5', '3', '7', '2', '8', '4'}, {'2', '8', '7', '4', '1', '9', '6', '3', '5'}, {'3', '4', '5', '2', '8', '6', '1', '7', '9'}},
		},
	}

	for _, tC := range testCases {
		board := tC.board
		solveSudoku(tC.board)
		for i := 0; i < len(tC.want); i++ {
			for j := 0; j < len(tC.want[i]); j++ {
				if tC.want[i][j] != tC.board[i][j] {
					t.Errorf("Test failed. The input board = %v is expected to %v but return %v", board, tC.want, tC.board)
					goto b
				}
			}
		}
	b:
	}
}
