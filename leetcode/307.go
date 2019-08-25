package leetcode

/*
Binary Indexed Tree经典解法
执行用时 :72 ms, 在所有 Go 提交中击败了90.48%的用户
内存消耗 :10.1 MB, 在所有 Go 提交中击败了28.57%的用户
 */
type NumArray struct {
	binaryIndexTree
}

type binaryIndexTree struct {
	tree []int
	maxIdx int
	nums[] int
}

func (this *binaryIndexTree) sum(j int) int {
	sum := 0
	for idx := j; idx > 0; {
		sum += this.tree[idx]
		idx -= this.lowbit(idx)
	}
	return sum
}

func (this *binaryIndexTree) upadte(idx int, val int) {
	for ; idx <= this.maxIdx; {
		this.tree[idx] += val
		idx += this.lowbit(idx)
	}
}

func (this * binaryIndexTree) lowbit(x int) int {
	return x - x & (x - 1)
}

func Constructor(nums []int) NumArray {
	tree := make([]int, len(nums) + 1)
	bit := binaryIndexTree {
		tree: tree,
		maxIdx:len(nums),
		nums:nums,
	}
	for i := 0; i < len(nums); i ++ {
		bit.upadte(i + 1, nums[i])
	}
	return NumArray{
		bit,
	}
}

func (this *NumArray) Update(i int, val int)  {
	delta := val - this.nums[i]
	this.nums[i] = val
	this.upadte(i + 1, delta)
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.binaryIndexTree.sum(j + 1) - this.binaryIndexTree.sum(i)
}

