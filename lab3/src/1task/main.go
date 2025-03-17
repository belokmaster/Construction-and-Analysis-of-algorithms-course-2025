package main

import (
	"fmt"
	"math"
)

const INF = math.MaxInt32

var n int
var graph [][]int
var dp [][]int
var parent [][]int

func tsp(mask, pos int) int {
	if mask == (1<<n)-1 {
		if graph[pos][0] != 0 {
			return graph[pos][0]
		}
		return INF
	}

	if dp[mask][pos] != -1 {
		return dp[mask][pos]
	}

	res := INF
	for city := 0; city < n; city++ {
		if (mask&(1<<city)) == 0 && graph[pos][city] != 0 {
			newCost := graph[pos][city] + tsp(mask|(1<<city), city)
			if newCost < res {
				res = newCost
				parent[mask][pos] = city
			}
		}
	}
	dp[mask][pos] = res
	return res
}

func findPath() []int {
	path := []int{0}
	mask, pos := 1, 0
	for len(path) < n {
		nextCity := parent[mask][pos]
		if nextCity == -1 {
			return nil
		}
		path = append(path, nextCity)
		mask |= (1 << nextCity)
		pos = nextCity
	}
	path = append(path, 0)
	return path
}

func main() {
	fmt.Scan(&n)
	graph = make([][]int, n)

	for i := range graph {
		graph[i] = make([]int, n)
		for j := range graph[i] {
			fmt.Scan(&graph[i][j])
		}
	}

	dp = make([][]int, 1<<n)
	parent = make([][]int, 1<<n)
	for i := range dp {
		dp[i] = make([]int, n)
		parent[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
			parent[i][j] = -1
		}
	}

	res := tsp(1, 0)
	if res == INF {
		fmt.Println("no path")
	} else {
		fmt.Println(res)
		path := findPath()
		if path == nil {
			fmt.Println("no path")
		} else {
			for _, city := range path {
				fmt.Print(city, " ")
			}
			fmt.Println()
		}
	}
}
