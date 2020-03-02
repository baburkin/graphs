package graphs

type digraph struct {
	vertices  []int
	edges     map[int][]int
	vOrder    int
	eOrder    int
	outdegree []int
	indegree  []int
}

// DirectedGraph interface provides API to work with directed graphs
type DirectedGraph interface {
	// Add edge v -> w
	AddEdge(v int, w int)
	// All adjacent edges to the int v
	Edges(v int) []int
	// All edges in the digraph
	AllEdges() map[int][]int
	// All vertices in the digraph
	AllVertices() []int
	// How many edges are originating from v
	OutDegree(v int) int
	// How many edges are incident to v
	InDegree(v int) int
	// Reverse direction of all edges
	Reverse() *DirectedGraph
}

func (g *digraph) AddEdge(v int, w int) {

}

func (g *digraph) Edges(v int) []int {
	return g.edges[v]
}

func (g *digraph) AllEdges() map[int][]int {
	return g.edges
}

func (g *digraph) OutDegree(v int) {
	return g.edges[v]
}

func checkIOError(err error) {
	if err != nil {
		panic(err)
	}
}

// InitDirectedGraph initializes a new instance of DirectedGraph
// and populates from io.Reader instance.
func InitDirectedGraph(filename string) *DirectedGraph {
	g := new(digraph)
	g.vertices = make([]int, 0, 10)
	g.edges = make(map[int][]int)
	// bytesRead, err := ioutil.ReadFile(filename)
	// checkIOError(err)
	// scanner := bufio.NewScanner(nil)

	return nil
}
