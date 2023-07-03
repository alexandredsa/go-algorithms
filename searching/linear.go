package searching

// LinearSearcher implements the Linear Search algorithm to search for a target element in an integer array.
type LinearSearcher struct{}

// Search performs a linear search on the given integer array for the target element.
// It returns the index of the target element if found, or -1 if not found.
func (ls LinearSearcher) Search(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i // Target found at index i
		}
	}

	return -1 // Target not found in the array
}
