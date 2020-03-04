package graphs

import "fmt"

type directedCycle struct {
	DirectedGraph
	marked  []bool // flag used to mark visited vertices
	onStack []bool // flag used to detect cycles
	edgeTo  []int  // the source edge that has path to the given edge
	cycle   []int  // stack that keeps a cycle (if exists) - pop the last element
}

func (dc *directedCycle) String() string {
	return fmt.Sprintf("Graph: [%v] marked: [%v], onStack: [%v] cycle: [%v]",
		dc.DirectedGraph, dc.marked, dc.onStack, dc.cycle)
}

func initDFSCycle(g DirectedGraph) *directedCycle {
	dc := new(directedCycle)
	dc.DirectedGraph = g
	dc.marked = make([]bool, g.VNum(), g.VNum())
	dc.onStack = make([]bool, g.VNum(), g.VNum())
	dc.edgeTo = make([]int, g.VNum(), g.VNum())
	dc.cycle = make([]int, 0, 0)
	return dc
}

func findCycleDFS(dc *directedCycle, v int) {
	dc.onStack[v] = true
	dc.marked[v] = true
	for _, w := range dc.Edges(v) {
		// if there is a cycle - return immediately
		if len(dc.cycle) > 0 {
			return
		} else if !dc.marked[w] { // do DFS when found unvisited vertex
			dc.edgeTo[w] = v
			findCycleDFS(dc, w)
		} else if dc.onStack[w] { // push the cycle to stack if w is on the stack
			dc.cycle = make([]int, 0)
			for x := v; x != w; x = dc.edgeTo[x] {
				dc.cycle = append(dc.cycle, x)
			}
			dc.cycle = append(dc.cycle, w)
			dc.cycle = append(dc.cycle, v)
		}
	}
	dc.onStack[v] = false
}

// IsDAG determines if the directed graph is a DAG (directed acyclic graph)
func IsDAG(g DirectedGraph) bool {
	dc := initDFSCycle(g)
	for v := 0; v < dc.VNum(); v++ {
		if !dc.marked[v] && len(dc.cycle) == 0 {
			findCycleDFS(dc, v)
		}
	}
	return len(dc.cycle) == 0
}

// GetCycleInGraph returns the list of vertices that form a cycle in a directed graph
func GetCycleInGraph(g DirectedGraph) []int {
	dc := initDFSCycle(g)
	for v := 0; v < dc.VNum(); v++ {
		if !dc.marked[v] && len(dc.cycle) == 0 {
			findCycleDFS(dc, v)
		}
	}
	c := dc.cycle
	// reverse the order as dc.cycle is a stack
	for v := 0; v < len(c)/2; v++ {
		c[v], c[len(c)-v-1] = c[len(c)-v-1], c[v]
	}

	return dc.cycle
}
