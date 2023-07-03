package graph

// BFS implements the Breadth-First Search algorithm to search for a target element in a graph.
type BFS struct{}

// Search performs a breadth-first search on the given graph starting from the specified start vertex.
// It returns true if the target element is found, or false otherwise.
func (bfs BFS) Search(graph [][]int, start int, target int) bool {
	if len(graph) == 0 {
		return false
	}

	queue := []int{start}
	visited := make([]bool, len(graph))
	visited[start] = true

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]

		if vertex == target {
			return true // Target found
		}

		for _, neighbor := range graph[vertex] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
	}

	return false // Target not found in the graph
}
