package leetcode

func singleNumber(nums []int) []int {
	xorSum := 0
	for _, v := range nums{
		xorSum ^= v
	}

	lowestNum := lowestBitNum(xorSum)
	num1 := 0
	num2 := 0
	for _, v := range nums{
		if (v & lowestNum) != 0 {
			num1 ^= v
		} else {
			num2 ^= v
		}
	}
	return []int{num1, num2}
}

func lowestBitNum(num int) int {
	return num & (^num + 1)
}