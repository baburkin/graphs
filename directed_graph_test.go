package graphs

import (
	"fmt"
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
	initDAG4Cycled = func() DirectedGraph {
		g := InitDirectedGraph(5)
		// Make a cycle 4 -> 2 -> 3 -> 4
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)
		g.AddEdge(3, 4)
		g.AddEdge(4, 2)
		return g
	}
)

func TestInitDirectedGraphByVNum(t *testing.T) {
	graph := InitDirectedGraph(1000)
	testData := []struct {
		parameter string
		actual    int
		expected  int
	}{
		{"VNum()", graph.VNum(), 1000},
		{"ENum()", graph.ENum(), 0},
	}
	for _, test := range testData {
		t.Logf("Parameter: [%v], expected: [%v], actual: [%v]", test.parameter, test.expected, test.actual)
		if test.actual != test.expected {
			t.Fail()
		}
	}
}

func TestStringReprOfDirectedGraph(t *testing.T) {
	graph := InitDirectedGraph(3)
	graph.AddEdge(1, 2)
	actual := fmt.Sprintf("%v", graph)
	expected := "DirectedGraph(vertices: [3]; edges: [[] [2] []]; in degree [0 0 1])"
	if actual != expected {
		t.Errorf("Expected: [%v], actual: [%v]", expected, actual)
	}
}

func TestAddEdges(t *testing.T) {
	graph := InitDirectedGraph(5)
	graph.AddEdge(0, 4)
	graph.AddEdge(4, 1)
	graph.AddEdge(4, 2)
	graph.AddEdge(4, 3)
	actual := graph.ENum()
	expected := 4
	if actual != expected {
		t.Errorf("Expected: [%v], actual: [%v]", expected, actual)
	}
}

func TestInitDirectedGraphFromFile(t *testing.T) {
	filename := "examples/directed_7.txt"
	graph, err := LoadDirectedGraph(filename)
	testData := []struct {
		method_name string
		method      func() int
		expected    int
	}{
		{"graph.VNum()", graph.VNum, 7},
		{"graph.ENum()", graph.ENum, 7},
	}
	if err != nil {
		t.Errorf("Failed loading graph from file %v. Error: %v", filename, err)
	}
	for _, data := range testData {
		actual := data.method()
		if data.expected != actual {
			t.Errorf("%s=%v, but should be %v", data.method_name, actual, data.expected)
		}
	}
}

func TestInitDirectedGraphFromFileErr(t *testing.T) {
	_, err := LoadDirectedGraph("examples/directed_err.txt")
	expectedErr := fmt.Sprintf(errInvalidVertex+": strconv.Atoi: parsing \"5k\": invalid syntax", "5k")
	if err.Error() != expectedErr {
		t.Errorf("The error returned should be equal to: [%v]. Actual error is: [%v]", expectedErr, err.Error())
	}
}
