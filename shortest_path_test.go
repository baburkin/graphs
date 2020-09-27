package graphs

import (
	"strings"
	"testing"
)

var (
	testDataShortestPath = []struct {
		name     string
		graph    string
		v        int
		w        int
		expected int
	}{
		{"digraph-7", "7\n5 0\n4 5\n6 5\n1 3\n5 2\n3 4\n3 6", 1, 0, 4},
		{"digraph-6", "6\n0 1\n1 2\n2 3\n3 4\n4 5\n1 5", 0, 5, 2},
		{"digraph-2a", "2\n0 1", 1, 1, 0},
		{"digraph-2b", "2\n0 1", 0, 1, 1},
	}
)

func TestShortestPathBFS(t *testing.T) {
	for _, test := range testDataShortestPath {
		t.Run(test.name, func(t *testing.T) {
			g, err := LoadDirectedGraphFromReader(strings.NewReader(test.graph))
			if err != nil {
				t.Logf("Could not init graph [%v], got an error: [%v]", test.graph, err)
			}
			path := ShortestPathBFS(g, test.v, test.w)
			t.Logf("Shortest path [%v -> %v] is %v; expected: %v", test.v, test.w, path, test.expected)
			if path != test.expected {
				t.Fail()
			}
		})
	}
}
