package leetcode

/**
执行用时 :2440 ms, 在所有 Go 提交中击败了94.29%的用户
内存消耗 :27.3 MB, 在所有 Go 提交中击败了25.00%的用户

处理当前数字curNum，首先把它加入list中，再需要考虑以它开头的其他数字（(curNum + i)* 10）, i = 0, ..., 9
 */

type LexOrder struct  {
	n int
	index int
	result []int
}

func lexicalOrder(n int) []int {
	result := make([]int, n)
	lexOrder := &LexOrder{
		n: n ,
		index:0,
		result:result,
	}

	iterLexOrder(1, lexOrder)
	return lexOrder.result
}

func iterLexOrder(curNum int, order * LexOrder) {
	if curNum > order.n  {
		return
	}
	for i := 0; i <= 9 && curNum + i <= order.n && order.index < order.n; i ++ {
		order.result[order.index] = curNum + i
		order.index ++
		iterLexOrder((curNum + i) * 10, order)
	}
}

