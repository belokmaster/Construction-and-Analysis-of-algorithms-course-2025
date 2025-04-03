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

func kmpSearch(text, pattern string) []int {
	var result []int
	combined := pattern + "@" + text
	p := prefixFunction(combined)
	patternLen := len(pattern)

	for i, v := range p {
		if v == patternLen {
			result = append(result, i-patternLen*2)
		}
	}

	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	pattern, _ := reader.ReadString('\n')
	text, _ := reader.ReadString('\n')

	pattern = strings.TrimSpace(pattern)
	text = strings.TrimSpace(text)

	indices := kmpSearch(text, pattern)

	if len(indices) > 0 {
		fmt.Print(indices[0])
		for i := 1; i < len(indices); i++ {
			fmt.Print(",", indices[i])
		}
		fmt.Println()
	} else {
		fmt.Println(-1)
	}
}
