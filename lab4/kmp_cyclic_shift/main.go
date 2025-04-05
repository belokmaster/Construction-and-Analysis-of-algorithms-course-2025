package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func prefixFunction(s string) []int {
	p := make([]int, len(s))

	for i := 1; i < len(s); i++ {
		k := p[i-1]
		for k > 0 && s[i] != s[k] {
			k = p[k-1]
		}

		if s[i] == s[k] {
			p[i] = k + 1
		} else {
			p[i] = k
		}
	}

	return p
}

func isCyclicShift(sample, text string) int {
	if len(sample) != len(text) {
		return -1
	}

	if sample == text {
		return 0
	}

	prefixTable := prefixFunction(sample)
	patternIndex := 0

	for textInd := 0; textInd < 2*len(text); textInd++ {
		if text[textInd%len(text)] == sample[patternIndex] {
			patternIndex++

			if patternIndex == len(sample) {
				return textInd - len(sample) + 1
			}
			continue
		}

		if patternIndex == 0 {
			continue
		}

		patternIndex = prefixTable[patternIndex-1]
		textInd--
	}

	return -1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	a, _ := reader.ReadString('\n')
	b, _ := reader.ReadString('\n')

	a = strings.TrimSpace(a)
	b = strings.TrimSpace(b)

	result := isCyclicShift(b, a)
	fmt.Println(result)
}
