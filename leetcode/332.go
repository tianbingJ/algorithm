package leetcode

import (
	"sort"
)

type FlightSlice [][]string

func (slice FlightSlice)Len() int {
	return len(slice)
}

func (slice FlightSlice) Less(i, j int) bool {
	if slice[i][0] == slice[j][0] {
		return slice[i][1] < slice[j][1]
	}
	return slice[i][0] < slice[j][0];
}

func (slice FlightSlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

var stack []string
var visitAdj map[string][]bool
var adj map[string][]string

func findItinerary(tickets [][]string) []string {
	start := "JFK"
	stack = make([]string, 0)
	stack = append(stack, start)

	sort.Sort(FlightSlice(tickets))
	adj = make(map[string][]string)

	for i := 0; i < len(tickets); i ++ {
		from := tickets[i][0]
		to := tickets[i][1]
		adj[from] = append(adj[from], to)
	}

	//init visitAdj
	visitAdj = make(map[string][]bool)
	for v := range adj {
		visited := make([]bool, len(adj[v]))
		visitAdj[v] = visited
	}
	dfs(start, tickets)
	return stack
}

func dfs(cur string, tickets [][]string) bool {
	if len(stack) == len(tickets) + 1{
		return true
	}

	for i := 0; i < len(adj[cur]); i ++ {
		if !visitAdj[cur][i] {
			//use edge here , flag and push stack
			stack = append(stack, adj[cur][i])
			visitAdj[cur][i] = true
			exist := dfs(adj[cur][i], tickets)
			if exist {
				return true
			}
			//edge not work, flag and pop stack
			visitAdj[cur][i] = false
			stack = stack[0: len(stack) - 1]
		}
	}
	return false
}
