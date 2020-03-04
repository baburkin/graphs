package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDAG(t *testing.T) {
	g := InitDirectedGraph(5)

	// Make a cycle 2 -> 3 -> 4 -> 5
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 2)

	assert.Equal(t, false, IsDAG(g))
}

func TestGetCycleInGraph(t *testing.T) {
	g := InitDirectedGraph(5)

	// Make a cycle 4 -> 2 -> 3 -> 4
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 2)

	assert.Equal(t, []int{4, 2, 3, 4}, GetCycleInGraph(g))
}
