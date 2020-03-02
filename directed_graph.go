package graphs

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type digraph struct {
	vertices []int
	edges    map[int][]int
	vOrder   int
	eOrder   int
	indegree []int
}

// DirectedGraph interface provides API to work with directed graphs
type DirectedGraph interface {
	// Add edge v -> w
	AddEdge(v int, w int)
	// All adjacent edges to the int v
	Edges(v int) []int
	// All edges in the digraph
	AllEdges() map[int][]int
	// All vertices in the digraph
	AllVertices() []int
	// How many edges are originating from v
	OutDegree(v int) int
	// How many edges are incident to v
	InDegree(v int) int
	// Reverse direction of all edges
	Reverse() DirectedGraph
}

func (g *digraph) AddEdge(v int, w int) {

}

func (g *digraph) Edges(v int) []int {
	return g.edges[v]
}

func (g *digraph) AllEdges() map[int][]int {
	return g.edges
}

func (g *digraph) AllVertices() []int {
	return g.vertices
}

func (g *digraph) OutDegree(v int) int {
	return len(g.edges[v])
}

func (g *digraph) InDegree(v int) int {
	return g.indegree[v]
}

func (g *digraph) Reverse() DirectedGraph {
	return g
}

func checkIOError(err error) {
	if err != nil {
		panic(err)
	}
}

// InitDirectedGraph initializes a new instance of DirectedGraph
// and populates from io.Reader instance.
func InitDirectedGraph(filename string) (DirectedGraph, error) {
	g := new(digraph)
	g.edges = make(map[int][]int)
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
		g.vOrder = size
		g.eOrder = 0
		g.vertices = make([]int, g.vOrder, g.vOrder)
		g.indegree = make([]int, g.vOrder, g.vOrder)
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
			g.edges[v] = append(g.edges[v], w)
			g.indegree[w]++
			g.eOrder++
		} else {
			return g, errors.New("The input data has wrong format")
		}
	}

	return g, scanner.Err()
}
