package graphs

type digraph struct {
	vertices []int
	edges    map[int][]int
	vOrder   int
	eOrder   int
}

// DirectedGraph interface provides API to work with directed graphs
type DirectedGraph interface {
	// Add edge v -> w
	AddEdge(v int, w int)
	// All adjacent edges to the int v
	Edges(v int) []int
	AllEdges() map[int][]int
	AllVertices() []int
	OutDegree(v int) int
	InDegree(v int) int
	Reverse() *DirectedGraph
}

func (*digraph) AddEdge(v int, w int) {

}

func (g *digraph) Edges(v int) []int {
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
