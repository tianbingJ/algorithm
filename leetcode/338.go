package leetcode

/*
直接计算
执行用时 :2140 ms, 在所有 Go 提交中击败了65.22%的用户
内存消耗 :31.2 MB, 在所有 Go 提交中击败了9.09%的用户
func countBits(num int) []int {
	result := make([]int, num + 1)

	for i := 0; i <= num; i ++ {
		count := 0
		for n := i; n > 0; n = n & (n - 1) {
			count ++
		}
		result[i] = count
	}
	return result
}
*/

/*
执行用时 :1960 ms, 在所有 Go 提交中击败了93.91%的用户
内存消耗 :36.6 MB, 在所有 Go 提交中击败了6.06%的用户

找规律：
数字	二进制  	1的个数
0		0		0
1		1		1
2		10		1
3		11		2
4		100		1
5		101		2
6		110		2
7		111		3
8		1000	1
9		1001	2

2的幂次方都只有一个1，其余情况：
5比1多出一个前置1
6比2多出一个前置1
7比3多处一个前置1，一次类推。
对于数字n(i非2的幂)，最大x满足 2^x < n， 则n的一个个数 = n % (2 ^x) + 1
 */
func countBits(num int) []int {
	result := make([]int, num + 1)
	for i := 1; true ; i ++ {
		n := 1 << uint(i)
		if n > num {
			break
		}
		result[n] = 1
	}

	power := 1
	for i := 2; i <= num; i ++ {
		if result[i] == 1 {
			power =  power << 1
			continue
		}
		result[i] = result[i % power] + 1
	}
	return result
}

/**
实际上计算当前数的时候能用到前面计算的结果就行。

更简单的一种解法是：
x & (x - 1)会去掉最后一个1，这个关系也可以利用.
执行用时 :1976 ms, 在所有 Go 提交中击败了92.17%的用户
 */
