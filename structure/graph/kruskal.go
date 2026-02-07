package graph

import "math"

type Edge struct {
	Start  int
	End    int
	Weight int
}

// Returns minimum cost for a Spanning Tree in a Graph (Greedy Method), It is an Optimisation problem
// edges is the edges of the Graph, n is the number of vertices in the graph
func KruskalMST(edges []Edge, n int) (cost int) {
	selected := make([]bool, len(edges)) // to check if a edge is already selected or not
	var u, v int
	edgeCount := 0

	ds := NewUnionFind(n)
	var min, idx int

	for edgeCount < n-1 {
		min = math.MaxInt
		// find the minimum cost
		// we can use min heap instead of looping and find the minimum
		for i := 0; i < len(edges); i++ {
			if !selected[i] && edges[i].Weight < min {
				min = edges[i].Weight
				idx = i
			}
		}

		u = edges[idx].Start
		v = edges[idx].End

		if ds.Find(u) != ds.Find(v) {
			cost += edges[idx].Weight
			edgeCount++
			ds.Union(u, v)
		}
		selected[idx] = true
	}

	return cost
}
