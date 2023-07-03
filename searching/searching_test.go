package searching_test

import (
	"testing"

	"github.com/alexandredsa/go-algorithms/searching"
)

type TestCase struct {
	description string
	input       []int
	target      int
	expected    int
}

func TestSearchingAlgorithms(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Empty array, returns -1",
			input:       []int{},
			target:      10,
			expected:    -1,
		},
		{
			description: "Single-element array, target found at index 0",
			input:       []int{1},
			target:      1,
			expected:    0,
		},
		{
			description: "Target element not found, returns -1",
			input:       []int{1, 2, 3, 4, 5},
			target:      6,
			expected:    -1,
		},
		{
			description: "Target element found at index 3",
			input:       []int{10, 20, 30, 40, 50},
			target:      40,
			expected:    3,
		},
	}

	searchers := []searching.Searcher{searching.BinarySearcher{}}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			for _, searcher := range searchers {
				result := searcher.Search(testCase.input, testCase.target)
				if result != testCase.expected {
					t.Errorf("Binary search failed. Expected index: %d, Got: %d", testCase.expected, result)
				}
			}
		})
	}
}
