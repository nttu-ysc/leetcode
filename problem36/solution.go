package problem36

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if !isValid(board, i, j, board[i][j]) {
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, i, j int, val byte) bool {
	for x := 0; x < 9; x++ {
		if x == i {
			continue
		}
		if board[x][j] == val {
			return false
		}
	}
	for y := 0; y < 9; y++ {
		if y == j {
			continue
		}
		if board[i][y] == val {
			return false
		}
	}
	row, col := i-i%3, j-j%3
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if x+row == i && y+col == j {
				continue
			}
			if board[x+row][y+col] == val {
				return false
			}
		}
	}
	return true
}
