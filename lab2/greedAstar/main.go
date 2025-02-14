package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strings"
)

type Edge struct {
	node   string
	weight float64
}

type Item struct {
	path      []string
	cost      float64
	heuristic float64
	index     int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost+pq[i].heuristic < pq[j].cost+pq[j].heuristic
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func heuristic(node, goal string) float64 {
	return math.Abs(float64(node[0] - goal[0]))
}

func aStar(graph map[string]map[string]float64, start, end string) []string {
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{path: []string{start}, cost: 0, heuristic: heuristic(start, end)})

	visited := make(map[string]float64)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)
		currentNode := current.path[len(current.path)-1]

		if currentNode == end {
			return current.path
		}

		if val, exists := visited[currentNode]; exists && val <= current.cost {
			continue
		}
		visited[currentNode] = current.cost

		for neighbor, weight := range graph[currentNode] {
			newCost := current.cost + weight
			newPath := append([]string{}, current.path...)
			newPath = append(newPath, neighbor)
			heap.Push(pq, &Item{path: newPath, cost: newCost, heuristic: heuristic(neighbor, end)})
		}
	}

	return nil
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
		startVertex, endVertex := parts[0], parts[1]
		var weight float64
		fmt.Sscanf(parts[2], "%f", &weight)

		if _, exists := graph[startVertex]; !exists {
			graph[startVertex] = make(map[string]float64)
		}
		graph[startVertex][endVertex] = weight
	}

	return graph
}

func main() {
	var start, end string
	fmt.Scanf("%s %s\n", &start, &end)
	graph := readGraph()

	path := aStar(graph, start, end)
	if path != nil {
		fmt.Println(strings.Join(path, ""))
	} else {
		fmt.Println("No path found")
	}
}
