package main

// This is a sample usage of graphs package

import (
	"fmt"
	"os"

	"github.com/baburkin/graphs"
)

func main() {
	// Initialize a directed graph from a file
	graph, err := graphs.InitDirectedGraphFromFile("examples/cycled_7.txt")
	if err != nil {
		fmt.Printf("Got an error: %v\n", err)
		os.Exit(2)
	}
	fmt.Printf("Initial graph: %v\n", graph)

	// Get a reversed graph
	reversedGraph := graph.Reverse()
	fmt.Printf("Reversed graph: %v\n", reversedGraph)

	if order, err := graphs.TopoSort(graph); err != nil {
		fmt.Printf("Cannot do topological sort: %v\n", err)
	} else {
		fmt.Printf("Topologically sorted graph (vertex order):\n%v\n", order)
	}
}
