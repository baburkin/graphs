package graphs

import (
	"fmt"
	"os"

	"github.com/baburkin/graphs"
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
