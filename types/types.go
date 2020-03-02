package types

// Bag type is a generic collection
type Bag interface {
	Contains() bool
}

// Vertex type is a generic type for graph's vertices
type Vertex struct {
}

// Edge type is a generic type for graph's edges
type Edge struct {
	source *Vertex
	dest   *Vertex
}

// DirectedEdge provides a special kind of edge which has direction.
// Used in DirectedGraph's
type DirectedEdge interface {
}

// UndirectedGraph interface provides API to work with undirected graphs
type UndirectedGraph interface {
}

// DirectedGraph interface provides API to work with directed graphs
type DirectedGraph interface {
	Edges() []DirectedEdge
	Vertices() []Vertex
}
