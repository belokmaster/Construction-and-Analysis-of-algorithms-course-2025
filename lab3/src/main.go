package main

import (
	"fmt"
	"time"
)

const (
	numCities = 20
)

func main() {
	start := time.Now()
	fmt.Printf("Количество городов: %d\n", numCities)
	generateAndSolveTSP()
	fmt.Printf("\nВремя выполнения: %s\n", time.Since(start))
}
