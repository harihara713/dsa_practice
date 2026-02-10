package graph

import (
	"math"
)

/*
Dijkstra's algorithm is an algorithm for finding the shortest paths between nodes in a graph, which may represent, for example, road networks. It was conceived by computer scientist
Edsger W. Dijkstra in 1956 and published three years later. The algorithm exists in many variants; Dijkstra's original variant found the shortest path between two nodes, but a more
common variant fixes a single node as the "source" node and finds shortest paths from the source to all other nodes in the graph, producing a shortest-path tree.

*/

// Dijkstra's Shortest Path algorithm, find the shortest path between two nodes in a graph (both directed & undirected)
// cost indicate the adjacency representation of a weighted graph, n is number of vertices, source is the source vertex from which calculationg the shortest path
func Dijkstra(cost [][]int, n int, source int) []int {
	distance := make([]int, n+1) // keep track the minimum distance between source node/vertex and other nodes/vertices
	selected := make([]bool, n+1)

	// update the distance between source vertex and other vertices
	for i := 1; i <= n; i++ {
		distance[i] = cost[source][i]
	}

	distance[source] = 0
	selected[source] = true

	var min, u int
	for i := 1; i < n; i++ {
		min = math.MaxInt

		// choose the minimum distance node
		for k := 1; k <= n; k++ {
			if !selected[k] && distance[k] < min {
				min = distance[k]
				u = k
			}
		}

		selected[u] = true
		// perform relaxation from minimum distanced node
		for v := 1; v <= n; v++ {
			if !selected[v] && distance[u] < distance[v]-cost[u][v] {
				distance[v] = distance[u] + cost[u][v]
			}
		}
	}

	return distance
}
