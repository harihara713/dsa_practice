package graph

import "math"

// Return a minimum cost for a spanning tree for a undirected graph
// n -> size of the vertices
func PrimMST(graph [][]int, n int) (cost int) {
	selected := make([]bool, n+1)
	near := make([]int, n+1)

	u := 1
	var v int
	min := math.MaxInt

	for i := 1; i <= n; i++ {
		if graph[1][i] < min {
			min = graph[1][i]
			v = i
		}
	}

	selected[u] = true
	selected[v] = true
	cost += graph[u][v]

	for i := 1; i <= n; i++ {
		if graph[u][i] < graph[v][i] {
			near[i] = u
		} else {
			near[i] = v
		}
	}

	for i := 1; i <= n-2; i++ {
		min = math.MaxInt
		for j := 1; j <= n; j++ {
			if !selected[j] && graph[j][near[j]] < min {
				min = graph[j][near[j]]
				v = j
			}
		}

		u = near[v]
		cost += graph[u][v]
		selected[v] = true

		// Updating near, if the new added vertex has minimum cost to other vertices
		for k := 1; k <= n; k++ {
			if !selected[k] && graph[k][v] < graph[k][near[k]] {
				near[k] = v
			}
		}
	}

	return cost
}
