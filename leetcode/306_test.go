package leetcode

import (
	"testing"
)

func TestIsAdditiveNumber(t *testing.T) {
	inputs := []string{
		"010",
		"011",
		"000",
		"112358",
		"199100199",
		"0123123",
		"01231234",
	}
	wants := []bool{
		false,
		true,
		true,
		true,
		true,
		true,
		false,
	}

	result := isAdditive("112358", 0, 1)
	if result != true {
		t.Fatal("fail")
	}

	for i := 0; i < len(inputs); i ++ {
		if isAdditiveNumber(inputs[i]) != wants[i] {
			t.Errorf("inputs: %s\n  wants: %v", inputs[i], wants[i])
		}
	}
}
