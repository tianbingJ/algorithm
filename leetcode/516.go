package leetcode

func longestPalindromeSubseq(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	dp := make([][]int, n)
	for i := 0; i < n; i ++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	result := 1
	for i := 0; i < n; i ++ {
		for j := i - 1; j >= 0; j -- {
			if s[i] == s[j] {
				dp[j][i] = max(dp[j][i], dp[j + 1][i - 1] + 2)
				result = max(dp[j][i], result)
			} else {
				dp[j][i] = max(dp[j][i], dp[j + 1][i])
				dp[j][i] = max(dp[j][i], dp[j][i - 1])
			}
		}
	}
	return result
}