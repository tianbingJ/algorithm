package leetcode

import "sort"

func wiggleSort(nums []int) {
	result := make([]int, len(nums))
	n := len(nums)
	sort.Ints(nums)
	midValue := nums[(n - 1)/ 2]

	for i, idx := 1, n - 1; i < n ; i, idx = i + 2, idx - 1 {
		if nums[idx] > midValue {
			result[i] = nums[idx]
		} else {
			result[i]= midValue
		}
	}

	i := n - 1
	if i % 2 == 1 {
		i -= 1
	}
	for idx := 0; i >= 0; i, idx = i - 2, idx + 1 {
		if nums[idx] < midValue {
			result[i] = nums[idx]
		} else {
			result[i] = midValue
		}
	}
	copy(nums, result)
}
