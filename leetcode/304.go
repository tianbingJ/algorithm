package leetcode

/**
预处理：
areaSum[i][j]表示包含从(0,0)到(i,j)表示的矩形区域的和。
SumRegion(row1, col1, row2, col2) =
	areaSum[row2][col2]  + areaSum[row1-1][col1-1] - areaSum[row1-1][col2] - aresSum[row2][col1 - 1]

执行用时 :52 ms, 在所有 Go 提交中击败了86.49%的用户
内存消耗 :8 MB, 在所有 Go 提交中击败了100.00%的用户
 */

type NumMatrix struct {
	aresSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	areaSum := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i ++ {
		areaSum[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j ++ {
			if j == 0 {
				areaSum[i][j] = matrix[i][j]
			} else {
				areaSum[i][j] = areaSum[i][j - 1] + matrix[i][j]
			}
		}
		if i > 0 {
			for j := 0; j < len(matrix[i]); j ++ {
				areaSum[i][j] += areaSum[i - 1][j]
			}
		}
	}
	return NumMatrix{
		aresSum:areaSum,
	}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := this.aresSum[row2][col2];
	if row1 > 0 && col1 > 0 {
		sum += this.aresSum[row1 - 1][col1 - 1]
	}
	if row1 > 0 {
		sum -= this.aresSum[row1 - 1][col2]
	}
	if col1 > 0 {
		sum -= this.aresSum[row2][col1 - 1]
	}
	return sum
}

