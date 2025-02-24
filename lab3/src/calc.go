package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateDistanceMatrix(n int) [][]int {
	if n < 2 {
		panic("Количество городов должно быть хотя бы 2")
	}

	rand.Seed(time.Now().UnixNano())

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := rand.Intn(100) + 1
			matrix[i][j] = dist
			matrix[j][i] = dist
		}
	}

	fmt.Println("Сгенерированная матрица расстояний:")

	fmt.Printf("    ")
	for i := 0; i < n; i++ {
		fmt.Printf("%3d ", i)
	}
	fmt.Println()

	fmt.Printf("    ")
	for i := 0; i < n; i++ {
		fmt.Print("----")
	}
	fmt.Println()

	for i := 0; i < n; i++ {
		fmt.Printf("%2d |", i)
		for j := 0; j < n; j++ {
			fmt.Printf("%3d ", matrix[i][j])
		}
		fmt.Println()
	}

	return matrix
}

func generateAndSolveTSP() {
	if numCities < 2 {
		fmt.Println("Количество городов должно быть хотя бы 2.")
		return
	}

	matrix := generateDistanceMatrix(numCities)

	start := rand.Intn(numCities)
	greedyResult := GreedyTSP(matrix, start)

	fmt.Printf("\nРешение задачи коммивояжёра (жадный алгоритм):\n")
	fmt.Printf("Путь: ")
	for _, city := range greedyResult.cities {
		fmt.Printf("%d ", city)
	}
	fmt.Printf("\nОбщее расстояние: %d\n", greedyResult.distance)

	branchAndBoundResult := TSP(matrix)

	fmt.Printf("\nРешение задачи коммивояжёра (метод ветвей и границ):\n")
	fmt.Printf("Путь: ")
	for _, city := range branchAndBoundResult.cities {
		fmt.Printf("%d ", city)
	}
	fmt.Printf("\nОбщее расстояние: %d\n", branchAndBoundResult.distance)
}
