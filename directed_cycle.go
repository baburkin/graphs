package graphs

type directedCycle struct {
	DirectedGraph
	marked  []bool
	onStack []bool
	edgeTo  []int
	cycle   []int
}

func initDFSCycle(g DirectedGraph) *directedCycle {
	dc := new(directedCycle)
	dc.DirectedGraph = g
	dc.marked = make([]bool, g.VNum(), g.VNum())
	dc.onStack = make([]bool, g.VNum(), g.VNum())
	dc.edgeTo = make([]int, g.VNum(), g.VNum())
	return dc
}

func dfsCycle(dc *directedCycle, v int) {
	dc.onStack[v] = true
	dc.marked[v] = true
	for _, w := range dc.Edges(v) {

		// if there is a cycle - return immediately
		if len(dc.cycle) > 0 {
			return
		} else if !dc.marked[w] { // do DFS when found unvisited vertex
			dc.edgeTo[w] = v
			dfsCycle(dc, w)
		} else if dc.onStack[w] { // push the cycle to stack if w is on the stack
			dc.cycle = make([]int, 0)
			for x := v; x != w; x = dc.edgeTo[x] {
				dc.cycle = append(dc.cycle, x)
			}
			dc.cycle = append(dc.cycle, w)
			dc.cycle = append(dc.cycle, v)
		}

	}
}

// IsDAG determines if the directed graph is a DAG (directed acyclic graph)
func IsDAG(g DirectedGraph) bool {
	dc := initDFSCycle(g)

	return true
}

// Rank shows the rank of the vertex if topological sort exists, -1 otherwise
func Rank(g DirectedGraph, v Vertex) int {
	return -1
}
