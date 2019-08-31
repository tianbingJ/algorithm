package leetcode

import "testing"

func TestIntegerBreak(t *testing.T) {
	inputs := []int{
		8,
		9,
		7,
		12,
		5,
		6,
		11,
		10,
		2,
		20,
		50,
	}

	wants := [] int {
		18,
		27,
		12,
		81,
		6,
		9,
		54,
		36,
		1,
		1458,
		86093442,
	}
	for i := 0; i < len(inputs); i ++ {
		result := integerBreak(inputs[i])
		if result != wants[i] {
			t.Errorf("input:%v result: %v want:%v\n", inputs[i], result, wants[i])
		}
	}
}
