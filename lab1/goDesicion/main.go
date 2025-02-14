package main

import (
	"fmt"
	"time"
)

type Square struct {
	x, y, w int
}

type Board struct {
	size          int
	tempX, tempY  int
	minSquares    int
	curSquareSize int
	countSquares  int
	board         [][]int
	bestBoard     [][]int
	squares       []Square
	bestSolution  []Square

	// Счётчики для времени и операций
	operationsCount int
	startTime       time.Time
}

func NewBoard(size int) *Board {
	board := make([][]int, size)
	for i := 0; i < size; i++ {
		board[i] = make([]int, size)
	}

	return &Board{
		board:           board,
		size:            size,
		squares:         []Square{},
		bestSolution:    []Square{},
		countSquares:    0,
		minSquares:      size * size,
		operationsCount: 0, // Инициализируем счётчик операций
	}
}

func (b *Board) initializeBoard() {
	halfSize := b.size / 2
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if i <= halfSize && j <= halfSize {
				b.board[i][j] = 1
			} else if j < halfSize {
				b.board[i][j] = 2
			} else if i < halfSize {
				b.board[i][j] = 3
			}
		}
	}

	b.countSquares = 3
	b.squares = []Square{
		{0, 0, halfSize + 1},
		{halfSize + 1, 0, halfSize},
		{0, halfSize + 1, halfSize},
	}

	b.tempX = halfSize
	b.tempY = halfSize
}

func (b *Board) isBoardFull() bool {
	if b.curSquareSize == 0 {
		b.curSquareSize = b.size / 2
	}

	for i := b.size / 2; i < b.size; i++ {
		for j := b.size / 2; j < b.size; j++ {
			if b.board[i][j] == 0 {
				b.tempY = i
				b.tempX = j

				if b.countSquares >= b.minSquares {
					return false
				}

				for {
					if b.newSquare() {
						b.curSquareSize = b.size / 2
						break
					} else {
						b.curSquareSize--
					}
				}
			}
		}
	}

	return true
}

func (b *Board) newSquare() bool {
	if b.tempX+b.curSquareSize > b.size || b.tempY+b.curSquareSize > b.size {
		return false
	}

	for i := b.tempY; i < b.tempY+b.curSquareSize; i++ {
		for j := b.tempX; j < b.tempX+b.curSquareSize; j++ {
			if b.board[i][j] != 0 {
				return false
			}
		}
	}

	b.countSquares++
	for i := b.tempY; i < b.tempY+b.curSquareSize; i++ {
		for j := b.tempX; j < b.tempX+b.curSquareSize; j++ {
			b.board[i][j] = b.countSquares
		}
	}

	b.squares = append(b.squares, Square{b.tempX, b.tempY, b.curSquareSize})

	// Увеличиваем счётчик операций
	b.operationsCount++

	return true
}

func (b *Board) backtrace() bool {
	lastSquare := b.squares[len(b.squares)-1]

	for len(b.squares) > 3 && lastSquare.w == 1 {
		b.deleteSquare(b.squares[len(b.squares)-1].x, b.squares[len(b.squares)-1].y, b.squares[len(b.squares)-1].w)
		lastSquare = b.squares[len(b.squares)-1]
	}

	if len(b.squares) <= 3 {
		return false
	}

	b.deleteSquare(b.squares[len(b.squares)-1].x, b.squares[len(b.squares)-1].y, b.squares[len(b.squares)-1].w)
	b.curSquareSize = lastSquare.w - 1

	// Увеличиваем счётчик операций
	b.operationsCount++

	return true
}

func (b *Board) canDeleteSquare(x, y, w int) bool {
	if x < 0 || x+w > b.size || y < 0 || y+w > b.size {
		return false
	}

	if len(b.squares) == 0 {
		return false
	}

	// Увеличиваем счётчик операций
	b.operationsCount++

	return true
}

