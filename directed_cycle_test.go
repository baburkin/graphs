package graphs

import (
	"reflect"
	"testing"
)

func TestIsDAG(t *testing.T) {
	g := initDAG4Cycled()
	t.Logf("Check if the directed graph is DAG: %v", g)
	if IsDAG(g) {
		t.Errorf("IsDAG() should have returned false")
	}
}

func TestGetCycleInGraph(t *testing.T) {
	g := initDAG4Cycled()
	t.Logf("Check the cycle in the directed graph: %v", g)
	expected := []int{4, 2, 3, 4}
	actual := GetCycleInGraph(g)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("GetCycleInGraph() should have returned [%v], but returned [%v]", expected, actual)
	}
}
