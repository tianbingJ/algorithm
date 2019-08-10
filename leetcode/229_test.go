package leetcode

import (
	"reflect"
	"testing"
)

func TestMajorityElements(t * testing.T) {

	var inputs [][]int = [][]int{
		[]int{1, 3, 3, 4},
		[]int{6,5,5},
		[]int{1, 1, 2, 2, 3, 3, 2},
		[]int{1, 1, 2, 2, 3, 3, 2, 1},
		[]int{1, 1, 3},
		[]int{1, 2, 3},
		[]int{1, 1, 1},
		[]int{1},
		[]int{1, 1, 2, 2, 3, 3},
	}
	var wants [][]int = [][]int {
		[]int{3},
		[]int{5},
		[]int{2},
		[]int{1,2},
		[]int{1},
		[]int{},
		[]int{1},
		[]int{1},
		[]int{},
	}

	for i := 0; i < len(inputs); i ++ {
		result := majorityElement(inputs[i])
		if !reflect.DeepEqual(result, wants[i]) {
			t.Errorf("input: %v  want: %v result:%v\n", inputs[i], wants[i], result)
		}
	}
}

