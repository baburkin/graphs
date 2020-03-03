package main

// This is a sample usage

import (
	"fmt"
	"os"

	"github.com/baburkin/graphs"
)

func main() {
	// Initialize graph from a Reader
	graph, err := graphs.InitDirectedGraphFromFile("examples/directed_7.txt")
	if err != nil {
		fmt.Printf("Got an error: %v\n", err)
		os.Exit(2)
	}
	fmt.Printf("Graph: %v\n", graph)

	reversedGraph := graph.Reverse()
	fmt.Printf("Reversed graph: %v\n", reversedGraph)

	if err, sortedGraph := graphs.TopoSort(&graph); err != nil {
		fmt.Println("Cannot do topological sort - the graph has cycles.")
	} else {
		fmt.Printf("Topologically sorted graph:\n%v\n", sortedGraph)
	}
}
