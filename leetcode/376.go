package leetcode

const (
	END_BIG   = 0
	END_SMALL = 1
)

/**
自增子序列思想
O(n^2的方法)
*/
func wiggleMaxLength1(nums []int) int {
	dp := make([][]int, len(nums))
	maxLength := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, 2)
		dp[i][END_BIG] = 1
		dp[i][END_SMALL] = 1

		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j][END_SMALL]+1 > dp[i][END_BIG] {
				dp[i][END_BIG] = dp[j][END_SMALL] + 1
			} else if nums[i] < nums[j] && dp[j][END_BIG]+1 > dp[i][END_SMALL] {
				dp[i][END_SMALL] = dp[j][END_BIG] + 1
			}
		}
		maxLength = max(maxLength, dp[i][END_BIG])
		maxLength = max(maxLength, dp[i][END_SMALL])
	}
	return maxLength
}

func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	n := len(nums)
	up := make([]int, n)
	down := make([]int, n)
	up[0] = 1
	down[0] = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i - 1] {
			up[i] = down[i - 1] + 1
			down[i] = down[i - 1]
		} else if nums[i] < nums[i - 1]{
			up[i] = up[i - 1]
			down[i] = up[i - 1] + 1

		} else {
			up[i] = up[i - 1]
			down[i] = down[i - 1]
		}
	}
	return max(up[n - 1], down[n - 1])
}
