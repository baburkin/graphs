package graphs

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type digraph struct {
	edges    map[int][]int
	V        int
	E        int
	indegree []int
}

// DirectedGraph interface provides API to work with directed graphs
type DirectedGraph interface {
	// Number of vertices
	VNum() int
	// Number of edges
	ENum() int
	// Add edge v -> w
	AddEdge(v int, w int) bool
	// All adjacent edges to the int v
	Edges(v int) []int
	// All edges in the digraph
	AllEdges() map[int][]int
	// All vertices in the digraph
	OutDegree(v int) int
	// How many edges are incident to v
	InDegree(v int) int
	// Reverse direction of all edges
	Reverse() DirectedGraph
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
	gNew := new(digraph)
	gNew.E = g.E
	gNew.V = g.V

	return gNew
}

func checkIOError(err error) {
	if err != nil {
		panic(err)
	}
}

// InitDirectedGraph initializes a new instance of DirectedGraph
// with a given number of vertices
func InitDirectedGraph(verticesNum int) (DirectedGraph, error) {
	g := new(digraph)
	g.V = verticesNum
	g.E = 0
	g.edges = make(map[int][]int)
	return g, nil
}

// InitDirectedGraph initializes a new instance of DirectedGraph
// from a file.
func InitDirectedGraphFromFile(filename string) (DirectedGraph, error) {
	g := new(digraph)
	r, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	if scanner.Scan() {
		gType := scanner.Text() // 'u', 'd', 'uw', 'dw' ...
		if gType != "d" {
			return nil, errors.New("the graph is not directed")
		}
		fmt.Printf("Directed graph it is!\n")
	}

	if scanner.Scan() {
		size, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		fmt.Printf("Our directed graph has [%d] elements.\n", size)
		g.V = size
		g.E = 0
		g.indegree = make([]int, g.V, g.V)
		g.edges = make(map[int][]int)
	}

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
			fmt.Printf("Adding edge: [%d -> %d].\n", v, w)
			if g.edges[v] == nil {
				g.edges[v] = make([]int, 1)
			}
			if !g.AddEdge(v, w) {
				fmt.Printf("Could not add edge [%d -> %d].\n", v, w)
			}
			g.indegree[w]++
			g.E++
		} else {
			return g, errors.New("The input data has wrong format")
		}
	}

	return g, scanner.Err()
}
