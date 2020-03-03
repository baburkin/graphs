package graphs

import "fmt"

type topoOrder struct {
	DirectedGraph
}

// DepthFirstOrder dfs = new DepthFirstOrder(G);
// order = dfs.reversePost();
// rank = new int[G.V()];
// int i = 0;
// for (int v : order)
// 	rank[v] = i++;

// Rank shows the rank of the vertex if topological sort exists, -1 otherwise
func Rank(g DirectedGraph, v Vertex) int {
	return -1
}

// TopoSort sorts the directed graph in topological order and returns the ordered slice of vertices
func TopoSort(g DirectedGraph) ([]int, error) {
	if IsDAG(g) {
		order := make([]int, g.VNum(), g.VNum())

		return order, nil
	}
	return nil, fmt.Errorf("The graph [%v] has a cycle", g)
}
