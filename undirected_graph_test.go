package graphs

import (
	"testing"
)

var (
	initGraph1 = func() Graph {
		g := InitGraph(7)
		g.AddEdge(1, 2)
		g.AddEdge(3, 6)
		return g
	}
	testDataHasEdge = []struct {
		name            string
		initFunc        graphFactoryFunc
		v, w            int
		expectedHasEdge bool
	}{
		{"testGraph1", initGraph1, 1, 2, true},
		{"testGraph1", initGraph1, 2, 1, true},
		{"testGraph1", initGraph1, 3, 6, true},
	}
)

func TestGraph_HasEdge(t *testing.T) {
	for _, test := range testDataHasEdge {
		t.Run(test.name, func(t *testing.T) {
			g := test.initFunc()
			t.Logf("Checking graph.HasEdge(%v,%v) method for: %v", g, test.v, test.w)
			result := g.HasEdge(test.v, test.w)
			t.Logf("Expected result: [%v], actual result: [%v]", result, test.expectedHasEdge)
			if result != test.expectedHasEdge {
				t.Fail()
			}
		})
	}
}
