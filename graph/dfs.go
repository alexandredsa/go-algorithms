package graph

// DFS implements the Depth-First Search algorithm to search for a target element in a graph.
type DFS struct{}

// Search performs a depth-first search on the given graph starting from the specified start vertex.
// It returns true if the target element is found, or false otherwise.
func (dfs DFS) Search(graph [][]int, start int, target int) bool {
	if len(graph) == 0 {
		return false
	}

	stack := []int{start}
	visited := make([]bool, len(graph))
	visited[start] = true

	for len(stack) > 0 {
		vertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if vertex == target {
			return true // Target found
		}

		for _, neighbor := range graph[vertex] {
			if !visited[neighbor] {
				stack = append(stack, neighbor)
				visited[neighbor] = true
			}
		}
	}

	return false // Target not found in the graph
}
