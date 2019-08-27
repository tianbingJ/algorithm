package leetcode

import "math"

/**
数学不好，遇到这种题都有点怕，先用直观的方式，分别考虑N个灯中每个灯被翻转的次数。
O(N^2)
不出意外，超时了

func bulbSwitch(n int) int {
	totalCount := 0
	for bulb := 1; bulb <= n; bulb++ {
		count := 0
		for i := 1; i <= bulb; i++ {
			if bulb % i == 0 {
				count ++
			}
		}
		if count % 2 == 1 {
			totalCount ++
		}
	}
	return totalCount
}
*/

/**
不断试case 发现是根号规律...
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2 MB, 在所有 Go 提交中击败了12.50%的用户

在因数的情况下灯才会翻转，除了完全平方数，因数全是成对出现。
对于n，小于n的完全平方数的个数为根号n，即完全平方数为 1^2.2^2, 3^2..sqrt(n)^2
 */
func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}