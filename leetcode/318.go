package leetcode

/*
方法一：比较直观的方法，统计每个字符串字符的个数。判断两个字符串是否相同的时候直接迭代26个字母计数即可。
复杂度O（N * N）+ O（N * M），N是单词的个数，M是单词的长度


执行用时 :56 ms, 在所有 Go 提交中击败了52.38%的用户
内存消耗 :6.4 MB, 在所有 Go 提交中击败了27.27%的用户

type stat struct {
	count[]int
}

func maxProduct(words []string) int {
	stats := make([]stat, len(words))
	for i := 0; i < len(stats); i ++ {
		stats[i].count = make([]int, 26)
	}
	for i := 0; i < len(words); i ++ {
		for _, v := range words[i] {
			stats[i].count[v - 'a'] ++
		}
	}

	result := 0
	for i := 1; i < len(words); i ++ {
		for j := 0; j < i; j ++ {
			if !hasSame(stats[i], stats[j]) && len(words[i]) * len(words[j]) > result {
				result = len(words[i]) * len(words[j])
			}
		}
	}
	return result
}

func hasSame(s1, s2 stat) bool {
	for i := 0; i < 26; i ++ {
		if s1.count[i] > 0 && s2.count[i] > 0 {
			return true
		}
	}
	return false
}
 */


/*
方法2：
跟1差不多，只不过状态不再用26的数组，而是用一个整数来表示
用时 :20 ms, 在所有 Go 提交中击败了90.48%的用户
内存消耗 :6.1 MB, 在所有 Go 提交中击败了63.64%的用户
 */


type stat int

func maxProduct(words []string) int {
	stats := make([]stat, len(words))
	for i := 0; i < len(words); i ++ {
		for _, v := range words[i] {
			stats[i] |= 1 << uint(v - 'a')
		}
	}

	result := 0
	for i := 1; i < len(words); i ++ {
		for j := 0; j < i; j ++ {
			if (stats[i] & stats[j]) == 0 && len(words[i]) * len(words[j]) > result {
				result = len(words[i]) * len(words[j])
			}
		}
	}
	return result
}

