// Copyright 2020 Dmitry Baburkin

// Package graphs implements some graphs data structures
// and algorithms
package graphs

import "fmt"

type digraphFactoryFunc func() DirectedGraph

type graphFactoryFunc func() Graph

// VertexIterator type is a generic iterable collection
type VertexIterator interface {
	HasNext() bool
	Next() (int, error)
}

// EdgeIterator interface provides a way to
type EdgeIterator interface {
	HasNext() bool
	Next() Edge
}

// Vertex type is a generic type for graph's vertices
type Vertex interface{}

// Edge type is a generic struct for (un)directed graph's edges
type Edge struct {
	v Vertex
	w Vertex
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
