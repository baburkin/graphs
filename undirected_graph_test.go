package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func initTestGraph1() Graph {
	g := InitGraph(7)
	g.AddEdge(1, 2)
	g.AddEdge(3, 6)
	return g
}

func TestHasEdge(t *testing.T) {
	g := initTestGraph1()
	assert.Equal(t, g.HasEdge(1, 2), true)
	assert.Equal(t, g.HasEdge(2, 1), true)
	assert.Equal(t, g.HasEdge(3, 6), true)
	assert.Equal(t, g.HasEdge(6, 3), true)
}
