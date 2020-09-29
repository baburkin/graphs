package graphs

import (
	"fmt"
	"os"
	"testing"
)

func ExampleTopoSort() {
	// Initialize a directed graph from a file
	graph, err := LoadDirectedGraph("examples/directed_7.txt")
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

var (
	testDataTopoSort = []struct {
		name     string
		initFunc digraphFactoryFunc
	}{
		{"testDAG1", initDAG1},
	}
)

func TestTopoSort(t *testing.T) {
	for _, test := range testDataTopoSort {
		t.Run(test.name, func(t *testing.T) {
			g := test.initFunc()
			gSorted, err := TopoSort(g)

			if err != nil {
				t.Fatalf("Could not topologically sort the graph %v, error: %v", g, err)
			}

			// Give the output of the initial and sorted graph
			t.Logf("Initial graph: %v\n", g)
			t.Logf("Topologically sorted graph: %v\n", gSorted)

			// Validate the topological sort
			sortedIndex := make([]int, len(gSorted), len(gSorted))
			for i, v := range gSorted {
				sortedIndex[v] = i
			}
			for v, edgesFromV := range g.AllEdges() {
				for _, w := range edgesFromV {
					t.Logf("Vertex [%v] should come before [%v]... ", v, w)
					if sortedIndex[v] >= sortedIndex[w] {
						t.Logf("Error: topological sort is incorrect")
						t.Fail()
					}
				}
			}
		})
	}
}

func TestTopoSortReverse(t *testing.T) {
	g := initDAG1()
	gSorted, err := TopoSortReverse(g)

	// gSorted[2], gSorted[7] = gSorted[7], gSorted[2] <-- used to test path for incorrect sort

	if err != nil {
		t.Fatalf("Could not topologically sort the graph %v, error: %v", g, err)
	}

	// Give the output of the initial and sorted graph
	t.Logf("Initial graph: %v\n", g)
	t.Logf("Topologically sorted graph: %v\n", gSorted)

	// Validate the topological sort
	sortedIndex := make([]int, len(gSorted), len(gSorted))
	for i, v := range gSorted {
		sortedIndex[v] = i
	}
	for v, edgesFromV := range g.AllEdges() {
		for _, w := range edgesFromV {
			t.Logf("Vertex [%v] should come before [%v]... ", w, v)
			if sortedIndex[w] < sortedIndex[v] {
				t.Logf("OK")
			} else {
				t.Logf("Eror")
				t.Fatal("Topological sort is incorrect")
			}
		}
	}
}
