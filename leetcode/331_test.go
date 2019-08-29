package leetcode

import "testing"

func TestIsValidSerialization(t *testing.T) {
	inputs := []string{
		"9,3,4,#,#,1,#,#,2,#,6,#,#",
		"1,#",
		"9,#,#,1",
	}
	wants := []bool {
		true,
		false,
		false,
	}

	for i := 0; i < len(inputs); i ++ {
		result := isValidSerialization(inputs[i])
		if result != wants[i] {
			t.Errorf("input:%v result:%v want:%v\n", inputs[i], result, wants[i])
		}
	}
}