package searching

// DFS implements the Depth-First Search algorithm to search for a target element in an integer array.
type DFS struct{}

// Search performs a depth-first search on the given integer array for the target element.
// It returns the index of the target element if found, or -1 if not found.
func (d DFS) Search(arr []int, target int) int {
	return dfs(arr, target, 0, len(arr)-1)
}

// dfs is a recursive function that performs the depth-first search.
func dfs(arr []int, target, start, end int) int {
	if start > end {
		return -1
	}

	mid := (start + end) / 2

	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return dfs(arr, target, mid+1, end) // Target is in the right half
	} else {
		return dfs(arr, target, start, mid-1) // Target is in the left half
	}
}
