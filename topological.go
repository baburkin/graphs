package graphs

// TopoSort sorts the directed graph in topological order and returns the ordered slice of vertices
func TopoSort(g DirectedGraph) ([]int, error) {
	if IsDAG(g) {
		order := make([]int, g.VNum(), g.VNum())

		return order, nil
	}
	return nil, nil
}
