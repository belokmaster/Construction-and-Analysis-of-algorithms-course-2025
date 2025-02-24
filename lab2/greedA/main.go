package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Edge struct {
	node   string
	weight float64
}

func greedyAlg(graph map[string]map[string]float64, start, end string) ([]string, int) {
	queue := [][]string{{start}}
	visited := make(map[string]bool)
	operations := 0 // Счётчик операций

	for len(queue) > 0 {
		operations++ // Операция извлечения пути из очереди
		currentPath := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		currentNode := currentPath[len(currentPath)-1]

		if currentNode == end {
			return currentPath, operations
		}

		for node := range graph[currentNode] {
			operations++ // Операция обхода рёбер
			if !visited[node] {
				queue = append(queue, currentPath) // Операция добавления в очередь
				newPath := append(currentPath, node)
				queue = append(queue, newPath) // Операция добавления нового пути в очередь

				visited[node] = true
				break
			}
		}
	}

	return nil, operations
}

func readGraph() map[string]map[string]float64 {
	graph := make(map[string]map[string]float64)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		inputString := scanner.Text()
		if inputString == "" {
			break
		}

		parts := strings.Split(inputString, " ")
		startVertex := parts[0]
		endVertex := parts[1]
		weight := 0.0
		fmt.Sscanf(parts[2], "%f", &weight)

		if _, exists := graph[startVertex]; !exists {
			graph[startVertex] = make(map[string]float64)
		}
		graph[startVertex][endVertex] = weight

		if _, exists := graph[endVertex]; !exists {
			graph[endVertex] = make(map[string]float64)
		}
	}

	for vertex := range graph {
		edges := graph[vertex]
		edgesArray := make([]Edge, 0, len(edges))

		for node, weight := range edges {
			edgesArray = append(edgesArray, Edge{node, weight})
		}

		for i := 0; i < len(edgesArray); i++ {
			for j := i + 1; j < len(edgesArray); j++ {
				if edgesArray[i].weight > edgesArray[j].weight {
					edgesArray[i], edgesArray[j] = edgesArray[j], edgesArray[i]
				}
			}
		}

		graph[vertex] = make(map[string]float64)
		for _, edge := range edgesArray {
			graph[vertex][edge.node] = edge.weight
		}
	}

	return graph
}

func main() {
	var start, end string
	fmt.Scanf("%s %s\n", &start, &end)

	startTime := time.Now()

	graph := readGraph()

	path, operations := greedyAlg(graph, start, end)

	duration := time.Since(startTime)

	fmt.Println("Path:", strings.Join(path, ""))
	fmt.Println("Operations:", operations)
	fmt.Printf("Execution Time: %s\n", duration)
}
