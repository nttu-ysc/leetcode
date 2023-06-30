# 37. Sudoku Solver
## 題目
Write a program to solve a Sudoku puzzle by filling the empty cells.

A sudoku solution must satisfy all of the following rules:

Each of the digits `1-9` must occur exactly once in each row.  
Each of the digits `1-9` must occur exactly once in each column.  
Each of the digits `1-9` must occur exactly once in each of the 9 `3x3` sub-boxes of the grid.  
The `'.'` character indicates empty cells.

---
## 想法
這題就是就是直接解數獨,因為演算法可能不同,所以可能會有多組答案出現  
對於每個格子判斷規則,如果可以填入就往下走  
***因為每個前面的格子都是符合規則,所以只需要檢查當下的格子***

首先就是對空白的格子做填寫, 填入的數字要符合同row, col, 九宮格 都只出現一次 因此可以按照此規則寫出 func `isValid` 來判斷每個填入數字是有有效  

對於`board` $i$ row, $j$ column, 填入 val 是否有效
```go
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
```

對於每個格子遞迴到完成所有表格
```go
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
```

[完整代碼](./solution.go)

