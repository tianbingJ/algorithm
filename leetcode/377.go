package leetcode

/**
搜索超时
 */
func combinationSum41(nums []int, target int) int {
	if  target == 0 {
		return 1
	}
	if target < 0 {
		return 0
	}
	return countCombination(nums, target)
}

func countCombination(nums[]int, target int) int {
	count := 0
	for i := 0; i < len(nums); i ++ {
		if target - nums[i] == 0 {
			count ++
		}
		if target - nums[i]  > 0 {
			count += countCombination(nums, target - nums[i])
		}
	}
	return count
}

/**
可以用动态规划
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.1 MB, 在所有 Go 提交中击败了33.33%的用户
 */
func combinationSum4(nums []int, target int) int {
	if  target == 0 {
		return 1
	}
	if target < 0 {
		return 0
	}
	n := len(nums)
	dp := make([]int, target + 1)
	for i := 0; i < n; i ++ {
		if nums[i] <= target {
			dp[nums[i]] = 1
		}
	}

	for i := 0; i <= target; i ++ {
		count := dp[i]
		for j := 0; j < n; j ++ {
			if i - nums[j] >= 0 {
				count += dp[i - nums[j]]
			}
		}
		dp[i] = count
	}
	return dp[target]
}
