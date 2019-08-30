package leetcode

import "testing"

func TestIncreasingTriplet(t *testing.T) {
	inputs := [][]int{
		{1, 2, -10, -8, -7},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{6, 1, 8, 7, 9, 2},
	}

	wants := []bool {
		true,
		true,
		false,
		true,
	}

	for i := 0; i < len(inputs); i ++ {
		result := increasingTriplet(inputs[i])
		if result != wants[i] {
			t.Errorf("input:%v result: %v want:%v\n", inputs[i], result, wants[i])
		}
	}
}
