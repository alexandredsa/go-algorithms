package sorting

/**
Quick Sort is a divide-and-conquer algorithm that selects a pivot,
partitions the array, and recursively sorts sub-arrays.

1. Choose a pivot:
Select a pivot element from the array.
This pivot can be chosen in different ways, such as selecting the last element,
the first element, or a random element.

2. Partitioning:
Rearrange the array elements such that all elements less than the pivot are moved
to the left side of the pivot, and all elements greater than the pivot are moved to the right side.
The pivot element is now in its sorted position.

3. Recursive calls:
Recursively repeat the above steps for the sub-arrays on the left and right sides of the pivot.
This means applying the partitioning step to each sub-array until the entire array is sorted.

4. Combine the results: As the recursive calls finish, the sub-arrays become sorted.
The sorted sub-arrays are combined to form the final sorted array.

Time complexity: O(n log n).
*/

type QuickSort struct {
}

func (q QuickSort) Sort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivotIndex := q.partition(arr)
	q.Sort(arr[:pivotIndex])
	q.Sort(arr[pivotIndex+1:])
}

func (q QuickSort) partition(arr []int) int {
	pivot := arr[len(arr)-1]
	i := 0

	for j := 0; j < len(arr)-1; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]
	return i
}
