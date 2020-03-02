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

	"github.com/baburkin/graphs"
)

func main() {
	// Create a new reader from a file
	in, err := os.Open("sample_graph.txt")
	if err != nil {
		fmt.Println("Error opening input file.")
		os.Exit(2)
	}

	// Initialize graph from a Reader
	graph := graphs.InitDirectedGraph(bufio.NewReader(in))

	if err, sortedGraph := graphs.TopoSort(graph); err != nil {
		fmt.Println("Cannot do topological sort - the graph has cycles.")
	} else {
		fmt.Printf("Topologically sorted graph:\n%v\n", sortedGraph)
	}
}

```