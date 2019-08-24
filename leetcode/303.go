package leetcode

/**
执行用时 :48 ms, 在所有 Go 提交中击败了93.23%的用户
内存消耗 :7.4 MB, 在所有 Go 提交中击败了62.50%的用户

type NumArray struct {
	prefixSum[]int
}

func Constructor(nums []int) NumArray {
	prefixSum := make([]int, len(nums))
	for i := 0; i < len(nums); i ++ {
		if i == 0 {
			prefixSum[i] = nums[0]
		} else {
			prefixSum[i] = prefixSum[i - 1] + nums[i];
		}
	}
	return NumArray{
		prefixSum:prefixSum,
	}
}


func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.prefixSum[j]
	} else {
		return this.prefixSum[j] - this.prefixSum[i - 1]
	}
}
*/
