package graphs

import "fmt"

// DepthFirstOrder dfs = new DepthFirstOrder(G);
// order = dfs.reversePost();
// rank = new int[G.V()];
// int i = 0;
// for (int v : order)
// 	rank[v] = i++;

// TopoSort sorts the directed graph in topological order and returns the ordered slice of vertices
func TopoSort(g DirectedGraph) ([]int, error) {
	if IsDAG(g) {
		order := make([]int, g.VNum(), g.VNum())

		return order, nil
	} else {
		return nil, fmt.Errorf("The graph [%v] has a cycle", g)
	}
}
