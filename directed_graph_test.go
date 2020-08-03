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
	assert.Equal(t,
		"DirectedGraph(vertices: [3]; edges: [[] [2] []]; in degree [0 0 1])",
		fmt.Sprintf("%v", graph))
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
	graph, err := LoadDirectedGraph("examples/directed_7.txt")

	assert.Nil(t, err)
	assert.Equal(t, graph.VNum(), 7)
	assert.Equal(t, graph.ENum(), 7)
}

func TestInitDirectedGraphFromFile2(t *testing.T) {
	graph, err := LoadDirectedGraph("examples/directed_err.txt")

	if assert.Error(t, err) {
		assert.EqualError(t, err, fmt.Sprintf(errInvalidVertex+
			": strconv.Atoi: parsing \"5k\": invalid syntax", "5k"))
	}

	assert.Equal(t, 2, graph.ENum())
}
