package graphs

import "fmt"

type dfsOrder struct {
	DirectedGraph
	pre         []int
	post        []int
	preOrder    []int
	postOrder   []int
	marked      []bool
	preCounter  int
	postCounter int
}

func (g *dfsOrder) String() string {
	return fmt.Sprintf("dfsOrder("+
		"Graph: %v, pre: %v, post: %v, preOrder: %v, postOrder: %v, "+
		"marked: %v, preCounter: %v, postCounter: %v)", g.DirectedGraph,
		g.pre, g.post, g.preOrder, g.postOrder, g.marked, g.preCounter, g.postCounter)
}

// initDFSOrder inits the data structure that finds order in DirectedGraph
// using DFS
func initDFSOrder(g DirectedGraph) *dfsOrder {
	newG := new(dfsOrder)
	newG.DirectedGraph = g
	newG.pre = make([]int, g.VNum(), g.VNum())
	newG.post = make([]int, g.VNum(), g.VNum())
	newG.preOrder = make([]int, 0, g.VNum())
	newG.postOrder = make([]int, 0, g.VNum())
	newG.marked = make([]bool, g.VNum(), g.VNum())
	newG.preCounter = 0
	newG.postCounter = 0
	for v := 0; v < g.VNum(); v++ {
		if !newG.marked[v] {
			findOrderDFS(newG, v)
		}
	}
	return newG
}

func findOrderDFS(g *dfsOrder, v int) {
	g.marked[v] = true
	g.pre[v] = g.preCounter
	g.preCounter++
	g.preOrder = append(g.preOrder, v)
	for _, w := range g.Edges(v) {
		if !g.marked[w] {
			findOrderDFS(g, w)
		}
	}
	g.postOrder = append(g.postOrder, v)
	g.post[v] = g.postCounter
	g.postCounter++
}

// ReversePostOrder provides the reverse post-DFS order
func ReversePostOrder(g DirectedGraph) []int {
	return initDFSOrder(g).postOrder
}

// PostOrder provides the post-DFS order of vertices
func PostOrder(g DirectedGraph) []int {
	gNew := initDFSOrder(g)
	n := len(gNew.postOrder)
	revPost := make([]int, n)
	for i, v := range gNew.postOrder {
		revPost[n-i-1] = v
	}
	return revPost
}
