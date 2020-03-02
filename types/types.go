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
	source *Vertex,
	dest *Vertex,
}

type UndirectedGraph struct {

}

type DirectedGraph struct {
	
}