func (b *Board) deleteSquare(x, y, w int) {
	if !b.canDeleteSquare(x, y, w) {
		return
	}

	for i := y; i < y+w; i++ {
		for j := x; j < x+w; j++ {
			b.board[i][j] = 0
		}
	}

	b.squares = b.squares[:len(b.squares)-1]
	b.countSquares--

	// Увеличиваем счётчик операций
	b.operationsCount++
}

func (b *Board) calculations() {
	b.initializeBoard()
	min := b.size * b.size
	b.minSquares = min

	for {
		if b.isBoardFull() {
			if b.countSquares < min {
				b.bestSolution = append([]Square{}, b.squares...)
				b.bestBoard = make([][]int, len(b.board))

				for i := range b.bestBoard {
					b.bestBoard[i] = make([]int, len(b.board[i]))
					copy(b.bestBoard[i], b.board[i])
				}

				min = b.countSquares
				b.minSquares = min
			}
		}

		if !b.backtrace() {
			break
		}
	}

	// Увеличиваем счётчик операций
	b.operationsCount++
}

func (b *Board) printCoords() {
	fmt.Printf("%d\n", b.minSquares)

	for i := 0; i < b.minSquares; i++ {
		temp := b.bestSolution[len(b.bestSolution)-1]
		b.bestSolution = b.bestSolution[:len(b.bestSolution)-1]
		fmt.Printf("%d %d %d\n", temp.x+1, temp.y+1, temp.w)
	}

	// Увеличиваем счётчик операций
	b.operationsCount++
}

func (b *Board) process() {
	startTime := time.Now() // Начинаем отсчёт времени
	b.startTime = startTime

	switch {
	case b.size%2 == 0:
		fmt.Println("4")

		part := b.size / 2
		fmt.Printf("1 1 %d\n", part)
		fmt.Printf("%d 1 %d\n", part+1, part)
		fmt.Printf("1 %d %d\n", part+1, part)
		fmt.Printf("%d %d %d\n", part+1, part+1, part)
		elapsedTime := time.Since(startTime)
		fmt.Printf("Время работы: %.8f сек.\n", elapsedTime.Seconds())
		fmt.Printf("Количество операций: %d\n", b.operationsCount)
		return
	case b.size%3 == 0:
		fmt.Println("6")

		part := b.size / 3
		fmt.Printf("1 1 %d\n", 2*part)
		fmt.Printf("%d 1 %d\n", 1+2*part, part)
		fmt.Printf("1 %d %d\n", 1+2*part, part)
		fmt.Printf("%d %d %d\n", 1+2*part, 1+part, part)
		fmt.Printf("%d %d %d\n", 1+part, 1+2*part, part)
		fmt.Printf("%d %d %d\n", 1+2*part, 1+2*part, part)
		return
	case b.size%5 == 0:
		fmt.Println("8")

		part := b.size / 5
		fmt.Printf("1 1 %d\n", 3*part)
		fmt.Printf("%d 1 %d\n", 1+3*part, 2*part)
		fmt.Printf("1 %d %d\n", 1+3*part, 2*part)
		fmt.Printf("%d %d %d\n", 1+3*part, 1+3*part, 2*part)
		fmt.Printf("%d %d %d\n", 1+2*part, 1+3*part, part)
		fmt.Printf("%d %d %d\n", 1+2*part, 1+4*part, part)
		fmt.Printf("%d %d %d\n", 1+3*part, 1+2*part, part)
		fmt.Printf("%d %d %d\n", 1+4*part, 1+2*part, part)
		return

	default:
		b.calculations()
		b.printCoords()
	}

	// Выводим количество операций и время работы
	elapsedTime := time.Since(startTime)
	fmt.Printf("Время работы: %.8f сек.\n", elapsedTime.Seconds())
	fmt.Printf("Количество операций: %d\n", b.operationsCount)
}

func main() {
	var N int
	fmt.Scan(&N)

	board := NewBoard(N)
	board.process()
}
