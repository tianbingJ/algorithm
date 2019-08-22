package leetcode

import "sort"

func hIndex(citations []int) int {
	sort.Sort(sort.IntSlice(citations))
	for i := len(citations) - 1; i >= 0; i -- {
		if citations[i] >= i {
			return i + 1
		}
	}
	return 0
}