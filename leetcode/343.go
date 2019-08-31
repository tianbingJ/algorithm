package leetcode

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2 MB, 在所有 Go 提交中击败了45.00%的用户
 */
func integerBreak(n int) int {

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 1

	for i := 2; i <= n; i++ {
		max := 0
		for j := 0; j < i; j++ {
			if max < j * (i - j) {
				max = j * (i - j)
			}
			if max < j * dp[i-j] {
				max = j * dp[i - j]
			}
		}
		dp[i] = max
	}
	return dp[n]
}
