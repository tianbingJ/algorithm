package leetcode

import "testing"

func TestValidUtf8(t *testing.T) {

	inputs := [][]int{
		{145},
		{197, 130, 1},
		{235, 140, 4},
	}

	wants := []bool{
		false,
		true,
		false,
	}
	for i := 0; i < len(inputs); i ++ {
		result := validUtf8(inputs[i])
		if result != wants[i] {
			t.Errorf("input:%v %v result:%v\n", inputs[i], wants[i], result)
		}
	}
}
