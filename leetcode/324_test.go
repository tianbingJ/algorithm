package leetcode

import (
	"reflect"
	"testing"
)

func TestWiggleSort(t *testing.T) {

	inputs := [][]int{
		{4,5,5,6},
		{1, 2, 4, 3, 5, 6},
	}

	wants := [][]int{
		{5,6,4,5},
		{3, 6 ,2 ,5, 1 ,4},
	}

	for i := 0; i < len(inputs); i ++ {
		input := make([]int, len(inputs[i]))
		copy(input, inputs[i])
		wiggleSort(inputs[i])
		if !reflect.DeepEqual(inputs[i], wants[i]) {
			t.Errorf("input:%v result:%v want:%v", input, inputs[i],  wants[i])
		}
	}
}
