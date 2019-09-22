package leetcode

import (
	"reflect"
	"testing"
)

func TestCalcEquation(t *testing.T) {

	equations := [][]string {
		{"a","b"},
		{"b","c"},
	}

	values := []float64{2.0, 3.0}

	queries := [][]string {
		{"a", "c"},
		{"b", "a"},
		{"a", "e"},
		{"a", "a"},
		{"x", "x"},
	}
	result := calcEquation(equations, values, queries)
	expect := []float64{6.0,0.5,-1.0,1.0,-1.0}
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("result:%v expect:%v\n", result, expect)
	}
}
