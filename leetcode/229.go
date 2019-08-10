package leetcode

/*
给定一个大小为 n 的数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1)。

超过1/3的元素最多有2个.

思路：
更改求众数的摩尔投票法
1.初版的方案是遇到一个数新的数计数加2，遇到不是当前数的计数-1，如果一个数字出现次数大于n/3,n/3 * 2（当前计数） - 剩下的n / 3 * 2，总体上是大于等于0的。
对于[6，5，5]这样的case，正序过不了，可以逆序再来一次。
但是对于[1,3,3,4]这样的case就无能无力了，这样的case得先遇到3才能过。
66个case过了64个，有两个过不了。由于已经比较麻烦了，意识到实现上有更好的方法。

2.维护两个2个众数的计数：
(1).如果2个坑没占满，则计数
(2).遇到跟某个数相同的 +1，
(3).遇到跟2个数都不相同的情况计数同时-1。最后需要判断这两个数是不是众数。
*/

func majorityElement(nums []int) []int {
	num1Count := 0
	num1 := 0
	num2Count := 0
	num2 := 0

	for i := 0; i < len(nums); i ++ {
		curNum := nums[i]

		if curNum == num1 {
			num1Count ++
			continue
		}
		if curNum == num2 {
			num2Count ++
			continue
		}
		if num1Count == 0 {
			num1 = curNum
			num1Count ++
			continue
		}
		if num2Count == 0 {
			num2 = curNum
			num2Count ++
			continue
		}
		num1Count --
		num2Count --
	}
	result := []int{}
	if isMajority(nums, num1) {
		result = append(result, num1)
	}
	if num1 != num2 && isMajority(nums, num2) {
		result = append(result, num2)
	}
	return result
}

func isMajority(nums [] int, value int) bool {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == value {
			count++
		}
	}
	return count > len(nums)/3
}