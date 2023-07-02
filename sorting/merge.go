package sorting

/**
Merge Sort is a divide-and-conquer algorithm that divides the array into two halves,
recursively sorts the sub-arrays, and then merges them.

1. Divide:
Divide the unsorted array into two halves by finding the middle index.

2. Recursive calls:
Recursively apply Merge Sort to the left and right halves of the array.

3. Merge:
Merge the two sorted sub-arrays back into a single sorted array.
This involves comparing elements from both sub-arrays and placing them in the correct order.

Time complexity: O(n log n).
**/

type MergeSort struct {
}

func (m MergeSort) Sort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2
	left := make([]int, mid)
	right := make([]int, len(arr)-mid)

	copy(left, arr[:mid])
	copy(right, arr[mid:])

	m.Sort(left)
	m.Sort(right)

	m.merge(arr, left, right)
}

func (m MergeSort) merge(arr, left, right []int) {
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}
