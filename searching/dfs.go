package searching

// DepthFirstSearcher implements the Depth-First Search algorithm to search for a target element in a graph represented by an adjacency list.
type DepthFirstSearcher struct {
	visited []bool
}

func (dfs DepthFirstSearcher) Name() string {
	return "Depth-First Search"
}

// Search performs a depth-first search on the graph represented by the adjacency list to find the target element.
// It returns true if the target element is found, or false if not found.
func (dfs DepthFirstSearcher) Search(graph [][]int, start int, target int) bool {
	dfs.visited = make([]bool, len(graph))
	return dfs.dfs(graph, start, target)
}

// dfs is a recursive helper function that performs the depth-first search.
func (dfs DepthFirstSearcher) dfs(graph [][]int, vertex int, target int) bool {
	if len(graph) == 0 {
		return false
	}

	if vertex == target {
		return true // Target found
	}

	dfs.visited[vertex] = true

	for _, neighbor := range graph[vertex] {
		if !dfs.visited[neighbor] {
			if dfs.dfs(graph, neighbor, target) {
				return true // Target found in the neighbor
			}
		}
	}

	return false // Target not found in the graph
}
