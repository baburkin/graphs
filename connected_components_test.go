package graphs

import (
	"reflect"
	"testing"
)

type digraphFactoryFunc func() DirectedGraph

var (
	testDataCCCount = []struct {
		name            string
		initGraphFunc   digraphFactoryFunc
		expectedCCCount int
	}{
		{"testDAG1", initDAG1, 1},
		{"testDAG2", initDAG2, 2},
		{"testDAG3", initDAG3, 1},
	}
	testDataCCIndex = []struct {
		name            string
		initGraphFunc   digraphFactoryFunc
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
		initGraphFunc  digraphFactoryFunc
		component      int
		expectedCCSize int
		expectedComps  []int
	}{
		{"testDAG1", initDAG1, 0, 8,
			[]int{0, 1, 2, 3, 4, 5, 6, 7}},
		{"testDAG2", initDAG2, 0, 3,
			[]int{0, 1, 2}},
		{"testDAG2", initDAG2, 1, 3,
			[]int{3, 4, 5}},
		{"testDAG2", initDAG2, 2, 0,
			[]int{}},
		{"testDAG3", initDAG3, 0, 6,
			[]int{0, 1, 2, 3, 4, 5}},
		{"testDAG3", initDAG3, 1, 0,
			[]int{}},
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

func TestConnCompComponents(t *testing.T) {
	for _, test := range testDataCCSize {
		t.Run(test.name, func(t *testing.T) {
			g := test.initGraphFunc()
			cc := InitConnectedComponents(g)
			t.Logf("Checking connected components from graph: %v", g)
			comps := cc.Component(test.component)
			t.Logf("Expected component: [%v], actual component: [%v]", test.expectedComps, comps)
			if !reflect.DeepEqual(test.expectedComps, comps) {
				t.Fail()
			}
		})
	}
}
