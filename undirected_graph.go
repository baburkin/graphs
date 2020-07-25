package graphs

import "fmt"

// Graph interface provides API to work with undirected graphs
// The edges are unweighted
type Graph interface {
	VNum() int                 // Number of vertices
	ENum() int                 // Number of edges
	AddEdge(v int, w int) bool // Add edge v -- w
	HasEdge(v int, w int) bool // Does the graph contain edge v -- w?
	Edges(v int) []int         // Edges adjacent to v
	AllEdges() map[int][]int   // All edges in the graph
}

type graph struct {
	edges map[int][]int // I wonder if map[int]map[int]bool would be a better choice?
	V     int
	E     int
}

func (g *graph) String() string {
	return fmt.Sprintf("Graph("+
		"vertices: [%v]; "+
		"edges: %v)", g.V, g.edges)
}

func (g *graph) hasVertex(v int) bool {
	return v >= 0 && v < g.V
}

func (g *graph) VNum() int {
	return g.V
}

func (g *graph) ENum() int {
	return g.E
}

func (g *graph) Edges(v int) []int {
	return g.edges[v]
}

func (g *graph) AllEdges() map[int][]int {
	return g.edges
}

// TODO: need to rewrite HasEdge & AddEdge using sets, not slices
func (g *graph) HasEdge(v int, w int) bool {
	for _, e := range g.edges[v] {
		if e == w {
			return true
		}
	}
	return false
}

func (g *graph) AddEdge(v int, w int) bool {
	if !g.hasVertex(v) || !g.hasVertex(w) {
		return false
	}
	if !g.HasEdge(v, w) {
		g.edges[v] = append(g.edges[v], w)
		g.edges[w] = append(g.edges[w], v)
		g.E++
	}
	return true
}

// InitGraph initializes a new instance of Graph
// with a given number of vertices
func InitGraph(verticesNum int) Graph {
	g := new(graph)
	g.V = verticesNum
	g.E = 0
	g.edges = make(map[int][]int, verticesNum)
	return g
}

// InitGraphFromGraph initializes a new instance of unidirected graph
// from a given instance of graph by copying all edges
func InitGraphFromGraph(g Graph) Graph {
	newG := new(graph)
	newG.V = g.VNum()
	newG.edges = make(map[int][]int, newG.V)
	for v := 0; v < newG.V; v++ {
		for _, w := range g.Edges(v) {
			newG.AddEdge(v, w)
		}
	}
	return newG
}
