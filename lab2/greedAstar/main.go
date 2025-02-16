package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Edge struct {
	dest   string
	weight float64
}

type Node struct {
	name      string
	cost      float64
	heuristic float64
	path      []string
}

func heuristic(current, end string) float64 {
	return math.Abs(float64(current[0]) - float64(end[0]))
}

func aStar(graph map[string][]Edge, start, end string) []string {
	openSet := []Node{{name: start, cost: 0, heuristic: heuristic(start, end), path: []string{start}}}
	visited := make(map[string]float64)

	for len(openSet) > 0 {
		bestIdx := 0
		for i, node := range openSet {
			if node.cost+node.heuristic < openSet[bestIdx].cost+openSet[bestIdx].heuristic {
				bestIdx = i
			}
		}

		current := openSet[bestIdx]
		openSet = append(openSet[:bestIdx], openSet[bestIdx+1:]...)

		if current.name == end {
			return current.path
		}

		visited[current.name] = current.cost

		for _, edge := range graph[current.name] {
			newCost := current.cost + edge.weight
			if oldCost, seen := visited[edge.dest]; seen && newCost >= oldCost {
				continue
			}

			newPath := append([]string{}, current.path...)
			newPath = append(newPath, edge.dest)
			openSet = append(openSet, Node{
				name:      edge.dest,
				cost:      newCost,
				heuristic: heuristic(edge.dest, end),
				path:      newPath,
			})
		}
	}

	return nil
}

func readGraph() (map[string][]Edge, string, string) {
	graph := make(map[string][]Edge)
	scanner := bufio.NewScanner(os.Stdin)

	var start, end string
	if scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%s %s", &start, &end)
	}

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) < 3 {
			continue
		}

		var weight float64
		fmt.Sscanf(parts[2], "%f", &weight)
		graph[parts[0]] = append(graph[parts[0]], Edge{dest: parts[1], weight: weight})
	}

	return graph, start, end
}

func main() {
	graph, start, end := readGraph()
	path := aStar(graph, start, end)

	fmt.Println(strings.Join(path, ""))
}
