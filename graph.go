package main

type Node struct {
	idx      int
	distance int32
}

type Graph interface {
	Reset(v int)
	Size() int
	Get(idx int) *Node
	Adjacent(idx int) []int
}

type CubicGraph struct {
	nodes []*Node
	size  int
}

func (g *CubicGraph) Reset(v int) {
	for _, n := range g.nodes {
		n.distance = -1
	}

	g.nodes[v].distance = 0
}

func (g *CubicGraph) Size() int {
	return g.size * g.size * g.size
}

func (g *CubicGraph) Get(idx int) *Node {
	return g.nodes[idx]
}

func (g *CubicGraph) Adjacent(idx int) []int {
	sizeSq := g.size * g.size

	i := idx / sizeSq
	j := (idx - sizeSq*i) / g.size
	k := idx - sizeSq*i - j*g.size

	result := make([]int, 0, 6)
	if i != 0 {
		result = append(result, idx-sizeSq)
	}
	if i != g.size-1 {
		result = append(result, idx+sizeSq)
	}
	if j != 0 {
		result = append(result, idx-g.size)
	}
	if j != g.size-1 {
		result = append(result, idx+g.size)
	}
	if k != 0 {
		result = append(result, idx-1)
	}
	if k != g.size-1 {
		result = append(result, idx+1)
	}

	return result
}

func cubicGraph(l int) *CubicGraph {
	nodes := make([]*Node, l*l*l)
	for i := 0; i < l*l*l; i++ {
		nodes[i] = &Node{idx: i}
	}

	return &CubicGraph{
		nodes: nodes,
		size:  l,
	}
}
