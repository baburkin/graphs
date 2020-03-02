package graphs

// NewDirectedGraph creates a new (empty) instance of DirectedGraph
func NewDirectedGraph() *DirectedGraph {
	var G *DirectedGraph
	G = new(DirectedGraph)
	return G
}
