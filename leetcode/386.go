package leetcode

import "fmt"

/**
执行用时 :2440 ms, 在所有 Go 提交中击败了94.29%的用户
内存消耗 :27.3 MB, 在所有 Go 提交中击败了25.00%的用户
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

/**
每次方法调用填充完相同位数的数字
*/
func iterLexOrder(curNum int, order * LexOrder) {
	if curNum > order.n  {
		return
	}
	for i := 0; i <= 9 && curNum + i <= order.n && order.index < order.n; i ++ {
		fmt.Println(curNum + i, order.index)
		order.result[order.index] = curNum + i
		order.index ++
		iterLexOrder((curNum + i) * 10, order)
	}
}

