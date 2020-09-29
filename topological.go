package graphs

import "fmt"

//goland:noinspection SpellCheckingInspection
type topoOrder struct {
	DirectedGraph
}

// Rank shows the rank of the vertex if topological sort exists, -1 otherwise
func Rank(g DirectedGraph, v Vertex) int {
	return -1
}

// TopoSort sorts the directed graph in topological order and returns the ordered slice of vertices
func TopoSort(g DirectedGraph) ([]int, error) {
	if IsDAG(g) {
		return DFSReversePostOrder(g), nil
	}
	return nil, fmt.Errorf("the graph [%v] has a cycle hence topological sort is not possible", g)
}

// TopoSortReverse returns the slice of vertices ordered reversely to topological order.
// It is useful in case where e.g. vertices (v, w) are tasks and an edge (v -> w) denotes
// that v depends on w.
//
// In the resulting order w should come before v, whereas in topological order it's vice versa.
func TopoSortReverse(g DirectedGraph) ([]int, error) {
	if IsDAG(g) {
		return DFSPostOrder(g), nil
	}
	return nil, fmt.Errorf("the graph [%v] has a cycle hence task dependency ordering is not possible", g)
}
