package graphs

import "fmt"

// VertexIterator type is a generic iterable collection
type VertexIterator interface {
	HasNext() bool
	Next() (int, error)
}

// Vertex type is a generic type for graph's vertices
type Vertex struct{}

// Edge type is a generic struct for (un)directed graph's edges
type Edge struct {
	v int
	w int
}

func (e Edge) String() string {
	return fmt.Sprintf("Edge(%v -> %v)", e.v, e.w)
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
