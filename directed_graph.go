package graphs

import (
	"io"
	"io/ioutil"
	"os"
)

type digraph struct {
	vertices []int
	edges    map[int][]int
}

// DirectedGraph interface provides API to work with directed graphs
type DirectedGraph interface {
	// Add edge v -> w
	AddEdge(v int, w int)
	// All adjacent edges to the int v
	Edges(v int) []int
	AllEdges() map[int][]int
	AllVertices() []int
	OutDegree(v int) int
	InDegree(v int) int
	Reverse() *DirectedGraph
}

func (*digraph) AddEdge(v int, w int) {

}

func checkIOError(err error) {
	if err != nil {
		panic("Error reading from file: " + err);
	}
}

// InitDirectedGraph initializes a new instance of DirectedGraph
// and populates from io.Reader instance.
func InitDirectedGraph(filename string) *DirectedGraph {
	var g *DirectedGraph
	g = make(digraph)
	g.vertices = make([]int, 10)
	g.edges = make(map[int][]int)
	bytesRead, err := ioutil.ReadFile(filename)
	checkIOError(err);
	return g;
}

func (g *digraph) 