package graphs

// Bag type is a generic collection
type Bag interface {
	Contains() bool
}

// Vertex type is a generic type for graph's vertices
type Vertex struct{}

// Edge type is a generic type for graph's edges
type Edge struct {
	v Vertex
	w Vertex
}

// DirectedEdge provides a special kind of edge which has direction.
// Used in DirectedGraph's
type DirectedEdge interface {
	V() Vertex
	W() Vertex
}

// UndirectedGraph interface provides API to work with undirected graphs
type UndirectedGraph interface {
	AddEdge(v Vertex, w Vertex)
	Edges() []Edge
	Vertices() []Vertex
}

type digraph struct {
	vertices []Vertex
	edges    []Edge
}

// DirectedGraph interface provides API to work with directed graphs
type DirectedGraph interface {
	// Add edge v -> w
	AddEdge(v Vertex, w Vertex)
	// All adjacent edges to the vertex v
	Edges(v Vertex) []DirectedEdge
	AllEdges() []DirectedEdge
	AllVertices() []Vertex
	OutDegree(v Vertex) int
	InDegree(v Vertex) int
	Reverse() *DirectedGraph
}
