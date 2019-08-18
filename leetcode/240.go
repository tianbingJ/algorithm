package leetcode

/**
1.暴力方法：O（MN），没有用到有序的特性
执行用时 :68 ms, 在所有 Go 提交中击败了20.56%的用户
*/
func searchMatrix1(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == target {
				return true;
			}
		}
	}
	return false
}

/**
方法2:对每一行进行二分，复杂度 O(M*lg(N)), 没有用到按列有序的特性
执行用时：40 ms，战胜68.22%的用户
*/
func searchMatrix2(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if binarySearch(matrix[i], target) {
			return true
		}
	}
	return false
}

func binarySearch(data []int, target int) bool {
	for l, h := 0, len(data)-1; l < h; {
		mid := (l + h) / 2;
		if data[mid] == target {
			return true
		}
		if data[mid] > target {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

/*
方法3：在方法2的循环中加入一点优化，快了一点点
执行用时：32 ms, 在所有 Go 提交中击败了85%的用户
	if matrix[i][0] > target {
		continue
	}
	if matrix[i][len(matrix[i]) - 1] < target {
		continue
	}
*/

/*
方法4：前几种方法一致没有利用到列的有序性， 后来想到可以逐步迭代进行，利用行的有序性和列的有序性进行迭代,但一直没想明白算法的正确性。
从左下角出发：
(1)如果当前值== target, return true
(2)如果当前值 > target, row --
(3)否则当前值 < target， col ++
正确性分析:
 a a | a | a a a
 a a | a | a a a
 - - - - - - - -
 a a | c | a a a
 - - - - - - - -
 a a | a | a a a
 a a | a | a a a

假设当前所在位置是c，由于按行、列递增的特性，不包括c所在行、列，其余部分被分为4个矩阵。
可以判断出c位置左上角的矩阵一定比c小，右下角的矩阵一定比c大，算法的思路是不断调整这四个矩阵的大小，逼近target的过程。


执行用时 :36 ms, 在所有 Go 提交中击败了85.05%的用户.
算法复杂度是O(M + N)，但是时间上跟优化后的第三种方法一样，测试数据不够复杂。

*/
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := len(matrix) - 1
	col := 0
	for {
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] > target {
			row--
		} else {
			col++
		}
		if row < 0 || col >= len(matrix[0]) {
			return false
		}
	}
}
