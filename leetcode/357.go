package leetcode

/**
考虑N位数字：
第一位可以是1 ~ 9的任意一个，有9种情况
第二位可以是0 ~ 9中除了第一位数字的任意一个
第三位有8种选择，依次类推
..

执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2 MB, 在所有 Go 提交中击败了75.00%的用户
 */
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	if n > 10 {
		n = 10
	}

	lastDp := 10
	dp := 0

	nFaco := 9
	for i := 2; i <= n; i ++ {
		nFaco = nFaco * (10 - i + 1);
		dp = nFaco + lastDp
		lastDp = dp
	}
	return dp
}
