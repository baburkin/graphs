package graphs

import (
	"fmt"
	"os"
	"testing"

	"github.com/baburkin/graphs"
	"github.com/stretchr/testify/assert"
)

func ExampleTopoSort() {
	// Initialize a directed graph from a file
	graph, err := graphs.InitDirectedGraphFromFile("examples/directed_7.txt")
	if err != nil {
		fmt.Printf("Got an error: %v\n", err)
		os.Exit(2)
	}
	fmt.Printf("Initial graph: %v\n", graph)

	// Attempt to topologically sort the graph
	if order, err := graphs.TopoSort(graph); err != nil {
		fmt.Printf("Cannot do topological sort: %v\n", err)
	} else {
		fmt.Printf("Topologically sorted graph (vertex order): %v\n", order)
	}
}

func TestTopoSort(t *testing.T) {
	// Initialize the graph and sort it
	g := graphs.InitDirectedGraph(8)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(1, 5)
	g.AddEdge(1, 7)
	g.AddEdge(3, 1)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(6, 4)
	g.AddEdge(6, 7)
	gSorted, err := graphs.TopoSort(g)

	// Assert that this graph can indeed be sorted topologically
	assert.Equal(t, nil, err)

	// Give the output of the initial and sorted graph
	fmt.Printf("Initial graph: %v\n", g)
	fmt.Printf("Topologically sorted graph: %v\n", gSorted)

	// Validate the topological sort
	sortedIndex := make([]int, len(gSorted), len(gSorted))
	for i, v := range gSorted {
		sortedIndex[v] = i
	}
	for v, edgesFromV := range g.AllEdges() {
		for _, w := range edgesFromV {
			fmt.Printf("Vertex [%v] should come after [%v]... ", v, w)
			assert.Less(t, sortedIndex[w], sortedIndex[v])
			fmt.Println("OK")
		}
	}

}
