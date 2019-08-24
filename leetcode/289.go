package leetcode

var dx = []int {-1, -1, -1, 0,  0, 1,  1, 1}
var dy = []int {-1,  1,  0, 1, -1, 0, -1, 1}

/**
方法1,直接按照原来的逻辑实现，不是原地解决。
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.2 MB, 在所有 Go 提交中击败了25.00%的用户
 */
func gameOfLife(board [][]int)  {
	rows := len(board)
	if rows == 0 {
		return
	}
	cols := len(board[0])
	resultBoard := make([][]int, rows)
	for i := 0; i < rows; i ++ {
		resultBoard[i] = make([]int, cols)
	}

	for i := 0; i < rows; i ++ {
		for j := 0; j < cols; j ++ {
			count := aliveCount(board, i, j)
			if board[i][j] == 1 {
				switch {
				case count < 2:
					resultBoard[i][j] = 0
				case count == 2 || count == 3 :
					resultBoard[i][j] = 1
				default:
					resultBoard[i][j] = 0
				}
			} else {
				if count == 3 {
					resultBoard[i][j] = 1
				} else {
					resultBoard[i][j] = 0
				}
			}
		}
	}

	for i := 0; i < rows; i ++ {
		for j := 0; j < cols; j ++ {
			board[i][j] = resultBoard[i][j]
		}
	}
}

func aliveCount(board[][]int, i, j int) int {
	rows := len(board)
	cols := len(board[0])
	count := 0
	for k := 0; k < 8; k ++ {
		x := i + dx[k]
		y := j + dy[k]
		if x >= 0 && x < rows && y >= 0 && y < cols && board[x][y] == 1{
			count ++
		}
	}
	return count
}

