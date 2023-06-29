package problem37

func solveSudoku(board [][]byte) {
	recursion(board, 0, 0)
}

func recursion(board [][]byte, i, j int) bool {
	if i == 9 {
		return true
	}
	if j >= 9 {
		return recursion(board, i+1, 0)
	}
	if board[i][j] != '.' {
		return recursion(board, i, j+1)
	}
	for char := '1'; char <= '9'; char++ {
		if !isValid(board, i, j, byte(char)) {
			continue
		}
		board[i][j] = byte(char)
		if recursion(board, i, j+1) {
			return true
		}
		board[i][j] = '.'
	}
	return false
}

func isValid(board [][]byte, i, j int, val byte) bool {
	for x := 0; x < 9; x++ {
		if board[x][j] == val {
			return false
		}
	}
	for y := 0; y < 9; y++ {
		if board[i][y] == val {
			return false
		}
	}
	row, col := i-i%3, j-j%3
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if board[x+row][y+col] == val {
				return false
			}
		}
	}
	return true
}
