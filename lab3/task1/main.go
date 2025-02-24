package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const (
	numCities = 10 // Количество городов
)

type Path struct {
	cities   []int
	distance int
}

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
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%3d ", matrix[i][j])
		}
		fmt.Println()
	}

	return matrix
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

func calculateLowerBound(matrix [][]int, path []int, n int) int {
	lowerBound := 0

	// Добавляем расстояния для уже выбранных рёбер
	for i := 0; i < len(path)-1; i++ {
		lowerBound += matrix[path[i]][path[i+1]]
	}

	// Добавляем минимальные рёбра для оставшихся городов
	visited := make([]bool, n)
	for _, city := range path {
		visited[city] = true
	}

	// Для каждого города, который ещё не был посещен, находим минимальное ребро
	for i := 0; i < n; i++ {
		if !visited[i] {
			minEdge := math.MaxInt32
			for j := 0; j < n; j++ {
				if !visited[j] && i != j && matrix[i][j] < minEdge {
					minEdge = matrix[i][j]
				}
			}
			lowerBound += minEdge
		}
	}

	return lowerBound
}

func branchAndBoundTSP(matrix [][]int) Path {
	n := len(matrix)
	bestPath := make([]int, n+1)
	bestDistance := math.MaxInt32

	// Используем стек для хранения состояний
	type State struct {
		path            []int
		currentDistance int
	}
	stack := []State{{path: []int{0}, currentDistance: 0}}

	for len(stack) > 0 {
		currentState := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(currentState.path) == n {
			currentState.currentDistance += matrix[currentState.path[len(currentState.path)-1]][currentState.path[0]]
			if currentState.currentDistance < bestDistance {
				bestDistance = currentState.currentDistance
				copy(bestPath, append(currentState.path, currentState.path[0]))
			}
			continue
		}

		lastCity := currentState.path[len(currentState.path)-1]
		for nextCity := 0; nextCity < n; nextCity++ {
			if !contains(currentState.path, nextCity) {
				newPath := append(currentState.path, nextCity)
				newDistance := currentState.currentDistance + matrix[lastCity][nextCity]

				// Вычисляем нижнюю границу
				lowerBound := calculateLowerBound(matrix, newPath, n)

				// Если нижняя граница меньше лучшего расстояния, продолжаем поиск
				if lowerBound < bestDistance {
					stack = append(stack, State{path: newPath, currentDistance: newDistance})
				}
			}
		}
	}

	return Path{cities: bestPath, distance: bestDistance}
}

func contains(path []int, city int) bool {
	for _, c := range path {
		if c == city {
			return true
		}
	}
	return false
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
	branchAndBoundResult := branchAndBoundTSP(matrix)

	fmt.Printf("\nРешение задачи коммивояжёра (метод ветвей и границ):\n")
	fmt.Printf("Путь: ")
	for _, city := range branchAndBoundResult.cities {
		fmt.Printf("%d ", city)
	}
	fmt.Printf("\nОбщее расстояние: %d\n", branchAndBoundResult.distance)

	// Визуализация графа
	visualizeGraph(matrix, branchAndBoundResult.cities)
}

func randomColor() color.RGBA {
	return color.RGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: 255,
	}
}

func visualizeGraph(matrix [][]int, path []int) {
	p := plot.New()

	pts := make(plotter.XYs, numCities)
	pointColors := make([]color.RGBA, numCities)
	for i := 0; i < numCities; i++ {
		pts[i].X = float64(rand.Intn(100))
		pts[i].Y = float64(rand.Intn(100))
		pointColors[i] = randomColor()
	}

	scatterData := make(plotter.XYs, numCities)
	for i := 0; i < numCities; i++ {
		scatterData[i] = pts[i]
	}
	line, err := plotter.NewScatter(scatterData)
	if err != nil {
		fmt.Println("Ошибка при добавлении точек на граф:", err)
		return
	}

	for i := range scatterData {
		line.GlyphStyle.Color = pointColors[i]
	}

	line.GlyphStyle.Radius = vg.Points(5)
	p.Add(line)

	for i := 0; i < numCities; i++ {
		for j := i + 1; j < numCities; j++ {
			if matrix[i][j] > 0 {
				lines := plotter.XYs{
					{X: pts[i].X, Y: pts[i].Y},
					{X: pts[j].X, Y: pts[j].Y},
				}

				linePlot, err := plotter.NewLine(lines)
				if err != nil {
					fmt.Println("Ошибка при создании линии:", err)
					return
				}
				linePlot.Color = color.Black
				p.Add(linePlot)
			}
		}
	}

	for i := 0; i < len(path)-1; i++ {
		lines := plotter.XYs{
			{X: pts[path[i]].X, Y: pts[path[i]].Y},
			{X: pts[path[i+1]].X, Y: pts[path[i+1]].Y},
		}

		linePlot, err := plotter.NewLine(lines)
		if err != nil {
			fmt.Println("Ошибка при создании линии для пути:", err)
			return
		}
		linePlot.Color = color.Black
		p.Add(linePlot)
	}

	err = p.Save(10*vg.Inch, 10*vg.Inch, "graph.png")
	if err != nil {
		fmt.Println("Ошибка при сохранении графа:", err)
		return
	}

	fmt.Println("Граф с путями сохранён в файл graph.png")
}

func main() {
	start := time.Now()
	fmt.Printf("Количество городов: %d\n", numCities)
	generateAndSolveTSP()
	fmt.Printf("\nВремя выполнения: %s\n", time.Since(start))
}
