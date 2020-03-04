package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitGraphByVNum(t *testing.T) {
	graph := InitDirectedGraph(1000)

	assert.Equal(t, graph.VNum(), 1000)
	assert.Equal(t, graph.ENum(), 0)
}

func TestAddEdges(t *testing.T) {

}
