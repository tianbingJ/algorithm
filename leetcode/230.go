package leetcode

/*
思路
1.简单来说可以中序遍历，直接返回列表中第k个元素。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	_, _, value := doKthSmallest(root, k, 0)
	return value
}

func doKthSmallest(root * TreeNode, k int, count int) (leftCount int, found bool , value int) {
	if root == nil {
		return count, false, 0
	}
	count, found, value = doKthSmallest(root.Left, k, count)
	if found {
		return leftCount, found, value
	}
	if count + 1 == k {
		return count +1, true, root.Val
	}
	return doKthSmallest(root.Right, k, count + 1)
}
