package BinaryIndexedTree

/*
树状数组，统计动态update，统计区间和
 */
type BinaryIndexTree struct {
	tree []int
	maxIdx int
}

func Constructor(nums []int) BinaryIndexTree {
	tree := make([]int, len(nums) + 1)
	bit := BinaryIndexTree{
		tree: tree,
		maxIdx:len(nums),
	}
	for i := 0; i < len(nums); i ++ {
		bit.upadte(i + 1, nums[i])
	}
	return bit
}

func (this *BinaryIndexTree) sum(j int) int {
	sum := 0
	for idx := j; idx > 0; {
		sum += this.tree[idx]
		idx -= this.lowbit(idx)
	}
	return sum
}

/**
val is delta
 */
func (this *BinaryIndexTree) upadte(idx int, val int) {
	for ; idx <= this.maxIdx; {
		this.tree[idx] += val
		idx += this.lowbit(idx)
	}
}

func (this * BinaryIndexTree) lowbit(x int) int {
	return x - x & (x - 1)
}
