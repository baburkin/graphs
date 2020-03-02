package types

// Bag type is a generic collection
type Bag interface {
	Contains() bool
}

// Vertex type is a generic type for graph's vertices
type Vertex interface{}

// Edge type is a generic type for graph's edges
type Edge interface {
	V() Vertex
	W() Vertex
}

// DirectedEdge provides a special kind of edge which has direction.
// Used in DirectedGraph's
type DirectedEdge interface {
	V() Vertex
	W() Vertex
}

// UndirectedGraph interface provides API to work with undirected graphs
type UndirectedGraph interface {
	AddEdge(v int, w int)
	Edges() []Edge
	Vertices() []Vertex
}

// DirectedGraph interface provides API to work with directed graphs
type DirectedGraph interface {
	AddEdge(v int, w int)
	Edges() []DirectedEdge
	Vertices() []Vertex
	OutDegree(vertex int) int
	InDegree(vertex int) int
	Reverse() *DirectedGraph
}
