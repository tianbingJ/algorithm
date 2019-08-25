package leetcode

import "testing"

func TestMaxProfit(t *testing.T) {
	inputs := [][]int {
		[]int{2, 1, 4},
		[]int{1,2,3,0,2},
		[]int{1,2,3,4,5},
		[]int{5,4,3,2,1},
		[]int{4,5,1,2,3},
		[]int{1, 3},
	}
	wants := []int{
		3,
		3,
		4,
		0,
		2,
		2,
	}

	for i := 0; i < len(inputs); i ++ {
		result :=  maxProfit(inputs[i])
		if result != wants[i] {
			t.Errorf("inputs:%v  result:%d wants:%d\n", inputs[i], result, wants[i])
		}
	}
}
