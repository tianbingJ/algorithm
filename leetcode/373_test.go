package leetcode

import (
	"reflect"
	"testing"
)

func TestKSmallestPairs(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
	}

	wants := [][]int{
		{1, 1},
		{1, 2},
		{2, 1},
		{1, 3},
		{2, 2},
		{3, 1},
		{1, 4},
		{2, 3},
		{3, 2},
		{4, 1},
	}

	result := kSmallestPairs(inputs[0], inputs[1], 10)
	if !reflect.DeepEqual(wants, result) {
		t.Errorf("result: %v  want: %v\n", result, wants)
	}
}
