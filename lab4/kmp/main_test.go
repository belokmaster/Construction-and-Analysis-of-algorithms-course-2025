package main

import (
	"testing"
)

func TestKmp(t *testing.T) {
	tests := []struct {
		text     string
		pattern  string
		expected []int
	}{
		{
			text:     "abcabcabc",
			pattern:  "abc",
			expected: []int{0, 3, 6}, // "abc" встречается трижды в "abcabcabc"
		},
		{
			text:     "hello world",
			pattern:  "world",
			expected: []int{6}, // "world" встречается один раз на позиции 6
		},
		{
			text:     "abcde",
			pattern:  "fgh",
			expected: []int{}, // Нет вхождений
		},
		{
			text:     "aaaaa",
			pattern:  "aa",
			expected: []int{0, 1, 2, 3}, // "aa" встречается 4 раза
		},
		{
			text:     "",
			pattern:  "a",
			expected: []int{}, // Пустой текст, подстрока не найдена
		},
	}

	for _, tt := range tests {
		t.Run(tt.text+"_"+tt.pattern, func(t *testing.T) {
			result := kmp(tt.text, tt.pattern)
			if !equal(result, tt.expected) {
				t.Errorf("For text '%s' and pattern '%s, expected %v, but got %v", tt.text, tt.pattern, tt.expected, result)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
