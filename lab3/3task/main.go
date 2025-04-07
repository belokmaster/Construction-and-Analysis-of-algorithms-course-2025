package main

import (
	"fmt"
	"math"
)

func branchAndBound(path []int, visited []bool, currentCost float64, N int, matrix [][]float64, bestCost *float64, bestPath *[]int) {
	if len(path) == N {
		if matrix[path[len(path)-1]][path[0]] == -1 {
			return
		}

		totalCost := currentCost + matrix[path[len(path)-1]][path[0]]
		if totalCost < *bestCost {
			*bestCost = totalCost
			*bestPath = make([]int, len(path))
			copy(*bestPath, path)
		}

		return
	}

	for nextCity := 0; nextCity < N; nextCity++ {
		if !visited[nextCity] && matrix[path[len(path)-1]][nextCity] != -1 {
			lowerBound := currentCost + matrix[path[len(path)-1]][nextCity]
			if lowerBound >= *bestCost {
				continue
			}

			visited[nextCity] = true
			path = append(path, nextCity)

			branchAndBound(path, visited, lowerBound, N, matrix, bestCost, bestPath)

			visited[nextCity] = false
			path = path[:len(path)-1]
		}
	}
}

func main() {
	var N int
	fmt.Scan(&N)

	matrix := make([][]float64, N)
	for i := 0; i < N; i++ {
		matrix[i] = make([]float64, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	bestCost := math.MaxFloat64
	var bestPath []int

	path := make([]int, 0, N)
	visited := make([]bool, N)
	path = append(path, 0)
	visited[0] = true

	branchAndBound(path, visited, 0, N, matrix, &bestCost, &bestPath)

	for i, city := range bestPath {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(city)
	}

	fmt.Println()
	fmt.Printf("%.1f\n", bestCost)
}
