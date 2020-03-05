package graphs

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleTopoSort() {
	// Initialize a directed graph from a file
	graph, err := InitDirectedGraphFromFile("examples/directed_7.txt")
	if err != nil {
		fmt.Printf("Got an error: %v\n", err)
		os.Exit(2)
	}
	fmt.Printf("Initial graph: %v\n", graph)

	// Attempt to topologically sort the graph
	if order, err := TopoSort(graph); err != nil {
		fmt.Printf("Cannot do topological sort: %v\n", err)
	} else {
		fmt.Printf("Topologically sorted graph (vertex order): %v\n", order)
	}
}

func initTestGraph() DirectedGraph {
	g := InitDirectedGraph(8)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(1, 5)
	g.AddEdge(1, 7)
	g.AddEdge(3, 1)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(6, 4)
	g.AddEdge(6, 7)
	return g
}

func TestTopoSort(t *testing.T) {
	g := initTestGraph()
	gSorted, err := TopoSort(g)

	// gSorted[2], gSorted[7] = gSorted[7], gSorted[2] <-- used to test path for incorrect sort

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
			fmt.Printf("Vertex [%v] should come before [%v]... ", v, w)
			if sortedIndex[v] < sortedIndex[w] {
				fmt.Println("OK")
			} else {
				fmt.Println("Eror")
				t.Fatal("Topological sort is incorrect")
			}
		}
	}

}

func TestTopoSortReverse(t *testing.T) {
	g := initTestGraph()
	gSorted, err := TopoSortRevers(g)

	// gSorted[2], gSorted[7] = gSorted[7], gSorted[2] <-- used to test path for incorrect sort

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
			fmt.Printf("Vertex [%v] should come before [%v]... ", w, v)
			if sortedIndex[w] < sortedIndex[v] {
				fmt.Println("OK")
			} else {
				fmt.Println("Eror")
				t.Fatal("Topological sort is incorrect")
			}
		}
	}
}
