package leetcode

/**
预处理，计算出到每个位置i的各个字符的计数情况，dp[i][0~26].
对于任意两个位置[start, end]统计各个字符是否都大于k。是比较暴力的方案。
时间复杂度：
O(n * n)

执行用时 :512 ms, 在所有 Go 提交中击败了26.67%的用户
内存消耗 :4.5 MB, 在所有 Go 提交中击败了7.14%的用户
*/
func longestSubstring1(s string, k int) int {
	n := len(s)
	if n <= 0 {
		return 0
	}
	dp := make([][26]int, n)
	dp[0][s[0]-'a'] ++
	for i := 1; i < n; i++ {
		copy(dp[i][:], dp[i-1][:])
		dp[i][s[i]-'a'] ++
	}

	max := 0
	for i := k - 1; i < n; i++ {
		for j := 0; j <= i && j + max  - 1 < i; j++ {
			valid := true
			for kk := 0; kk < 26; kk++ {
				v := 0
				if j == 0 {
					v = dp[i][kk]
				} else {
					v = dp[i][kk] - dp[j - 1][kk]
				}
				if v > 0 && v < k {
					valid = false
					break
				}
			}
			if valid && i - j + 1 > max  {
				max = i - j + 1
			}
		}
	}
	return max
}


func longestSubstring(s string, k int) int {
	n := len(s)
	if n <= 0 {
		return 0
	}
	dp := make([][26]int, n)
	dp[0][s[0]-'a'] ++
	for i := 1; i < n; i++ {
		copy(dp[i][:], dp[i-1][:])
		dp[i][s[i]-'a'] ++
	}

	return doLongestSubstring(s, k, 0, n - 1)
}

func doLongestSubstring(s string, k int, start, end int) int {
	if end - start + 1 < k {
		return 0
	}
	cnt := make([]int, 26)
	for i := start; i <= end; i ++ {
		cnt[s[i] - 'a'] ++
	}

	allValid := true
	for i := start; i <= end; i ++ {
		if cnt[s[i] - 'a'] < k {
			allValid = false
		}
	}

	if allValid {
		return end - start + 1
	}

	for i := start; i <= end; i ++ {
		if cnt[s[i] - 'a'] < k {
			return  max(doLongestSubstring(s, k, start, i - 1), doLongestSubstring(s, k, i + 1, end))
		}
	}
	return end - start + 1
}
