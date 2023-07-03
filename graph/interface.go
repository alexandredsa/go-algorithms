package graph

type GraphSearcher interface {
	Search(graph [][]int, start int, target int) bool
}
