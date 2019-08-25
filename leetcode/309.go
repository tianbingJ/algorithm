package leetcode

const (
	buy = 1
	sell = 0
)

/*
dp[i][state] 表示第i天处于state状态时的最大利润,:
state = buy(持有状态), sell（表示不持有状态, 包含卖出和冷冻期）

执行用时 :4 ms, 在所有 Go 提交中击败了54.55%的用户
内存消耗 :2.6 MB, 在所有 Go 提交中击败了10.00%的用户
dp[i][buy] = max
 */

func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i ++ {
		dp[i] = make([]int, 2)
	}

	dp[0][buy] = - prices[0]
	dp[0][sell] = 0
	for i := 1; i < len(prices); i ++ {
		//buy,(-prices[i]表新买股票prices[i])
		dp[i][buy] = max(dp[i - 1][buy], - prices[i])
		if i >= 2 {
			dp[i][buy] =  max(dp[i - 2][sell] - prices[i], dp[i][buy])
		}

		//sell
		dp[i][sell] = dp[i - 1][sell]
		dp[i][sell] = max(dp[i - 1][sell], dp[i - 1][buy] + prices[i])
	}
	profit := 0
	for i := 0; i < len(prices); i ++ {
		profit = max(profit, dp[i][buy])
		profit = max(profit, dp[i][sell])
	}
	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

