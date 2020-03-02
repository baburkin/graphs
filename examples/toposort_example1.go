package main

// This is a sample usage

import (
	"fmt"

	"github.com/baburkin/graphs"
)

func main() {
	// Initialize graph from a Reader
	graph, err := graphs.InitDirectedGraph("directed_7.txt")
	if err != nil {
		fmt.Printf("Got an error: %v\n", err)
	}

	if err, sortedGraph := graphs.TopoSort(graph); err != nil {
		fmt.Println("Cannot do topological sort - the graph has cycles.")
	} else {
		fmt.Printf("Topologically sorted graph:\n%v\n", sortedGraph)
	}
}
