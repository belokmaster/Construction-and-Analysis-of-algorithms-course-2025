package main

import (
	"math"
)

var (
	finalPath []int
	finalRes  int
)

func firstMin(matrix [][]int, i int) int {
	min := math.MaxInt32

	for k := 0; k < len(matrix); k++ {
		if matrix[i][k] < min && i != k {
			min = matrix[i][k]
		}
	}

	return min
}

func secondMin(matrix [][]int, i int) int {
	first, second := math.MaxInt32, math.MaxInt32

	for j := 0; j < len(matrix); j++ {
		if i == j {
			continue
		}

		if matrix[i][j] <= first {
			second = first
			first = matrix[i][j]
		} else if matrix[i][j] <= second && matrix[i][j] != first {
			second = matrix[i][j]
		}
	}

	return second
}

func copyToFinal(currPath []int, n int) {
	finalPath = make([]int, n+1)
	copy(finalPath, currPath)
	finalPath[n] = currPath[0]
}

func TSPRec(matrix [][]int, currBound, currWeight, level int, currPath []int, visited []bool) {
	n := len(matrix)

	if level == n {
		if matrix[currPath[level-1]][currPath[0]] != 0 {
			currRes := currWeight + matrix[currPath[level-1]][currPath[0]]
			if currRes < finalRes {
				copyToFinal(currPath, n)
				finalRes = currRes
			}
		}
		return
	}

	for i := 0; i < n; i++ {
		if matrix[currPath[level-1]][i] != 0 && !visited[i] {
			temp := currBound
			currWeight += matrix[currPath[level-1]][i]

			if level == 1 {
				currBound -= (firstMin(matrix, currPath[level-1]) + firstMin(matrix, i)) / 2
			} else {
				currBound -= (secondMin(matrix, currPath[level-1]) + firstMin(matrix, i)) / 2
			}

			if currBound+currWeight < finalRes {
				currPath[level] = i
				visited[i] = true
				TSPRec(matrix, currBound, currWeight, level+1, currPath, visited)
			}

			currWeight -= matrix[currPath[level-1]][i]
			currBound = temp

			visited = make([]bool, n)
			for j := 0; j < level; j++ {
				if currPath[j] != -1 {
					visited[currPath[j]] = true
				}
			}
		}
	}
}

func TSP(matrix [][]int) Path {
	n := len(matrix)
	currBound := 0
	currPath := make([]int, n+1)
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		currBound += (firstMin(matrix, i) + secondMin(matrix, i))
	}

	currBound = int(math.Ceil(float64(currBound) / 2))

	visited[0] = true
	currPath[0] = 0

	finalRes = math.MaxInt32
	TSPRec(matrix, currBound, 0, 1, currPath, visited)

	return Path{cities: finalPath, distance: finalRes}
}
