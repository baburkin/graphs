package graphs

import (
	"testing"
)

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
	testDataCCCount = []struct {
		name            string
		initGraphFunc   graphFactoryFunc
		expectedCCCount int
	}{
		{"testDAG1", initDAG1, 1},
		{"testDAG2", initDAG2, 2},
		{"testDAG3", initDAG3, 1},
	}
	testDataCCIndex = []struct {
		name            string
		initGraphFunc   graphFactoryFunc
		vertex          int
		expectedCCIndex int
	}{
		{"testDAG1", initDAG1, 3, 0},
		{"testDAG2", initDAG2, 0, 0},
		{"testDAG2", initDAG2, 2, 0},
		{"testDAG2", initDAG2, 3, 1},
		{"testDAG2", initDAG2, 5, 1},
		{"testDAG3", initDAG3, 5, 0},
	}
	testDataCCSize = []struct {
		name           string
		initGraphFunc  graphFactoryFunc
		component      int
		expectedCCSize int
	}{
		{"testDAG1", initDAG1, 0, 8},
		{"testDAG2", initDAG2, 0, 3},
		{"testDAG2", initDAG2, 1, 3},
		{"testDAG2", initDAG2, 2, 0},
		{"testDAG3", initDAG3, 0, 6},
		{"testDAG3", initDAG3, 1, 0},
	}
)

func TestConnCompCount(t *testing.T) {
	for _, test := range testDataCCCount {
		t.Run(test.name, func(t *testing.T) {
			g := test.initGraphFunc()
			cc := InitConnectedComponents(g)
			t.Logf("Checking number of connected components for graph: %v", g)
			count := cc.Count()
			t.Logf("Expected count: [%v], actual count: [%v]", test.expectedCCCount, count)
			if test.expectedCCCount != count {
				t.Fail()
			}
		})
	}
}

func TestConnCompIndex(t *testing.T) {
	for _, test := range testDataCCIndex {
		t.Run(test.name, func(t *testing.T) {
			g := test.initGraphFunc()
			cc := InitConnectedComponents(g)
			t.Logf("Checking connected components indices for vertices from graph: %v", g)
			index := cc.Index(test.vertex)
			t.Logf("Expected CC index: [%v], actual CC index: [%v]", test.expectedCCIndex, index)
			if test.expectedCCIndex != index {
				t.Fail()
			}
		})
	}
}

func TestConnCompSize(t *testing.T) {
	for _, test := range testDataCCSize {
		t.Run(test.name, func(t *testing.T) {
			g := test.initGraphFunc()
			cc := InitConnectedComponents(g)
			t.Logf("Checking connected components sizes for components from graph: %v", g)
			size := cc.CompSize(test.component)
			t.Logf("Expected CC size: [%v], actual size: [%v]", test.expectedCCSize, size)
			if test.expectedCCSize != size {
				t.Fail()
			}
		})
	}
}
