package leetcode

import "sort"

func hIndex(citations []int) int {

	sort.Sort(sort.IntSlice(citations))
	for i := 0; i <  len(citations); i ++ {
		if citations[i] >= len(citations) - i {
			return len(citations) - i
		}
	}
	return 0
}