package leetcode

import (
	"testing"
)

func TestPushDominoes(t * testing.T) {

	inputs := []string {
		"RR.L",
		".L.R...LR..L..",
	}
	wants := []string {
		"RR.L",
		"LL.RR.LLRRLL..",
	}
	for i := 0; i < len(inputs);i ++ {
		result := pushDominoes(inputs[i])
		if result != wants[i] {
			t.Errorf("input: %s, want: %s, result: %s\n", inputs[i], wants[i], result)
		}
	}
}
