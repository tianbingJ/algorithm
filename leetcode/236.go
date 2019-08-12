package leetcode

type Solution struct {
	visitP bool
	visitQ bool
	pathP  []*TreeNode
	pathQ  []*TreeNode
	path   []*TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	solution := &Solution{
	}
	iter(solution, root, p, q)
	i := 0
	for ; i < len(solution.pathP) && i < len(solution.pathQ); i++ {
		if solution.pathP[i] != solution.pathQ[i] {
			break
		}
	}
	return solution.pathP[i-1]
}

func iter(solution *Solution, root, p, q *TreeNode) {
	if root == nil {
		return
	}
	if solution.visitP && solution.visitQ {
		return
	}
	solution.path = append(solution.path, root)
	if root.Val == p.Val {
		solution.visitP = true
		solution.pathP = make([]*TreeNode, len(solution.path))
		copy(solution.pathP, solution.path)
	}
	if root.Val == q.Val {
		solution.visitQ = true
		solution.pathQ = make([]*TreeNode, len(solution.path))
		copy(solution.pathQ, solution.path)
	}
	iter(solution, root.Left, p, q)
	iter(solution, root.Right, p, q)
	solution.path = solution.path[0:len(solution.path) - 1]
}
