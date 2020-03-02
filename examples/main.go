package examples

// This is a sample usage

import (
	"fmt"
	"os"

	"github.com/baburkin/graphs"
)

func main() {
	// Create a graph var
	// graph := graphs.NewDirectedGraph()
	graph := graphs.NewDirectedGraph()

	// Create a new reader from a file
	in, err := os.Open("sample_graph.txt")
	if err != nil {
		fmt.Println("Error opening input file.")
		os.Exit(2)
	}

	// Initialize graph from a Reader
	graph.Init(in)

	if err, sortedGraph := graph.TopoSort(); err != nil {
		fmt.Println("Cannot do topological sort - the graph has cycles.")
	} else {
		fmt.Printf("Topologically sorted graph:\n%v\n", sortedGraph)
	}
}
