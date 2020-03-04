package graphs

// Graph interface provides API to work with undirected graphs
// The edges are unweighted
type Graph interface {
	VNum() int                 // Number of vertices
	ENum() int                 // Number of edges
	AddEdge(v int, w int) bool // Add edge v -- w
	Edges(v int) []int         // All adjacent edges to v
	AllEdges() map[int][]int   // All edges in the graph
}
