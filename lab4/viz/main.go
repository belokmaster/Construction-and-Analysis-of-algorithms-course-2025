package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Step представляет один шаг выполнения алгоритма KMP
type Step struct {
	TextIndex      int    `json:"textIndex"`
	PatternIndex   int    `json:"patternIndex"`
	Match          bool   `json:"match"`
	Shift          bool   `json:"shift"`
	Status         string `json:"status"`
	FailureValue   int    `json:"failureValue,omitempty"`
	Comparisons    int    `json:"comparisons"`
	PrefixFunction []int  `json:"prefixFunction,omitempty"` // Динамическая информация о совпадениях
}

// KMPResult представляет результат работы алгоритма KMP с шагами
type KMPResult struct {
	FailureFunction []int  `json:"failureFunction"`
	Steps           []Step `json:"steps"`
	Found           bool   `json:"found"`
	Positions       []int  `json:"positions"`
	Comparisons     int    `json:"comparisons"`
	Error           string `json:"error,omitempty"`
}

// buildFailureFunction строит функцию префиксов для подстроки
func buildFailureFunction(pattern string) []int {
	m := len(pattern)
	failure := make([]int, m)
	failure[0] = 0
	j := 0

	for i := 1; i < m; i++ {
		if pattern[i] == pattern[j] {
			j++
			failure[i] = j
		} else {
			for j > 0 && pattern[i] != pattern[j] {
				j = failure[j-1]
			}
			if pattern[i] == pattern[j] {
				j++
			}
			failure[i] = j
		}
	}
	return failure
}

// searchKMP выполняет поиск подстроки в тексте с шагами для анимации
func searchKMP(text, pattern string) KMPResult {
	if pattern == "" {
		return KMPResult{Error: "Pattern cannot be empty"}
	}
	if text == "" {
		return KMPResult{Error: "Text cannot be empty"}
	}

	failure := buildFailureFunction(pattern)
	positions := []int{}
	steps := []Step{}
	comparisons := 0
	i := 0 // индекс в тексте
	j := 0 // индекс в подстроке
	n := len(text)
	m := len(pattern)

	// Создаем массив для хранения максимальных длин совпадений для каждого символа текста
	maxPrefix := make([]int, n)

	for i < n {
		comparisons++
		match := text[i] == pattern[j]
		step := Step{
			TextIndex:      i,
			PatternIndex:   j,
			Match:          match,
			Shift:          false,
			Comparisons:    comparisons,
			PrefixFunction: make([]int, n),
		}

		// Копируем предыдущие максимальные значения
		copy(step.PrefixFunction, maxPrefix)

		if match {
			// Обновляем максимальную длину совпадения для текущей позиции
			if j+1 > maxPrefix[i] {
				maxPrefix[i] = j + 1
			}
			step.PrefixFunction[i] = maxPrefix[i]
			step.Status = "Match. Advancing."
			steps = append(steps, step)
			i++
			j++
			if j == m {
				positions = append(positions, i-j)
				// Обновляем все позиции, где было полное совпадение
				for k := i - m; k < i; k++ {
					if k >= 0 && k < n {
						maxPrefix[k] = k - (i - m) + 1 // Длина совпадения от начала
					}
				}
				step = Step{
					TextIndex:      i - 1,
					PatternIndex:   j - 1,
					Match:          true,
					Shift:          true,
					Status:         "Pattern found! Shifting.",
					FailureValue:   failure[j-1],
					Comparisons:    comparisons,
					PrefixFunction: make([]int, n),
				}
				copy(step.PrefixFunction, maxPrefix)
				steps = append(steps, step)
				j = failure[j-1]
			}
		} else if j > 0 {
			step.Status = "Mismatch. Consulting failure function."
			step.FailureValue = failure[j-1]
			steps = append(steps, step)
			j = failure[j-1]
			step = Step{
				TextIndex:      i,
				PatternIndex:   j,
				Match:          false,
				Shift:          true,
				Status:         "Shifting pattern.",
				Comparisons:    comparisons,
				PrefixFunction: make([]int, n),
			}
			copy(step.PrefixFunction, maxPrefix)
			steps = append(steps, step)
		} else {
			step.Status = "Mismatch. Shifting one position."
			steps = append(steps, step)
			i++
		}
	}

	return KMPResult{
		FailureFunction: failure,
		Steps:           steps,
		Found:           len(positions) > 0,
		Positions:       positions,
		Comparisons:     comparisons,
	}
}

// handleIndex обслуживает главную страницу
func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	tmpl.Execute(w, nil)
}

// handleKMP обрабатывает запросы к API для выполнения KMP
func handleKMP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Text    string `json:"text"`
		Pattern string `json:"pattern"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	result := searchKMP(input.Text, input.Pattern)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/kmp", handleKMP)

	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
