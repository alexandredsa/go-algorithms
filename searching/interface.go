package searching

type ArraySearcher interface {
	Search(arr []int, target int) int
}

type GraphSearcher interface {
	Name() string
	Search(graph [][]int, start int, target int) bool
}
