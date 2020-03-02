package graphs

import "io"

// InitDirectedGraph initializes a new instance of DirectedGraph
// and populates from io.Reader instance.
func InitDirectedGraph(rd io.Reader) *DirectedGraph {
	var G *DirectedGraph
	G = new(DirectedGraph)
	return G
}
