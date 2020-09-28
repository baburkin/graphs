package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type initGraphFunc func() DirectedGraph

var (
	initDAG1 = func() DirectedGraph {
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
	initDAG2 = func() DirectedGraph {
		g := InitDirectedGraph(6)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 0) // introducing a component with a cycle
		g.AddEdge(3, 4)
		g.AddEdge(5, 4)
		return g
	}
	initDAG3 = func() DirectedGraph {
		g := initDAG2()
		g.AddEdge(2, 4) // now the graph has only one component
		return g
	}
	testDataCCSize = []struct {
		name          string
		initGraphFunc initGraphFunc
		expectedCount int
	}{
		{"testDAG1", initDAG1, 1},
		{"testDAG2", initDAG2, 2},
		{"testDAG3", initDAG3, 1},
	}
)

func TestConnCompCount(t *testing.T) {
	for _, test := range testDataCCSize {
		t.Run(test.name, func(t *testing.T) {
			g := test.initGraphFunc()
			cc := InitConnectedComponents(g)
			t.Logf("Checking connected components for graph: %v", g)
			count := cc.Count()
			t.Logf("Expected count: [%v], actual count: [%v]", test.expectedCount, count)
			if test.expectedCount != count {
				t.Fail()
			}
		})
	}
}

func TestConnComp1(t *testing.T) {
	g := initDAG1()

	cc := InitConnectedComponents(g)
	assert.Equal(t, 8, cc.CompSize(0))
	assert.Equal(t, 0, cc.Index(3))
}

func TestConnectedComponentsInDAG2(t *testing.T) {
	g := initDAG2()

	cc := InitConnectedComponents(g)
	assert.Equal(t, 3, cc.CompSize(0))
	assert.Equal(t, 3, cc.CompSize(1))
	// only two components should be found
	assert.Equal(t, 0, cc.CompSize(2))

	assert.Equal(t, 0, cc.Index(0))
	assert.Equal(t, 0, cc.Index(1))
	assert.Equal(t, 0, cc.Index(2))
	assert.Equal(t, 1, cc.Index(3))
	assert.Equal(t, 1, cc.Index(4))
	assert.Equal(t, 1, cc.Index(5))
}

// TestConnectedComponentsInDAG3 - tests that after adding an edge
// connecting two vertices in two separate components make the graph
// a single-component one
func TestConnectedComponentsInDAG3(t *testing.T) {
	g := initDAG2()

	cc := InitConnectedComponents(g)
	assert.Equal(t, 3, cc.CompSize(0))

	g.AddEdge(2, 4)

	cc = InitConnectedComponents(g)
	assert.Equal(t, 6, cc.CompSize(0))
	assert.Equal(t, 0, cc.CompSize(1))
}
