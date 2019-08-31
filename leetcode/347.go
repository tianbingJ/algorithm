package leetcode

import "sort"

/*
leetcode的运行时间不太稳定。
执行用时 :20 ms, 在所有 Go 提交中击败了95.54%的用户
内存消耗 :6.3 MB, 在所有 Go 提交中击败了78.79%的用户

两个map，一个统计num的频率，另外一个记录一个频率对应的数字列表。

时间复杂度： O(N + MlogM)，M表示不同频率的个数。
如果M接近N的话，时间复杂度则接近O(NlogN)。
 */

func topKFrequent(nums []int, k int) []int {

	countM := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		countM[nums[i]] ++
	}

	//freq -> numList
	freqM := make(map[int][]int)
	freqs := make([]int, 0)
	for num, freq := range countM {
		if _, ok := freqM[freq]; !ok {
			freqs = append(freqs, freq)
		}
		freqM[freq] = append(freqM[freq], num)
	}

	sort.Ints(freqs)
	result := make([]int, 0)
	for i := len(freqs) - 1; i >= 0 && k > 0; i -- {
		numList := freqM[freqs[i]]
		result = append(result, numList...)
		k -= len(numList)
	}
	return result
}
