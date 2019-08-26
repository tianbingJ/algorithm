package leetcode

import "testing"

func TestNthSuperUglyNumber(t *testing.T) {
	result := nthSuperUglyNumber(12, []int{2, 7, 13, 19})
	if result != 32 {
		t.Errorf("result :%d", result)
	}
}