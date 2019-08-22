package leetcode

func findDuplicate(nums []int) int {
	n := len(nums) - 1
	left, right := 1, n
	for left < right {
		mid := (left + right) / 2;
		count := 0
		for _, v := range nums {
			if v <= mid {
				count++
			}
		}
		if count <= mid {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
