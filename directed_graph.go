package graphs

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// DirectedGraph interface provides API to work with directed graphs
type DirectedGraph interface {
	Graph                   // DirectedGraph is an extension of Graph API
	OutDegree(v int) int    // All vertices in the digraph
	InDegree(v int) int     // How many edges are incident to v
	Reverse() DirectedGraph // Reverse direction of all edges
}

type digraph struct {
	edges    map[int][]int
	V        int
	E        int
	indegree []int
}

func (g *digraph) String() string {
	return fmt.Sprintf("DirectedGraph("+
		"vertices: [%v]; "+
		"edges: %v; "+
		"in degree %v)", g.V, g.edges, g.indegree)
}

func (g *digraph) isVertexValid(v int) bool {
	return v >= 0 && v < g.V
}

func (g *digraph) VNum() int {
	return g.V
}

func (g *digraph) ENum() int {
	return g.E
}

func (g *digraph) AddEdge(v int, w int) bool {
	if g.isVertexValid(v) && g.isVertexValid(w) {
		g.edges[v] = append(g.edges[v], w)
		g.indegree[w]++
		g.E++
		return true
	}
	return false
}

func (g *digraph) Edges(v int) []int {
	return g.edges[v]
}

func (g *digraph) AllEdges() map[int][]int {
	return g.edges
}

func (g *digraph) OutDegree(v int) int {
	return len(g.edges[v])
}

func (g *digraph) InDegree(v int) int {
	return g.indegree[v]
}

func (g *digraph) Reverse() DirectedGraph {
	gNew := InitDirectedGraph(g.V)
	for v, e := range g.edges {
		if len(e) > 0 {
			for _, w := range e {
				gNew.AddEdge(w, v)
			}
		}
	}
	return gNew
}

func checkIOError(err error) {
	if err != nil {
		panic(err)
	}
}

// InitDirectedGraph initializes a new instance of DirectedGraph
// with a given number of vertices
func InitDirectedGraph(verticesNum int) DirectedGraph {
	g := new(digraph)
	g.V = verticesNum
	g.E = 0
	g.edges = make(map[int][]int, verticesNum)
	g.indegree = make([]int, verticesNum, verticesNum)
	return g
}

// InitDirectedGraphFromFile initializes a new instance of DirectedGraph
// from a file.
func InitDirectedGraphFromFile(filename string) (DirectedGraph, error) {
	r, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	// First number should indicate the order of the graph (number of vertices)
	if scanner.Scan() {
		size, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		// fmt.Printf("Our directed graph has [%d] elements.\n", size)
		g := InitDirectedGraph(size)

		for scanner.Scan() {
			v, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return g, err
			}
			if scanner.Scan() {
				w, err := strconv.Atoi(scanner.Text())
				if err != nil {
					return g, err
				}

				if !g.AddEdge(v, w) {
					return g, fmt.Errorf("Could not add edge [%d -> %d]. "+
						"Most likely, indices are out of range\n", v, w)
				}
			} else {
				return g, errors.New("The input data has wrong format")
			}
		}
		return g, scanner.Err()
	}

	return nil, fmt.Errorf("Could not initialize directed graph from file [%v]."+
		"Most likely, the format in the file is incorrect.", filename)
}
