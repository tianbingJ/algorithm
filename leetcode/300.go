package leetcode

/**
动态规划
执行用时 :8 ms, 在所有 Go 提交中击败了67.72%的用户
内存消耗 :2.4 MB, 在所有 Go 提交中击败了52.38%的用户
 */
func lengthOfLIS1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	result := 1
	dp[0] = 1

	for i := 1; i < len(nums); i ++ {
		maxI := 1
		for j := 0; j < i; j ++ {
			if nums[i] > nums[j] && dp[j] >= maxI {
				maxI = dp[j] + 1
			}
		}
		dp[i] = maxI
		if maxI > result {
			result = maxI
		}
	}
	return result
}


/**
 NlgN的解法
非常有技巧的解法
参考 https://leetcode.com/problems/longest-increasing-subsequence/discuss/74824/JavaPython-Binary-search-O(nlogn)-time-with-explanation
 */

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tails := make([]int, len(nums))
	size := 0
	for _, num := range nums {
		i := 0
		j := size
		for ; i != j; {
			m := (i + j) / 2
			if tails[m] < num {
				i = m + 1
			} else {
				j = m
			}
		}
		tails[i] = num
		if i == size {
			size ++
		}
	}
	return size
}
