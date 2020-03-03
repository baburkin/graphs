package graphs

import "fmt"

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
		return ReversePostOrder(g), nil
	}
	return nil, fmt.Errorf("The graph [%v] has a cycle", g)
}
