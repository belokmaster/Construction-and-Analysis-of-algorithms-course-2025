package main

import (
	"testing"
)

func TestBranchAndBound(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]float64
		expected []int
		cost     float64
	}{
		{
			name: "Simple 3 cities",
			matrix: [][]float64{
				{-1, 1, 3},
				{3, -1, 1},
				{1, 2, -1},
			},
			expected: []int{0, 1, 2},
			cost:     3.0,
		},
		{
			name: "4 cities with clear path",
			matrix: [][]float64{
				{-1, 3, 4, 1},
				{1, -1, 3, 4},
				{9, 2, -1, 4},
				{8, 9, 2, -1},
			},
			expected: []int{0, 3, 2, 1},
			cost:     6.0,
		},
		{
			name: "Fully connected 2 cities",
			matrix: [][]float64{
				{-1, 5},
				{5, -1},
			},
			expected: []int{0, 1},
			cost:     10.0,
		},
		{
			name: "No possible path",
			matrix: [][]float64{
				{-1, -1},
				{-1, -1},
			},
			expected: nil,
			cost:     INF,
		},
		{
			name: "Asymmetric costs",
			matrix: [][]float64{
				{-1, 10, 15, 20},
				{5, -1, 9, 10},
				{6, 13, -1, 12},
				{8, 8, 9, -1},
			},
			expected: []int{0, 1, 3, 2},
			cost:     35.0,
		},
		{
			name: "One city has very high outgoing costs",
			matrix: [][]float64{
				{-1, 1, 100, 100},
				{100, -1, 1, 100},
				{100, 100, -1, 1},
				{1, 100, 100, -1},
			},
			expected: []int{0, 1, 2, 3},
			cost:     4.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			N = len(tt.matrix)
			matrix = tt.matrix
			bestPath = nil
			bestCost = INF

			path := make([]int, 0, N)
			visited := make([]bool, N)
			path = append(path, 0)
			visited[0] = true

			branchAndBound(path, visited, 0)

			if len(bestPath) != len(tt.expected) {
				t.Errorf("Expected path length %d, got %d", len(tt.expected), len(bestPath))
			} else {
				for i := range bestPath {
					if bestPath[i] != tt.expected[i] {
						t.Errorf("At position %d: expected %d, got %d", i, tt.expected[i], bestPath[i])
						break
					}
				}
			}

			if bestCost != tt.cost {
				t.Errorf("Expected cost %.1f, got %.1f", tt.cost, bestCost)
			}
		})
	}
}
