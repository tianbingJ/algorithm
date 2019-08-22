package leetcode

import (
	"github.com/gpmgo/gopm/modules/log"
	"testing"
)

func TestNumSquares(t *testing.T) {
	input := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	want :=  []int{1,2,3,1,2,3,4,2,1,2,  3, 3, 2, 3, 4, 1, 2, 2, 3, 2}

	for i := 0; i < len(input); i ++ {
		result := numSquares(input[i])
		if result != want[i] {
			log.Error("input:%d want: %d result:%d\n", input[i], want[i], result)
		}
	}
}
