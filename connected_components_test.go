package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func initTestOneComponentDAG() DirectedGraph {
	g := InitDirectedGraph(8)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(1, 5)
	g.AddEdge(1, 7)
	g.AddEdge(3, 1)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(6, 4)
	g.AddEdge(6, 7)
	return g
}

func TestConnectedComponentsInDAG(t *testing.T) {
	g := initTestOneComponentDAG()

	cc, size := ConnectedComponents(g)
	assert.Equal(t, 8, size[0])
	assert.Equal(t, 0, cc[3])
}
