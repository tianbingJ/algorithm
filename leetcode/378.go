package leetcode

/*

 */
func kthSmallest1(matrix [][]int, k int) int {
	rows := len(matrix)
	cols := len(matrix[0])
	lo := matrix[0][0]
	hi := matrix[rows-1][cols-1]
	for lo < hi {
		mid := (hi + lo) / 2
		if count(matrix, rows, cols, mid) < k {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return hi
}

func count(matrix [][]int, rows, cols, val int) int {
	i := rows - 1
	j := 0
	count := 0
	for i >= 0 && j < cols {
		if matrix[i][j] > val {
			i--
		} else {
			j++
			count += i + 1
		}
	}
	return count
}
