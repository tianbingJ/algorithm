package leetcode

import "testing"

func TestCalculate(t * testing.T) {
	inputs := []string{
		"42",
		"3+2*2",
		" 3 + 2 * 2 ",
		"3",
		"3*2*2",
		"3*2-2",
		"3-2-2",
		" 3/2 ",
	}
	wants := []int{
		42,
		7,
		7,
		3,
		12,
		4,
		-1,
		1,
	}

	for i := 0; i < len(inputs); i ++ {
		result := calculate(inputs[i])
		if result != wants[i] {
			t.Errorf("input:%s want:%d  result:%d\n", inputs[i], wants[i], result)
		}
	}
}