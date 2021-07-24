package graphs

import "sort"

type ccGraph struct {
	graph   Graph   // Input graph
	ccIndex []int   // ccIndex[v] returns the component ID of a vertex v
	ccSize  []int   // ccSize[i] shows the size of i-th component
	marked  []bool  // internal data structure used for DFS
	count   int     // number of connected components
	comp    [][]int // comp[i] is a slice of indices that belong to i-th component
}

// ConnectedComponents interface provides the API for getting
// the number of connected components in a graph
type ConnectedComponents interface {
	// Graph returns the source graph
	Graph() Graph
	// Index returns the connected component index for the vertex v
	Index(v int) int
	// Component returns a sorted slice of vertices in the given connected component
	Component(compID int) []int
	// CompSize returns size of the given connected component
	CompSize(compID int) int
	// Count returns total number of connected components
	Count() int
	// AllIndices returns all connected components: AllIndices()[v] returns
	// connected component index for the vertex v
	AllIndices() []int
	// AllCompSizes returns connected component size: AllCompSizes()[i] returns
	// number of vertices in i-th connected component
	AllCompSizes() []int
	// AllComponents returns all connected components: AllComponents()[i] returns
	// a slice of vertices that belong to the i-th connected component
	AllComponents() [][]int
}

func (cc *ccGraph) Graph() Graph {
	return cc.graph
}

func (cc *ccGraph) Index(v int) int {
	return cc.ccIndex[v]
}

func (cc *ccGraph) Component(compID int) []int {
	return cc.comp[compID]
}

func (cc *ccGraph) CompSize(compID int) int {
	return cc.ccSize[compID]
}

func (cc *ccGraph) Count() int {
	return cc.count
}

func (cc *ccGraph) AllIndices() []int {
	return cc.ccIndex
}

func (cc *ccGraph) AllCompSizes() []int {
	return cc.ccSize
}

func (cc *ccGraph) AllComponents() [][]int {
	return cc.comp
}

// InitConnectedComponents searches the given graph for connected components
// and returns arrays c []int and s []int, where the value c[v] is an id
// of a component, to which the given vertex 'v' belongs, and s[i] is the size
// of component with id 'i'.
func InitConnectedComponents(g Graph) ConnectedComponents {
	cc := new(ccGraph)
	switch g.(type) {
	case *digraph:
		cc.graph = InitGraphFromGraph(g)
	default:
		cc.graph = g
	}
	cc.ccIndex = make([]int, g.VNum(), g.VNum())
	cc.ccSize = make([]int, g.VNum(), g.VNum())
	cc.comp = make([][]int, 1)
	cc.marked = make([]bool, g.VNum(), g.VNum())
	cc.count = 0
	for v := 0; v < g.VNum(); v++ {
		if !cc.marked[v] {
			runDFS(cc, v)
			cc.count++
			cc.comp = append(cc.comp, []int{})
		}
	}
	cc.comp = cc.comp[:cc.count]
	for i := range cc.comp {
		sort.Ints(cc.comp[i])
	}
	return cc
}

func runDFS(cc *ccGraph, v int) {
	cc.marked[v] = true
	cc.ccIndex[v] = cc.count
	cc.ccSize[cc.count]++
	cc.comp[cc.count] = append(cc.comp[cc.count], v)
	for _, w := range cc.graph.Edges(v) {
		if !cc.marked[w] {
			runDFS(cc, w)
		}
	}
}
