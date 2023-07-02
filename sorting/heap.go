package sorting

/**
Heap Sort is an in-place sorting algorithm that utilizes a binary heap data structure.

1. Build the heap:
Starting from the first non-leaf node (n/2-1), iterate backwards and heapify each node,
which ensures that the heap property is satisfied.

2. Extract the maximum repeatedly:
Swap the root element (maximum element) with the last element in the heap.
Reduce the heap size by one.
Heapify the new root to maintain the heap property.
Repeat this process until the heap is empty.

Time complexity: O(n log n).
**/

type HeapSort struct {
}

func (h HeapSort) Sort(arr []int) {
	n := len(arr)

	// Build the heap
	for i := n/2 - 1; i >= 0; i-- {
		h.heapify(arr, n, i)
	}

	// Extract the maximum repeatedly
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0] // Swap the root (maximum) with the last element
		h.heapify(arr, i, 0)            // Heapify the new root
	}
}

func (h HeapSort) heapify(arr []int, n, i int) {
	largest := i     // Assume the largest element is at the root
	left := 2*i + 1  // Left child index
	right := 2*i + 2 // Right child index

	// If the left child is larger than the root
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// If the right child is larger than the largest so far
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// If the largest element is not at the root
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i] // Swap the root with the largest element
		h.heapify(arr, n, largest)                  // Recursively heapify the affected sub-tree
	}
}
