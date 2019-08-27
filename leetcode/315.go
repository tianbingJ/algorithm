package leetcode

/*
按照惯例，先实现直接了当的方法：
执行用时 :560 ms, 在所有 Go 提交中击败了36.59%的用户
内存消耗 :9.4 MB, 在所有 Go 提交中击败了41.67%的用户
func countSmaller(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i ++ {
		count := 0
		for j := i + 1; j < len(nums); j ++ {
			if nums[j] < nums[i] {
				count ++
			}
		}
		result[i] = count
	}
	return result
}
*/

/*
利用归并排序性质：
把数组分成两份segment1 和色segment2， 分别对segment1和segment2进行从到小排序。
当计算segment1和segment2之间逆序对数：
每次从两个有序的列表中选择最大的那个数X，如果相等则，使用右侧的数字（确保处理左侧数字L的时候，右侧的数字都是比L要小）。
1.如果X在segment1，则segment2中剩余的数组都比X大，则count += len(segment2中剩余的个数)
2.如果X在segment2，则不计算逆序对数。

数字排序之后无法得到最初的下表信息，因此把数字和下表绑定在一个，即类型Elem。

执行用时 :328 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :11.1 MB, 在所有 Go 提交中击败了25.00%的用户
*/


type Elem struct {
	num int
	index int
}

func countSmaller(nums []int) []int {
	n := len(nums)
	counts := make([]int, n)
	elems := make([]*Elem, n)
	for i := 0; i < n; i ++ {
		elems[i] = &Elem{
			num:nums[i],
			index:i,
		}
	}
	mergeCount(elems, 0, n-1, counts)
	return counts
}

func mergeCount(nums []*Elem, start, end int, counts []int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeCount(nums, start, mid, counts)
	mergeCount(nums, mid+1, end, counts)
	doMerge(nums, start, mid, mid + 1, end, counts)
}

func doMerge(nums []*Elem, start1, end1, start2, end2 int, counts []int) {
	help := make([]*Elem, end2-start1+1)
	helpIdx := 0

	idx1, idx2 := start1, start2
	for ; idx1 <= end1 && idx2 <= end2; helpIdx++ {
		if nums[idx2].num >= nums[idx1].num {
			help[helpIdx] = nums[idx2]
			idx2++
		} else {
			help[helpIdx] = nums[idx1]
			counts[nums[idx1].index] += end2 - idx2 + 1
			idx1++
		}
	}

	for ; idx1 <= end1; idx1, helpIdx = idx1+1, helpIdx+1 {
		help[helpIdx] = nums[idx1]
	}
	for ; idx2 <= end2; idx2, helpIdx = idx2+1, helpIdx+1 {
		help[helpIdx] = nums[idx2]
	}
	for i := 0; i < len(help); i ++ {
		nums[start1 + i] = help[i]
	}
}
