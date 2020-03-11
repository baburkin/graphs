package graphs

type ccGraph struct {
	graph   Graph
	ccIndex []int
	ccSize  []int
	marked  []bool
	count   int
}

// ConnectedComponents searches the given graph for connected components
// and returns arrays c []int and s []int, where the value c[v] is an id
// of a component, to which the given vertex 'v' belongs, and s[i] is the size
// of component with id 'i'.
func ConnectedComponents(g Graph) ([]int, []int) {
	cc := new(ccGraph)
	switch g.(type) {
	case *digraph:
		cc.graph = InitGraphFromGraph(g)
	default:
		cc.graph = g
	}
	cc.ccIndex = make([]int, g.VNum(), g.VNum())
	cc.ccSize = make([]int, g.VNum(), g.VNum())
	cc.marked = make([]bool, g.VNum(), g.VNum())
	cc.count = 0
	for v := 0; v < g.VNum(); v++ {
		if !cc.marked[v] {
			runDFS(cc, v)
			cc.count++
		}
	}
	return cc.ccIndex, cc.ccSize
}

func runDFS(cc *ccGraph, v int) {
	cc.marked[v] = true
	cc.ccIndex[v] = cc.count
	cc.ccSize[cc.count]++
	for _, w := range cc.graph.Edges(v) {
		if !cc.marked[w] {
			runDFS(cc, w)
		}
	}
}
