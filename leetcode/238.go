package leetcode

/*
1.比较笨的方法，并且不符合题意
执行用时 :3812 ms, 在所有 Go 提交中击败了87.74%的用户
内存消耗 :30.9 MB, 在所有 Go 提交中击败了10.53%的用户

2.题目要求不能用除法
两个数组，一个从左到右算前n项的乘积leftProduction[n]；另外一个从右到左算前n项的乘积rightProduction[n]；
对于第k个结果 = leftProduction[k-1] * rightProduction[k + 1]
竟然超时了
这个可以优化，只用一个数组就可以了，中间结果保存在原数组里。

3.分析：
由于都是整数，数量很大，并且结果不会溢出，可得出结论，里面的数字大部分是1，-1或者有0，否则很容易溢出。
统计0的个数，1的个数，-1的个数。可以把这几个结果忽略掉。其他的数字就比较容易处理了，由于这样的数字最多也就30多个，空间复杂度是常数.
行用时 :3744 ms, 在所有 Go 提交中击败了91.51%的用户
*/
func productExceptSelf1(nums []int) []int {
	production := 1
	zeroCount := 0
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCount++
		} else {
			production *= nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if zeroCount > 1 {
			result[i] = 0
		} else if zeroCount == 1 {
			if nums[i] == 0 {
				result[i] = production
			} else {
				result[i] = 0
			}
		} else {
			result[i] = production / nums[i]

		}
	}
	return result
}

func productExceptSelf2(nums []int) []int {
	leftProduction := make([]int, len(nums))
	leftProduction[0] = nums[0]
	rightProduction := make([]int, len(nums))
	rightProduction[len(nums)-1] = nums[len(nums)-1];
	for i := 1; i < len(nums); i++ {
		leftProduction[i] = leftProduction[i-1] * nums[i];
	}
	for i := len(nums) - 2; i >= 0; i-- {
		rightProduction[i] = rightProduction[i+1] * nums[i];
	}
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result[i] = rightProduction[1];
		} else if i == len(nums)-1 {
			result[i] = leftProduction[len(nums)-2]
		} else {
			result[i] = leftProduction[i-1] * rightProduction[i+1];
		}
	}
	return result
}

func productExceptSelf(nums []int) []int {
	count0, count1, countN1 := 0, 0, 0
	other := make(map[int]int)
	otherProduction := 1
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		switch nums[i] {
		case 1:
			count1++
		case -1:
			countN1++
		case 0:
			count0++
		default:
			other[i] = nums[i]
			otherProduction *= nums[i]
		}
	}

	//1 和 -1 的乘积
	oneProduction := 1
	if countN1 %2 == 1 {
		oneProduction = -1
	}
	if count0 > 1 {
		for i := 0; i < len(nums);i ++ {
			result[i] = 0
		}
		return result
	}

	if count0 == 1 {
		for i := 0; i < len(nums);i ++ {
			if nums[i] == 0 {
				result[i] = otherProduction * oneProduction
			} else {
				result[i] = 0
			}
		}
		return result
	}

	for i := 0; i < len(nums);i ++ {
		if nums[i] == 1 {
			if oneProduction > 0 {
				result[i] = otherProduction
			} else{
				result[i] = - otherProduction
			}
		} else if nums[i] == -1 {
			if oneProduction > 0 {
				result[i] = - otherProduction
			} else{
				result[i] =  otherProduction
			}
		} else {
			result[i] = productionWithoutIndex(other, i) * oneProduction
		}
	}
	return result
}

func productionWithoutIndex(mp map[int]int, index int) int {
	result := 1
	for k, v :=  range mp {
		if k != index {
			result *= v
		}
	}
	return result
}

