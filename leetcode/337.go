package leetcode

const (
	Rob    = 0
	NotRob = 1
)

/**
动态规划：
对于每个节点考虑抢劫和不抢劫的情况，同时把结果存在下来,效率有点低.
执行用时 :16 ms, 在所有 Go 提交中击败了47.17%的用户
内存消耗 :7.7 MB, 在所有 Go 提交中击败了11.11%的用户
 */
func rob(root *TreeNode) int {

	if root == nil {
		return 0
	}
	var dp = make(map[*TreeNode][]int)
	doRob(root, dp)
	return max(dp[root][Rob], dp[root][NotRob])
}

func doRob(root *TreeNode, dp map[*TreeNode][]int) int {
	if root == nil {
		return 0
	}
	if _, ok := dp[root]; ok {
		return max(dp[root][Rob], dp[root][NotRob])
	}
	dp[root] = make([]int, 2)
	//not rob root
	dp[root][NotRob] = doRob(root.Left, dp) + doRob(root.Right, dp)

	//rob root
	sum := root.Val
	if root.Left != nil {
		sum += doRob(root.Left.Left, dp) + doRob(root.Left.Right, dp)
	}
	if root.Right != nil {
		sum += doRob(root.Right.Left, dp) + doRob(root.Right.Right, dp)
	}
	dp[root][Rob] = sum
	return max(dp[root][Rob], dp[root][NotRob])
}
