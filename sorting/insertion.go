package sorting

/**
Insertion Sort is a simple sorting algorithm that builds the final sorted array
one item at a time.

1. Iterate through the array:
   - Starting from the second element, compare it with the previous elements
     and insert it into the correct position in the sorted subarray.

2. Repeat until the entire array is sorted.

Time complexity: O(n^2), where n is the number of elements in the array.
In the best case, when the array is already sorted, the time complexity can be O(n).
**/

type InsertionSort struct{}

func (i InsertionSort) Sort(arr []int) {
	for j := 1; j < len(arr); j++ {
		key := arr[j]
		i := j - 1

		for i >= 0 && arr[i] > key {
			arr[i+1] = arr[i]
			i--
		}

		arr[i+1] = key
	}
}
