package leetcode

import "strings"
/**

O(N)的时间复杂度，O(1)的空间复杂度。
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.8 MB, 在所有 Go 提交中击败了100.00%的用户
 */

func isValidSerialization(preorder string) bool {
	serial := strings.Split(preorder, ",")
	if len(serial) == 0 {
		return false
	}
	if serial[0] == "#" {
		return len(serial) == 1
	}
	valid , endIndex  := validate(serial, 0)
	return valid && endIndex == len(serial) - 1
}

func validate(sequence []string, idx int) (valid bool,  endIdx int) {
	if idx >= len(sequence) {
		return false, -1
	}
	if sequence[idx] == "#" {
		return true, idx
	}
	leftValid, leftEndIdx := validate(sequence, idx + 1)
	if !leftValid {
		return false, -1
	}
	return validate(sequence, leftEndIdx + 1)
}