package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var operationCount int // Переменная для подсчета операций

func prefixFunction(s string) []int {
	p := make([]int, len(s))

	for i := 1; i < len(s); i++ {
		k := p[i-1]
		operationCount++ // Каждое присваивание или проверка считается операцией
		for k > 0 && s[i] != s[k] {
			k = p[k-1]
			operationCount++ // Операция в цикле
		}

		if s[i] == s[k] {
			p[i] = k + 1
		} else {
			p[i] = k
		}
		operationCount++ // Каждое присваивание или проверка
	}

	return p
}

func kmpSearch(text, pattern string) []int {
	var result []int
	combined := pattern + "@" + text
	p := prefixFunction(combined)
	patternLen := len(pattern)

	for i, v := range p {
		operationCount++ // Перебор элементов в префикс-функции
		if v == patternLen {
			result = append(result, i-patternLen*2)
		}
	}

	return result
}

func main() {
	// Засекаем время начала работы программы
	startTime := time.Now()

	// Чтение входных данных
	reader := bufio.NewReader(os.Stdin)
	pattern, _ := reader.ReadString('\n')
	text, _ := reader.ReadString('\n')

	pattern = strings.TrimSpace(pattern)
	text = strings.TrimSpace(text)

	// Поиск подстроки с использованием KMP
	indices := kmpSearch(text, pattern)

	// Засекаем время завершения работы программы
	elapsedTime := time.Since(startTime)

	// Выводим индексы вхождений или -1, если не найдено
	if len(indices) > 0 {
		fmt.Print(indices[0])
		for i := 1; i < len(indices); i++ {
			fmt.Print(",", indices[i])
		}
		fmt.Println()
	} else {
		fmt.Println(-1)
	}

	// Выводим время выполнения и количество операций
	fmt.Printf("Time taken: %s\n", elapsedTime)
	fmt.Printf("Operations count: %d\n", operationCount)
}
