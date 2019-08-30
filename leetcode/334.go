package leetcode

import "fmt"

/**
先实现比较笨的方法，得到用于比较的运行时间
超时了
func increasingTriplet(nums []int) bool {

	for i := 0; i < len(nums); i ++ {
		for j := i + 1; j < len(nums); j ++ {
			for  k := j + 1; k < len(nums); k ++ {
				if nums[k] > nums[j]  && nums[j] > nums[i] {
					return true
				}
			}
		}
	}
	return false
}
*/

/*
错了两次，发现题目理解错了。看成了 加1递增。
执行用时 :4 ms, 在所有 Go 提交中击败了92.86%的用户
内存消耗 :3.1 MB, 在所有 Go 提交中击败了41.18%的用户

想到用partition sort 求最长递增子序列的思想。
又错了一次的case：1, 2, -10 , -8, -7

思想：
1.先考虑判断是否存在两个递增子序列的情况
遍历数组维护一个已经存在的最小值，如果当前节点大于最小值，则存在两个递增子序列。
2.判断是否存在三个递增子序列的情况类似，只不过保存的状态信息多一些，条件判断复杂一点：
(1)维护一个最小值oneMin，遇到更小的数字则更新它。
(2)维护一个bool变量existTwo，表示是否存在两个的自增子序列[twoFirst, twoSecond]。
(3)遍历数组，对于当前值nums[i]判断：
(4)如果存在两个递增的子序列:
i)nums[i]大于 twoSecond，则构成了三个递增的子序列 return true
ii)oneMin < nums[i] < twoSenond,：
   则更新twoFirst = oneMin, twoSecond = nums[i]
   这一步非常重要，考虑 case 1，2， -10， -8， -7，(twoFirst, twoSecond)由（1，2）更新成(-10, -8)
iii)如果 twoFirst <nums[i] < twoSecond,则说明由nums[i]更新twoSecond可以得到一个更小的二元递增组。
(5)不存在两个递增的子序列：
i)如果nums[i] > oneMin， 则(oneMin,nums[i])构成一个两个递增的子序列。
(6)	if nums[i] < oneMin { oneMin = nums[i]}


*/
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	oneMin := nums[0]
	existTwo := false
	var twoFirst, twoSecond int
	for i := 1; i < len(nums); i++ {
		//前面存在两个递增的子序列
		if existTwo {
			if nums[i] > twoSecond {
				fmt.Printf("one: %d two: %d three %d\n", twoFirst, twoSecond, nums[i])
				return true
			}
			if nums[i] > oneMin {
				twoFirst = oneMin
				twoSecond = nums[i]
			} else if nums[i] > twoFirst {
				twoSecond = nums[i]
			}
		} else {
			if nums[i] > oneMin {
				existTwo = true
				twoFirst = oneMin
				twoSecond = nums[i]
			}
		}
		if nums[i] < oneMin {
			oneMin = nums[i]
		}
	}
	return false
}

/*
时间复杂度满足，但是实现比较复杂。
思想类似，但是有更简单的实现，这种想法太精巧了,但是没法得到正确的3个数字

        for (int n : nums) {
            if (n <= one) {
                one = n;
            } else if (n <= two) {
                two = n;
            } else {
                return true;
            }
        }
*/
