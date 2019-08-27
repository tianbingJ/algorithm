package leetcode

import (
	"reflect"
	"testing"
)

func TestCountSmaller(t *testing.T) {

	inputs := [][]int{
		{0, 2, 1},
		{5,2,6,1},
		{5,4,3,2, 1},
		{5,4,4,4,1},
		{1,2},
		{2,1},
		{1},
		{1,2,3,4,5,4,3,2,1},
	}
	wants := [][]int{
		{0, 1, 0},
		{2,1,1,0},
		{4,3,2,1, 0},
		{4,1,1,1, 0},
		{0, 0},
		{1, 0},
		{0},
		{0,1,2,3,4,3,2,1,0},
	}
	for i := 0; i < len(inputs); i ++ {
		result := countSmaller(inputs[i])
		if !reflect.DeepEqual(wants[i], result) {
			t.Errorf("input: %v  wants: %v, but got result: %v\n", inputs[i], wants[i], result)
		}
	}
}