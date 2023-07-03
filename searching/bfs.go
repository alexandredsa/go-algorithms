package searching

// BreadthFirstSearcher implements the Breadth-First Search algorithm to search for a target element in a graph represented by an adjacency list.
type BreadthFirstSearcher struct {
	visited []bool
}

func (bfs BreadthFirstSearcher) Name() string {
	return "Breadth-First Search"
}

// Search performs a breadth-first search on the graph represented by the adjacency list to find the target element.
// It returns true if the target element is found, or false if not found.
func (bfs BreadthFirstSearcher) Search(graph [][]int, start int, target int) bool {
	if len(graph) == 0 {
		return false
	}

	bfs.visited = make([]bool, len(graph))
	queue := []int{start}

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]

		if vertex == target {
			return true // Target found
		}

		bfs.visited[vertex] = true

		for _, neighbor := range graph[vertex] {
			if !bfs.visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}

	return false // Target not found in the graph
}
