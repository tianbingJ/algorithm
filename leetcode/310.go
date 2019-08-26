package leetcode

/*
方法1:
分别从每个顶点进行宽度优先搜索, 最大距离最近的顶点即为最小高度:
最后一个case超时了，看测试数据，边的个数有10001条
时间复杂度是O N * (M),如果M接近于N * N 的话，整体的时间复杂度是O(N^3)。

优化：
无环 N个顶点有N - 1条边

type graph struct {
	nodes      int
	edges      [][]int
	minheights []int
}

func (g *graph) solve() [] int {
	minDis := g.nodes
	for i := 0; i < g.nodes; i++ {

		g.minheights[i] = g.calHeight(i)
		if g.minheights[i] < minDis {
			minDis = g.minheights[i]
		}
	}

	result := []int{}
	for i := 0; i < g.nodes; i++ {
		if g.minheights[i] == minDis {
			result = append(result, i)
		}
	}
	return result
}

func (g *graph) calHeight(startIndex int) int {
	queue := []int{}
	queue = append(queue, startIndex)
	dis := make([]int, g.nodes)
	dis[startIndex] = 0
	height := 0
	for ; len(queue) > 0; {
		front := queue[0]
		queue = queue[1:]
		for i := 0; i < len(g.edges[front]); i++ {
			to := g.edges[front][i]
			if to == startIndex {
				continue
			}
			if dis[to] > 0 {
				continue
			}
			dis[to] = dis[front] + 1
			if dis[to] > height {
				height = dis[to]
			}
			queue = append(queue, to)
		}
	}
	return height
}

func findMinHeightTrees(n int, edges [][]int) []int {
	edgesAdj := make([][]int, n)
	for i := 0; i < len(edges); i++ {
		from := edges[i][0]
		to := edges[i][1]
		edgesAdj[from] = append(edgesAdj[from], to)
		edgesAdj[to] = append(edgesAdj[to], from)
	}

	graph := &graph{
		nodes:      n,
		edges:      edgesAdj,
		minheights: make([]int, n),
	}
	return graph.solve()
}

*/

func findMinHeightTrees(n int, edges [][]int) []int {

	if n ==  1 {
		return  []int {0}
	}
	if n ==  2 {
		return  []int {0, 1}
	}

	edgesAdj := make([][]int, n)
	degrees := make([]int, n)
	for i := 0; i < len(edges); i++ {
		from := edges[i][0]
		to := edges[i][1]
		edgesAdj[from] = append(edgesAdj[from], to)
		edgesAdj[to] = append(edgesAdj[to], from)
		degrees[from] ++
		degrees[to] ++
	}

	queue := []int{}
	for i := 0; i < n; i ++ {
		if degrees[i] == 1 {
			queue = append(queue, i)
		}
	}

	totalSize := n
	for ; totalSize > 2; {
		size := len(queue)
		for j := 0; j < size; j ++ {
			front := queue[j]
			for i := 0; i < len(edgesAdj[front]); i ++ {
				node := edgesAdj[front][i]
				degrees[node] --
				if degrees[node] == 1 {
					queue = append(queue, node)
				}
			}
		}
		queue = queue[size:]
		totalSize -= size
	}
	return queue
}
