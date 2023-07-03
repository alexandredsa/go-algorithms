package graph_test

import (
	"testing"

	"github.com/alexandredsa/go-algorithms/graph"
)

type GraphTestCase struct {
	description string
	graph       [][]int
	start       int
	target      int
	expected    bool
}

func TestGraphSearchingAlgorithms(t *testing.T) {
	testCases := []GraphTestCase{
		{
			description: "Empty graph, target not found",
			graph:       [][]int{},
			start:       0,
			target:      10,
			expected:    false,
		},
		{
			description: "Graph with single vertex, target found",
			graph:       [][]int{{}},
			start:       0,
			target:      0,
			expected:    true,
		},
		{
			description: "Graph with multiple vertices, target not found",
			graph:       [][]int{{1, 2}, {0, 3}, {0, 4}, {1}, {2}},
			start:       0,
			target:      5,
			expected:    false,
		},
		{
			description: "Graph with multiple vertices, target found",
			graph:       [][]int{{1, 2}, {0, 3}, {0, 4}, {1}, {2}},
			start:       0,
			target:      3,
			expected:    true,
		},
	}

	searchers := []graph.GraphSearcher{
		graph.DFS{},
		graph.BFS{},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			for _, searcher := range searchers {
				result := searcher.Search(testCase.graph, testCase.start, testCase.target)
				if result != testCase.expected {
					t.Errorf("Graph search failed. Expected result: %v, Got: %v", testCase.expected, result)
				}
			}
		})
	}
}
