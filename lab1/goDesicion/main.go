package main

import (
	"fmt"
)

type Square struct {
	x, y, w int // Координаты (x, y) и размер квадрата (w)
}

type Board struct {
	size          int      // Размер игрового поля
	tempX, tempY  int      // Временные координаты для размещения квадратов
	minSquares    int      // Минимальное количество квадратов, необходимых для покрытия поля
	curSquareSize int      // Текущий размер квадрата
	countSquares  int      // Количество размещенных квадратов
	board         [][]int  // Игровое поле (матрица)
	bestBoard     [][]int  // Лучшее игровое поле (для хранения решения)
	squares       []Square // Список всех размещенных квадратов
	bestSolution  []Square // Лучшее решение (список квадратов)
}

// Конструктор для создания нового игрового поля.
func NewBoard(size int) *Board {
	board := make([][]int, size)
	for i := 0; i < size; i++ {
		board[i] = make([]int, size)
	}

	return &Board{
		board:        board,
		size:         size,
		squares:      []Square{},
		bestSolution: []Square{},
		countSquares: 0,
		minSquares:   size * size,
	}
}

// создаются три начальных квадрата, каждый из которых имеет определённый размер и положение на поле.
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

// Функция isBoardFull проверяет, заполнилось ли игровое поле
// (или если есть место для добавления новых квадратов),
// и пытается добавить новые квадраты в свободные участки поля.
func (b *Board) isBoardFull() bool {
	// Здесь происходит проверка, если curSquareSize (текущий размер квадрата) равен 0.
	// Это может быть на самом начале или после удаления квадрата.
	// Если так, то устанавливается размер квадрата как половина размера поля (b.size / 2).
	if b.curSquareSize == 0 {
		b.curSquareSize = b.size / 2
	}

	// Поиск свободных клеток на поле
	for i := b.size / 2; i < b.size; i++ {
		for j := b.size / 2; j < b.size; j++ {
			if b.board[i][j] == 0 {
				b.tempY = i
				b.tempX = j

				// Здесь проверяется, не превышает ли количество уже размещенных
				// квадратов (значение b.countSquares) минимальное количество квадратов
				if b.countSquares >= b.minSquares {
					return false
				}

				// В этом блоке кода происходит попытка разместить новый квадрат
				for {
					// Если добавление нового квадрата прошло успешно (функция newSquare возвращает true),
					// то размер следующего квадрата сбрасывается обратно на половину размера поля (b.size / 2),
					// и цикл прерывается с помощью break.
					if b.newSquare() {
						b.curSquareSize = b.size / 2
						break
						// Если квадрат не удалось разместить (функция newSquare возвращает false),
						// то размер текущего квадрата уменьшается на 1 (b.curSquareSize--),
						// и программа пытается снова разместить квадрат меньшего размера.
					} else {
						b.curSquareSize--
					}
				}
			}
		}
	}

	return true
}

// Функция newSquare пытается разместить новый квадрат на поле.
// Она проверяет, можно ли разместить квадрат в выбранном месте (с учетом размеров поля и уже занятых клеток).
// Если все проверки проходят успешно, квадрат размещается, а информация о нем сохраняется в массиве.
// Если разместить квадрат не удается, функция возвращает false.
func (b *Board) newSquare() bool {
	// Проверка, вмещается ли квадрат в поле
	if b.tempX+b.curSquareSize > b.size || b.tempY+b.curSquareSize > b.size {
		return false
	}

	// Проверка свободных клеток для квадрата
	for i := b.tempY; i < b.tempY+b.curSquareSize; i++ {
		for j := b.tempX; j < b.tempX+b.curSquareSize; j++ {
			if b.board[i][j] != 0 {
				return false
			}
		}
	}

	// Размещение квадрата на поле
	b.countSquares++
	for i := b.tempY; i < b.tempY+b.curSquareSize; i++ {
		for j := b.tempX; j < b.tempX+b.curSquareSize; j++ {
			b.board[i][j] = b.countSquares
		}
	}

	// Добавление квадрата в список
	b.squares = append(b.squares, Square{b.tempX, b.tempY, b.curSquareSize})

	return true
}

func (b *Board) backtrace() bool {
	// Получение последнего квадрата
	lastSquare := b.squares[len(b.squares)-1]

	// Проверка на условия отката
	// Длина списка квадратов b.squares должна быть больше 3 (то есть есть хотя бы 4 квадрата на поле).
	// Размер последнего квадрата (lastSquare.w) должен быть равен 1.
	for len(b.squares) > 3 && lastSquare.w == 1 {
		b.deleteSquare(b.squares[len(b.squares)-1].x, b.squares[len(b.squares)-1].y, b.squares[len(b.squares)-1].w)
		lastSquare = b.squares[len(b.squares)-1]
	}

	if len(b.squares) <= 3 {
		return false
	}

	b.deleteSquare(b.squares[len(b.squares)-1].x, b.squares[len(b.squares)-1].y, b.squares[len(b.squares)-1].w)
	b.curSquareSize = lastSquare.w - 1

	return true
}

func (b *Board) canDeleteSquare(x, y, w int) bool {
	if x < 0 || x+w > b.size || y < 0 || y+w > b.size {
		return false
	}

	if len(b.squares) == 0 {
		return false
	}

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
}

func (b *Board) printCoords() {
	fmt.Printf("%d\n", b.minSquares)

	for i := 0; i < b.minSquares; i++ {
		temp := b.bestSolution[len(b.bestSolution)-1]
		b.bestSolution = b.bestSolution[:len(b.bestSolution)-1]
		fmt.Printf("%d %d %d\n", temp.x+1, temp.y+1, temp.w)
	}
}

func (b *Board) process() {
	switch {
	case b.size%2 == 0:
		fmt.Println("4")

		part := b.size / 2
		fmt.Printf("1 1 %d\n", part)
		fmt.Printf("%d 1 %d\n", part+1, part)
		fmt.Printf("1 %d %d\n", part+1, part)
		fmt.Printf("%d %d %d\n", part+1, part+1, part)
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
}

func main() {
	var N int
	fmt.Scan(&N)

	board := NewBoard(N)
	board.process()
}
