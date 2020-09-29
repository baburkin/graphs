package graphs

import (
	"math"
)

func prependInt(x []int, y int) []int {
	x = append(x, 0)
	copy(x[1:], x)
	x[0] = y
	return x
}

// ShortestPathBFS returns the number of step that form
// the shortest path between `vSource` and `vTarget` vertices:
//
// vSource -> v1 -> v2 -> ... -> vTarget (number of arrows/hops)
//
// If vSource == vTarget then zero-length slice is returned
func ShortestPathBFS(g Graph, vSource int, vTarget int) int {
	var queue []int
	var marked []bool
	var distTo []int
	var edgeTo []int

	marked = make([]bool, g.VNum(), g.VNum())
	distTo = make([]int, g.VNum(), g.VNum())
	edgeTo = make([]int, g.VNum(), g.VNum())
	queue = make([]int, 0, g.VNum()-1)

	edgeTo[vSource] = vSource
	distTo[vSource] = 0
	marked[vSource] = true

	// Init distTo with "Infinity"
	for v := 0; v < g.VNum(); v++ {
		if v != vSource {
			distTo[v] = math.MaxInt64
		}
	}

	// BFS algorithm to search for shortest path
	queue = append(queue, vSource)
	for len(queue) != 0 {
		v := queue[0]
		queue = queue[1:]
		for _, w := range g.Edges(v) {
			if !marked[w] {
				edgeTo[w] = v
				distTo[w] = distTo[v] + 1
				marked[w] = true
				queue = append(queue, w)
			}
		}
	}

	// Path means succession of vertices (hops) between vSource and vTarget
	// for v := edgeTo[vTarget]; distTo[v] != 0; v = edgeTo[v] {
	//	path = prependInt(path, v)
	//}

	return distTo[vTarget]
}
