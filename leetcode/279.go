package leetcode

/**

对于数字f(n)，结果= min(f(i) + f(n - i)), 0 < i < n

动态规划， N平方的时间复杂度，
执行用时 :1284 ms, 在所有 Go 提交中击败了6.92%的用户
内存消耗 :5.9 MB, 在所有 Go 提交中击败了47.50%的用户
 */
func numSquares1(n int) int {
	nums := make([]int, 0)
	dp := make([]int, n + 1)
	for i := 1; i * i <= n; i ++ {
		dp[i * i] = 1
		nums = append(nums, i * i)
	}

	for i := 2; i <= n ;i ++ {
		if dp[i] == 1 {
			continue
		}
		min := i

		for j := i - 1; j >= 1; j -- {
			if dp[j] + dp[i - j] < min {
				min = dp[j] + dp[i-j];
			}
		}
		dp[i] = min
	}
	return dp[n]
}

