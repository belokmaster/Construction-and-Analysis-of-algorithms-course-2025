package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	numCities = 10 // Количество городов
)

type Path struct {
	cities   []int
	distance int
}

var (
	finalPath []int
	finalRes  int
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

	// Вывод матрицы с подписями городов
	fmt.Println("Сгенерированная матрица расстояний:")

	// Вывод заголовков (номеров городов сверху)
	fmt.Printf("    ")
	for i := 0; i < n; i++ {
		fmt.Printf("%3d ", i)
	}
	fmt.Println()

	// Разделительная линия
	fmt.Printf("    ")
	for i := 0; i < n; i++ {
		fmt.Print("----")
	}
	fmt.Println()

	// Вывод строк с номерами городов слева
	for i := 0; i < n; i++ {
		fmt.Printf("%2d |", i) // Номер города слева
		for j := 0; j < n; j++ {
			fmt.Printf("%3d ", matrix[i][j])
		}
		fmt.Println()
	}

	return matrix
}

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

func greedyTSP(matrix [][]int, start int) Path {
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

func generateAndSolveTSP() {
	if numCities < 2 {
		fmt.Println("Количество городов должно быть хотя бы 2.")
		return
	}

	matrix := generateDistanceMatrix(numCities)

	// Решаем задачу с жадным алгоритмом
	start := rand.Intn(numCities)
	greedyResult := greedyTSP(matrix, start)

	fmt.Printf("\nРешение задачи коммивояжёра (жадный алгоритм):\n")
	fmt.Printf("Путь: ")
	for _, city := range greedyResult.cities {
		fmt.Printf("%d ", city)
	}
	fmt.Printf("\nОбщее расстояние: %d\n", greedyResult.distance)

	// Решаем задачу методом ветвей и границ
	branchAndBoundResult := TSP(matrix)

	fmt.Printf("\nРешение задачи коммивояжёра (метод ветвей и границ):\n")
	fmt.Printf("Путь: ")
	for _, city := range branchAndBoundResult.cities {
		fmt.Printf("%d ", city)
	}
	fmt.Printf("\nОбщее расстояние: %d\n", branchAndBoundResult.distance)
}

func main() {
	start := time.Now()
	fmt.Printf("Количество городов: %d\n", numCities)
	generateAndSolveTSP()
	fmt.Printf("\nВремя выполнения: %s\n", time.Since(start))
}
