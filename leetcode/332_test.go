package leetcode

import (
	"reflect"
	"testing"
)

func TestFindItinerary(t *testing.T) {
	inputs := [][][]string{
		{
			{"MUC", "LHR"},
			{"JFK", "MUC"},
			{"SFO", "SJC"},
			{"LHR", "SFO"},
		},
		{
			{"JFK", "SFO"},
			{"JFK", "ATL"},
			{"SFO", "ATL"},
			{"ATL", "JFK"},
			{"ATL", "SFO"},
		},
	}
	wants := [][]string{
		{"JFK", "MUC", "LHR", "SFO", "SJC"},
		{"JFK","ATL","JFK","SFO","ATL","SFO"},
	}

	for i := 0; i < len(inputs); i ++ {
		result := findItinerary(inputs[i])
		if !reflect.DeepEqual(result, wants[i]) {
			t.Errorf("input :%v  result:%v  want:%v\n", inputs[i], result, wants[i])
		}
	}
}
