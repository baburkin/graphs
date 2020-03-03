package main

// This is a sample usage

import (
	"fmt"
	"os"

	"github.com/baburkin/graphs"
)

func main() {
	// Initialize graph from a Reader
	graph, err := graphs.InitDirectedGraphFromFile("examples/cycled_7.txt")
	if err != nil {
		fmt.Printf("Got an error: %v\n", err)
		os.Exit(2)
	}
	fmt.Printf("Graph: %v\n", graph)

	reversedGraph := graph.Reverse()
	fmt.Printf("Reversed graph: %v\n", reversedGraph)

	if order, err := graphs.TopoSort(graph); err != nil {
		fmt.Printf("Cannot do topological sort: %v\n", err)
	} else {
		fmt.Printf("Topologically sorted graph (vertex order):\n%v\n", order)
	}
}
