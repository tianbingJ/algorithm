package leetcode

import "testing"

func TestGetMoneyAmount(t *testing.T) {
	inputs := [] int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100}
	wants :=  [] int {0, 1, 2, 4, 6, 8, 10, 12, 14, 16, 400}
	for i := 0; i < len(inputs); i ++ {
		result := getMoneyAmount(inputs[i])
		if result != wants[i] {
			t.Errorf("input:%d %d result:%d\n", inputs[i], wants[i], result)
		}
	}
}