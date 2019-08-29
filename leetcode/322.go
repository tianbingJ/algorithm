package leetcode

import "math"

/**
动态规划，有点像背包问题

效率不是很高

执行用时 :20 ms, 在所有 Go 提交中击败了55.77%的用户
内存消耗 :6.1 MB, 在所有 Go 提交中击败了54.29%的用户
 */
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount + 1)
	for i := 1; i <= amount; i ++ {
		m := math.MaxInt32
		for j := 0; j < len(coins); j ++ {
			coin := coins[j]
			if i - coin >= 0 && dp[i - coin] != -1 && m > dp[i - coin] +1 {
				m = dp[i - coin] + 1
			}
		}
		if m == math.MaxInt32 {
			dp[i] = -1
		} else {
			dp[i] = m
		}
	}
	return dp[amount]
}
