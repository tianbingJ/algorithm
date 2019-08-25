package leetcode

import (
	"strconv"
)
/**
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2 MB, 在所有 Go 提交中击败了100.00%
 */

func isAdditiveNumber(num string) bool {
	for i := 0; i < len(num); i++ {
		for j := i + 1; j < len(num)-1; j++ {
			if isAdditive(num, i, j) {
				return true
			}
		}
	}
	return false
}

func isAdditive(num string, firstEndIndex, secondEndIndex int) bool {
	idx1Start := 0
	idx1End := firstEndIndex
	idx2Start := firstEndIndex + 1
	idx2End := secondEndIndex

	for {
		num1 := num[idx1Start : idx1End+1]
		if len(num1) > 1 && num1[0] == '0' {
			return false
		}
		num2 := num[idx2Start : idx2End+1]
		if len(num2) > 1 && num2[0] == '0' {
			return false
		}
		n1, _ := strconv.Atoi(num1)
		n2, _ := strconv.Atoi(num2)
		sum := strconv.Itoa(n1 + n2)
		if len(num) < idx2End +len(sum) + 1{
			return false
		}
		if num[idx2End+1: idx2End + len(sum) + 1] != sum {
			return false
		}
		idx1Start = idx2Start
		idx1End = idx2End
		idx2Start = idx1End + 1
		idx2End = idx1End + len(sum)
		if idx2End > len(num)-1 {
			return false
		}
		if idx2End == len(num)-1 {
			break
		}
	}
	return true
}
