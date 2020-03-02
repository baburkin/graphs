# Simple Golang library for graph processing

## Introduction

This library provides simple API (data structures and set of algorithms) to use in programs that run algorithms on graphs.

## Sample Usage

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a graph var
	graph := graphs.DirectedGraph()

	// Create a new reader from a file
	in := bufio.NewReader(os.File("sample_graph.txt"))

	// Initialize graph from a Reader
	graph.Init(in)

	if err, sortedGraph := graph.TopoSort(); err != nil {
		fmt.Println("Cannot do topological sort - the graph has cycles.")
	} else {
		fmt.Printf("Topologically sorted graph:\n%v\n", sortedGraph)
	}
}
```