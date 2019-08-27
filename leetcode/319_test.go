package leetcode

import "testing"

func TestBulbSwitch(t *testing.T) {

	inputs := []int{99999}
	wants := []int{99999}
	for i := 0; i < len(inputs); i ++ {
		result := bulbSwitch(inputs[i])
		if result != wants[i] {
			t.Errorf("want %d  result: %d", wants[i], result)
		}
	}
}