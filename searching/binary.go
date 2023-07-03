package searching

// BinarySearcher implements the Binary Search algorithm to search for a target element in a sorted integer array.
type BinarySearcher struct{}

// Search performs a binary search on the given sorted integer array for the target element.
// It returns the index of the target element if found, or -1 if not found.
func (bs BinarySearcher) Search(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid // Target found at index mid
		} else if arr[mid] < target {
			low = mid + 1 // Target is in the right half
		} else {
			high = mid - 1 // Target is in the left half
		}
	}

	return -1 // Target not found in the array
}
