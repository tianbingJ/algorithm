package leetcode

import (
	"testing"
)

func TestNumArray(t *testing.T) {
	inputs := []int {1, 3, 5}
	numArray := Constructor(inputs)
	result := numArray.SumRange(0, 2)
	if result != 9 {
		t.Errorf("error")
	}
	numArray.Update(1, 2)
	result = numArray.SumRange(0, 2)
	if result != 8 {
		t.Errorf("error result %d\n", result)
	}
}
