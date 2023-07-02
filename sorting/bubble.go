package sorting

/**
Bubble Sort is a simple comparison-based sorting algorithm that repeatedly steps through the list,
compares adjacent elements, and swaps them if they are in the wrong order.
The pass through the list is repeated until the list is sorted.

1. Compare adjacent elements:
Starting from the beginning of the list, compare each pair of adjacent elements.

2. Swap if necessary:
If the elements are in the wrong order, swap them.

3. Repeat until sorted:
Continue making passes through the list until no more swaps are needed,
indicating that the list is now sorted.

Time complexity: O(n^2).

Bubble Sort is not efficient for large data sets and is mainly used for educational purposes or small-sized arrays.
**/

type BubbleSort struct {
}

func (b BubbleSort) Sort(arr []int) {
	n := len(arr)
	swapped := true

	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
		n--
	}
}
