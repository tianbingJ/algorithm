package leetcode

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	inputs := [][]int{
		{1,1,1,2,2,3},
	}

	wants := [][]int{
		{1,2},
	}
	for i := 0; i < len(inputs); i ++ {
		result := topKFrequent(inputs[i], 2)
		if !reflect.DeepEqual(result,  wants[i]){
			t.Errorf("input:%v result: %v want:%v\n", inputs[i], result, wants[i])
		}
	}
}