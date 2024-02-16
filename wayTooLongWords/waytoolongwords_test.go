package main

import (
	"strings"
	"testing"
)

func TestProcessWord(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello", "hello"},               // Test case for a word with length <= 10
		{"abcdefghijk", "a9k"},           // Test case for a word with length > 10
		{"testing", "testing"},           // Test case for a word with length <= 10
		{"internationalization", "i18n"}, // Test case for a word with length > 10
		{"a", "a"},                       // Test case for a single-character word
	}

	for _, tc := range testCases {
		result := processWord(tc.input)
		if result != tc.expected {
			t.Errorf("Expected %s for input %s, but got %s", tc.expected, tc.input, result)
		}
	}
}

func BenchmarkProcessWord(b *testing.B) {
	// Benchmark the processWord function with a long word to simulate real-world usage
	longWord := strings.Repeat("a", 100)
	for i := 0; i < b.N; i++ {
		processWord(longWord)
	}
}
