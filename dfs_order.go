package graphs

type dfsOrder struct {
	DirectedGraph
	pre       []int
	post      []int
	postOrder []int
	preOrder  []int
	marked    []bool
}

// initDFSOrder inits the data structure required to find DFS order
func initDFSOrder(g DirectedGraph) *dfsOrder {
	newG := new(dfsOrder)
	newG.DirectedGraph = g
	newG.pre = make([]int, g.VNum(), g.VNum())
	newG.post = make([]int, g.VNum(), g.VNum())
	newG.preOrder = make([]int, g.VNum(), g.VNum())
	newG.postOrder = make([]int, g.VNum(), g.VNum())
	newG.marked = make([]bool, g.VNum(), g.VNum())
	return newG
}

func findOrderDFS(g *dfsOrder, v int) {
	g.marked[v] = true
}
