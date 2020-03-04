package graphs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDirectedGraphByVNum(t *testing.T) {
	graph := InitDirectedGraph(1000)

	assert.Equal(t, graph.VNum(), 1000)
	assert.Equal(t, graph.ENum(), 0)
}

func TestStringReprOfDirectedGraph(t *testing.T) {
	graph := InitDirectedGraph(3)
	graph.AddEdge(1, 2)
	assert.Equal(t, fmt.Sprintf("%v", graph),
		"DirectedGraph(vertices: [3]; edges: map[1:[2]]; in degree [0 0 1])")
}

func TestAddEdges(t *testing.T) {
	graph := InitDirectedGraph(5)
	graph.AddEdge(0, 4)
	graph.AddEdge(4, 1)
	graph.AddEdge(4, 2)
	graph.AddEdge(4, 3)

	assert.Equal(t, graph.ENum(), 4)
}

func TestInitDirectedGraphFromFile(t *testing.T) {
	graph, err := InitDirectedGraphFromFile("examples/directed_7.txt")

	assert.Nil(t, err)
	assert.Equal(t, graph.VNum(), 7)
	assert.Equal(t, graph.ENum(), 7)
}
