package graphs

// IsDAG determines if the directed graph is a DAG (directed acyclic graph)
func (*Topological) IsDAG() bool {
	return false
}

// Rank shows the rank of the vertex if topological sort exists, -1 otherwise
func (*Topological) Rank(v Vertex) int {
	return 0
}
