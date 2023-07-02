package sorting

/**
Selection Sort is an in-place comparison sorting algorithm.

1. Find the minimum element:
    - Start with the first element as the minimum.
    - Iterate through the remaining elements to find the minimum element.

2. Swap with the current element:
    - Swap the minimum element with the current element.

3. Repeat for the remaining unsorted portion:
    - Move to the next position and repeat the above steps until the entire array is sorted.

Time complexity: O(n^2).

Selection Sort has a consistent time complexity regardless of the input data.
It performs a constant number of swaps, which can be beneficial when memory writes are expensive.
**/

type SelectionSort struct{}

func (s SelectionSort) Sort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		minIndex := i

		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
