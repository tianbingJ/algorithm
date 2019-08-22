package leetcode

import "testing"

func TestFindDupliate(t * testing.T) {

	input := []int{1,3,4,2,2}
	result := findDuplicate(input)
	if result != 2 {
		t.Error("ERROR")
	}
}