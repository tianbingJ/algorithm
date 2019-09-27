package leetcode

import "testing"

func TestLongestPalindromeSubseq(t *testing.T) {
	inputs := []string{"bbbab", "abcdcba", "abcaaa"}
	wants := []int{4, 7, 4}
	for i := 0; i < len(inputs); i ++ {
		result := longestPalindromeSubseq(inputs[i])
		if result != wants[i] {
			t.Errorf("input: %v, wants: %v, result: %v\n", inputs[i], wants[i], result)
		}
	}

}
