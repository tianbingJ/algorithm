package leetcode

import "testing"

func TestCanMeasureWater(t *testing.T) {
	inputs := [][]int {
		{3, 3, 6},
		{22003, 31237, 0},
		{3, 5, 4},
		{2, 6, 5},
		{1, 100, 100},
		{3, 3, 5},
		{8, 8, 8},
	}
	wants := []bool {
		true,
		true,
		true,
		false,
		true,
		false,
		true,
	}
	for i := 0; i < len(inputs); i ++ {
		result := canMeasureWater(inputs[i][0], inputs[i][1], inputs[i][2])
		if result != wants[i] {
			t.Fatalf("input:%v result: %v want:%v\n", inputs[i], result, wants[i])
		}
	}
}
