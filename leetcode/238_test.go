package leetcode

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t * testing.T) {
	result := productExceptSelf([]int {1, 0})
	if !reflect.DeepEqual(result, []int{0, 1}){
		t.Error("fail")
	}
}