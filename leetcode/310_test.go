package leetcode

import (
	"reflect"
	"testing"
)

func TestFindMinHeightTrees(t *testing.T) {

	inputs := [][][]int{
		{
			[]int{0, 1},
			[]int{0, 2},
		},
		{
			[]int{0, 3},
			[]int{1, 3},
			[]int{2, 3},
			[]int{4, 3},
			[]int{5, 4},
		},
		{
			[]int{1, 0},
			[]int{1, 2},
			[]int{1, 3},
		},
	}

	nodes := []int{3, 6, 4}

	wants := [][]int{
		[]int{0},
		[]int{3, 4},
		[]int{1},
	}

	for i := 0; i < len(inputs); i++ {
		result := findMinHeightTrees(nodes[i], inputs[i])
		if !reflect.DeepEqual(result, wants[i]) {
			t.Errorf("inputs: %v result: %v wants :%v\n", inputs[i], result, wants[i])
		}
	}
}
