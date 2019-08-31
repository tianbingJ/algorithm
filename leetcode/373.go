package leetcode

import (
	"container/heap"
	"math"
)

/**
nums1和nums2构成所有的pairs
方法1.暴力：
排序所有的pairs, 取前k个。
时间复杂度: O(MN Log(MN)) M和N分别是两个队列的长度

方法2：使用堆
遍历所有pair, 维护一个K大小的大根堆
最后返回堆中的元素。时间复杂度 O(M * N * LogK)
考虑到K < M * N，这种方法在k比较小的时候比方法1要快

执行用时 :544 ms, 在所有 Go 提交中击败了15.79%的用户
内存消耗 :21.9 MB, 在所有 Go 提交中击败了33.33%的用户
 */

type Pairs struct {
	pairs [][]int
}

func (p * Pairs) Len() int {
	return len(p.pairs)
}
// Less reports whether the element with
// index i should sort before the element with index j.
func (p *Pairs)Less(i, j int) bool {
	return p.pairs[i][0] + p.pairs[i][1] > p.pairs[j][0] + p.pairs[j][1]
}
// Swap swaps the elements with indexes i and j.
func (p *Pairs)Swap(i, j int) {
	p.pairs[i], p.pairs[j] = p.pairs[j], p.pairs[i]
}

func (p *Pairs) Push(x interface{}) {
	switch  v := x.(type) {
	case []int:
		p.pairs = append(p.pairs, v)
	}
}

func (p *Pairs) Pop() interface{}   {
	v := p.pairs[len(p.pairs) - 1]
	p.pairs = p.pairs[0: len(p.pairs) - 1]
	return v
}

func kSmallestPairs1(nums1 []int, nums2 []int, k int) [][]int {
	if k == 0 {
		return [][]int{}
	}
	pairs := make([][]int, len(nums1) * len(nums2))
	for i := 0; i < len(nums1); i ++{
		for j := 0; j < len(nums2); j ++ {
			pairs[i * len(nums2) + j] = []int{nums1[i], nums2[j]}
		}
	}

	var heapData * Pairs = & Pairs{
		pairs:make([][]int, 0),
	}
	heap.Init(heapData)
	for i := 0; i < len(pairs);i ++ {
		if i < k {
			heap.Push(heapData, pairs[i])
		} else {
			heap.Push(heapData, pairs[i])
			heap.Pop(heapData)
		}
	}

	result := make([][]int, 0)
	for i := 0; i < k && len(heapData.pairs) > 0; i ++ {
		data := heap.Pop(heapData)
		switch v := data.(type){
		case []int:
			result = append(result, v)
		}
	}
	return result
}

/**

方法3
类似前面的丑数问题,记录每个数字分别
step[i]表示(nums1[i] ,X)是已经产生的丑数，其中X可以是：nums2[[0 ~ step[i])]

时间复杂度：
O(K * N1)
执行用时 :32 ms, 在所有 Go 提交中击败了52.63%的用户
内存消耗 :7.2 MB, 在所有 Go 提交中击败了100.00%的用户
 */
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	steps := make([]int, len(nums1))

	result := make([][]int, 0)
	for i := 0; i < k; i ++ {
		min := math.MaxInt32
		idxNums1 := -1
		idxNums2 := -1
		for j := 0; j < len(nums1); j ++ {
			if steps[j] >= len(nums2) {
				continue
			}
			if nums1[j] + nums2[steps[j]] < min {
				min = nums1[j] + nums2[steps[j]]
				idxNums1 = j
				idxNums2 = steps[j]
			}
		}

		if idxNums1 == -1 {
			break
		}
		steps[idxNums1] ++
		result = append(result, []int{nums1[idxNums1], nums2[idxNums2]})
	}
	return result
}
