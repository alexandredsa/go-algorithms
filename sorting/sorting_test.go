package sorting

import (
	"reflect"
	"testing"
)

func TestSortingAlgorithms(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{64, 25, 12, 22, 11}, []int{11, 12, 22, 25, 64}},
		{[]int{9, 8, 7, 6, 5}, []int{5, 6, 7, 8, 9}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}, []int{1, 1, 2, 3, 3, 4, 5, 5, 6, 9}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 3, 2, 5, 4}, []int{1, 2, 3, 4, 5}},
		{[]int{10, 20, 30, 40, 50}, []int{10, 20, 30, 40, 50}},
		{[]int{50, 40, 30, 20, 10}, []int{10, 20, 30, 40, 50}},
	}

	sorters := []Sorter{
		QuickSort{},
		MergeSort{},
		HeapSort{},
		InsertionSort{},
	}

	for _, testCase := range testCases {
		arr := make([]int, len(testCase.input))
		copy(arr, testCase.input)

		for _, sorter := range sorters {
			sorter.Sort(arr)
			if !reflect.DeepEqual(arr, testCase.expected) {
				t.Errorf("Sorting failed. Expected: %v, Got: %v", testCase.expected, arr)
			}
		}
	}
}
