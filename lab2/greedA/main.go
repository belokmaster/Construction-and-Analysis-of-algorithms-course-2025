package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Структура для представления ребра с весом
type Edge struct {
	node   string
	weight float64
}

// Функция для реализации жадного алгоритма
func greedyAlg(graph map[string]map[string]float64, start, end string) []string {
	// Очередь для хранения текущих путей
	queue := [][]string{{start}}
	// Множество для отслеживания посещенных узлов
	visited := make(map[string]bool)

	for len(queue) > 0 {
		// Извлекаем текущий путь из очереди
		currentPath := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		currentNode := currentPath[len(currentPath)-1]

		// Если последняя вершина пути - целевая, возвращаем путь
		if currentNode == end {
			return currentPath
		}

		// Перебираем соседей последней вершины текущего пути
		for node := range graph[currentNode] {
			// Если путь через этот узел ещё не посещался
			if !visited[node] {
				// Добавляем текущий путь обратно в очередь
				queue = append(queue, currentPath)
				// Добавляем новый путь с присоединённым узлом в очередь
				newPath := append(currentPath, node)
				queue = append(queue, newPath)
				// Отмечаем путь как посещённый
				visited[node] = true
				// Прерываем цикл, чтобы учитывать жадность алгоритма (берём первый подходящий вариант)
				break
			}
		}
	}

	return nil
}

// Функция для считывания графа
func readGraph() map[string]map[string]float64 {
	graph := make(map[string]map[string]float64)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		inputString := scanner.Text()
		if inputString == "" {
			break
		}

		// Разделяем входную строку на начальную вершину, конечную вершину и вес рёбер
		parts := strings.Split(inputString, " ")
		startVertex := parts[0]
		endVertex := parts[1]
		weight := 0.0
		fmt.Sscanf(parts[2], "%f", &weight)

		// Добавляем ребро в граф (ориентированный граф)
		if _, exists := graph[startVertex]; !exists {
			graph[startVertex] = make(map[string]float64)
		}
		graph[startVertex][endVertex] = weight

		// Добавляем конечную вершину в граф, если её ещё нет
		if _, exists := graph[endVertex]; !exists {
			graph[endVertex] = make(map[string]float64)
		}
	}

	// Сортируем список смежности по весу рёбер (для жадного алгоритма)
	for vertex := range graph {
		// Сортировка соседей по весу рёбер
		// В Go нет прямого аналога сортировки по ключу, поэтому используем простую сортировку
		// на основе весов для каждого соседа
		edges := graph[vertex]
		edgesArray := make([]Edge, 0, len(edges))

		for node, weight := range edges {
			edgesArray = append(edgesArray, Edge{node, weight})
		}

		// Сортируем по весу
		for i := 0; i < len(edgesArray); i++ {
			for j := i + 1; j < len(edgesArray); j++ {
				if edgesArray[i].weight > edgesArray[j].weight {
					edgesArray[i], edgesArray[j] = edgesArray[j], edgesArray[i]
				}
			}
		}

		// Перезаписываем граф с отсортированными соседями
		graph[vertex] = make(map[string]float64)
		for _, edge := range edgesArray {
			graph[vertex][edge.node] = edge.weight
		}
	}

	return graph
}

func main() {
	// Читаем начальную и конечную вершины из первой строки
	var start, end string
	fmt.Scanf("%s %s\n", &start, &end)

	// Считываем граф из ввода
	graph := readGraph()

	// Запускаем жадный алгоритм и выводим найденный путь
	path := greedyAlg(graph, start, end)
	fmt.Println(strings.Join(path, ""))
}
