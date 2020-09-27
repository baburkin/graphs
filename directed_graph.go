package graphs

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"strconv"
)

// DirectedGraph interface provides API to work with directed graphs
type DirectedGraph interface {
	Graph                   // DirectedGraph is an extension of Graph API
	OutDegree(v int) int    // How many edges are originating from the v
	InDegree(v int) int     // How many edges are incident to v
	Reverse() DirectedGraph // Reverse direction of all edges
}

type digraph struct {
	graph
	indegree []int // number of incident vertices for each vertex
}

var (
	errInvalidVertex = "graphs.LoadDirectedGraph: invalid vertex [%v] (should be int)"
)

func (g *digraph) String() string {
	return fmt.Sprintf("DirectedGraph("+
		"vertices: [%v]; "+
		"edges: %v; "+
		"in degree %v)", g.V, g.edges, g.indegree)
}

// AddEdge for digraph is not symmetrical as for an undirected graph (graph)
func (g *digraph) AddEdge(v int, w int) bool {
	if g.hasVertex(v) && g.hasVertex(w) {
		g.edges[v] = append(g.edges[v], w)
		g.indegree[w]++
		g.E++
		return true
	}
	return false
}

// OutDegree
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

// InitDirectedGraph initializes a new instance of DirectedGraph
// with a given number of vertices
func InitDirectedGraph(verticesNum int) DirectedGraph {
	g := new(digraph)
	g.V = verticesNum
	g.E = 0
	g.edges = make([][]int, verticesNum, verticesNum)
	g.indegree = make([]int, verticesNum, verticesNum)
	return g
}

// LoadDirectedGraph initializes a new instance of DirectedGraph
// from a file.
func LoadDirectedGraph(filename string) (DirectedGraph, error) {
	r, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return LoadDirectedGraphFromReader(r)
}

// LoadDirectedGraphFromReader initializes a new instance of DirectedGraph
// from any implementor of io.Reader interface.
func LoadDirectedGraphFromReader(reader io.Reader) (DirectedGraph, error) {
	scanner := bufio.NewScanner(reader)
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
				return g, errors.Wrapf(err, errInvalidVertex, scanner.Text())
			}
			if !scanner.Scan() {
				return g, fmt.Errorf("graphs.LoadDirectedGraph: " +
					"the input data from has wrong format")
			}
			w, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return g, errors.Wrapf(err, errInvalidVertex, scanner.Text())
			}

			if !g.AddEdge(v, w) {
				return g, fmt.Errorf("Could not add edge [%d -> %d]. "+
					"Most likely, indices are out of range\n", v, w)
			}
		}
		return g, scanner.Err()
	}

	return nil, fmt.Errorf("could not initialize directed graph from object [%v]; "+
		"most likely, the input format is incorrect", reader)
}
