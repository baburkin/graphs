# Simple Golang library for graph processing

## Introduction

This library provides simple API (data structures and set of algorithms) to use in programs
that run algorithms on graphs.

Currently, it has algorithms for computing:

* the shortest path in the graph using BFS (breadth-first search)
* an order in a directed graph using DFS (depth-first search)
* a topological order (and its reverse) in a directed graph
* all connected components (self-connected islands) in a graph

## Sample Usage

```go
package main

// This is a sample usage of topological sort algorithm

import (
	"fmt"
	"os"

	"github.com/baburkin/graphs"
)

func main() {
	// Initialize a directed graph from a file
	graph, err := graphs.LoadDirectedGraph("example/directed_7.txt")
	if err != nil {
		fmt.Printf("Got an error: %v\n", err)
		os.Exit(2)
	}
	fmt.Printf("Initial graph: %v\n", graph)

	// Get a reversed graph
	reversedGraph := graph.Reverse()
	fmt.Printf("Reversed graph: %v\n", reversedGraph)

	// Try to topologically sort the graph
	if order, err := graphs.TopoSort(graph); err != nil {
		fmt.Printf("Cannot do topological sort: %v\n", err)
	} else {
		fmt.Printf("Topologically sorted graph (vertex order): %v\n", order)
	}

	// Try to topologically sort the reversed graph
	if order, err := graphs.TopoSort(reversedGraph); err != nil {
		fmt.Printf("Cannot do topological sort for reversed graph: %v\n", err)
	} else {
		fmt.Printf("Topologically sorted reversed graph (vertex order): %v\n", order)
	}
}
```
