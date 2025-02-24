package main

import "math"

type Path struct {
	cities   []int
	distance int
}

func GreedyTSP(matrix [][]int, start int) Path {
	n := len(matrix)
	visited := make([]bool, n)
	visited[start] = true
	path := []int{start}
	totalDistance := 0
	currentCity := start

	for len(path) < n {
		nextCity := -1
		minDist := math.MaxInt32

		for i := 0; i < n; i++ {
			if !visited[i] && matrix[currentCity][i] < minDist {
				nextCity = i
				minDist = matrix[currentCity][i]
			}
		}

		visited[nextCity] = true
		path = append(path, nextCity)
		totalDistance += minDist
		currentCity = nextCity
	}

	totalDistance += matrix[currentCity][start]
	path = append(path, start)

	return Path{cities: path, distance: totalDistance}
}
