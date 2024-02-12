package main

import (
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
