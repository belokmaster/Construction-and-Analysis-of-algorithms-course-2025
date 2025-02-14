package main

import (
	"fmt"
	"math/bits"
)

var N int
var board []int
var fullMask int
var bestCount int
var bestSolution []struct {
	x, y, w int
}
var currentSolution []struct {
	x, y, w int
}

type SegmentTree struct {
	tree []int
	n    int
}

func (st *SegmentTree) build(n int) {
	st.n = n
	st.tree = make([]int, 4*n)
}

func (st *SegmentTree) update(node, start, end, idx, value int) {
	if start == end {
		st.tree[node] = value
	} else {
		mid := (start + end) / 2
		leftChild := 2*node + 1
		rightChild := 2*node + 2
		if idx <= mid {
			st.update(leftChild, start, mid, idx, value)
		} else {
			st.update(rightChild, mid+1, end, idx, value)
		}
		st.tree[node] = st.tree[leftChild] | st.tree[rightChild]
	}
}

func (st *SegmentTree) query(node, start, end, L, R int) int {
	if R < start || end < L {
		return 0
	}
	if L <= start && end <= R {
		return st.tree[node]
	}
	mid := (start + end) / 2
	leftChild := 2*node + 1
	rightChild := 2*node + 2
	leftQuery := st.query(leftChild, start, mid, L, R)
	rightQuery := st.query(rightChild, mid+1, end, L, R)
	return leftQuery | rightQuery
}

// Поиск первого свободного места с использованием сегментного дерева
func findFirstFreeCell(st *SegmentTree) (int, int) {
	for i := 0; i < N; i++ {
		freeMask := ^board[i] & fullMask
		// Для каждой строки запрашиваем свободное место
		if freeMask != 0 {
			col := bits.TrailingZeros(uint(freeMask))
			return i, col
		}
	}
	return -1, -1
}

// Оценка нижней границы для числа прямоугольников
func estimateLowerBound(freeCells int) int {
	maxArea := (N - 1) * (N - 1)
	return (freeCells + maxArea - 1) / maxArea
}

// Оптимизация DFS с сегментным деревом
func dfs(countUsed int, st *SegmentTree) {
	if countUsed >= bestCount {
		return
	}

	row, column := findFirstFreeCell(st)

	if row == -1 {
		bestCount = countUsed
		bestSolution = append([]struct{ x, y, w int }{}, currentSolution...)
		return
	}

	freeCells := 0
	for i := 0; i < N; i++ {
		freeCells += bits.OnesCount(uint(^board[i] & fullMask))
	}

	if countUsed+estimateLowerBound(freeCells) >= bestCount {
		return
	}

	maxSize := min(N-row, N-column, N-1)

	for w := maxSize; w > 0; w-- {
		mask := (1<<w - 1) << column
		overlap := false
		for i := row; i < row+w; i++ {
			if board[i]&mask != 0 {
				overlap = true
				break
			}
		}
		if overlap {
			continue
		}

		for i := row; i < row+w; i++ {
			board[i] |= mask
		}

		currentSolution = append(currentSolution, struct {
			x, y, w int
		}{x: row + 1, y: column + 1, w: w})
		dfs(countUsed+1, st)
		currentSolution = currentSolution[:len(currentSolution)-1]

		for i := row; i < row+w; i++ {
			board[i] &^= mask
		}
	}
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func baseSituation(n int) {
	half := n / 2
	bestSolution = []struct{ x, y, w int }{
		{x: 1, y: 1, w: half},
		{x: 1, y: half + 1, w: half},
		{x: half + 1, y: 1, w: half},
		{x: half + 1, y: half + 1, w: half},
	}
}

func main() {
	//start := time.Now()

	fmt.Scan(&N)
	board = make([]int, N)
	fullMask = (1 << N) - 1
	bestCount = N * N

	st := SegmentTree{}
	st.build(N)

	if N%2 == 0 {
		baseSituation(N)

		fmt.Println(4)
		for _, solution := range bestSolution {
			fmt.Printf("%d %d %d\n", solution.x, solution.y, solution.w)
		}
		//elapsed := time.Since(start) // Время выполнения
		//fmt.Printf("Время выполнения: %.2f секунд\n", elapsed.Seconds())
		return
	}

	dfs(0, &st)

	//elapsed := time.Since(start)

	fmt.Println(bestCount)
	for _, solution := range bestSolution {
		fmt.Printf("%d %d %d\n", solution.x, solution.y, solution.w)
	}

	//fmt.Printf("Время выполнения: %.2f секунд\n", elapsed.Seconds())
}
