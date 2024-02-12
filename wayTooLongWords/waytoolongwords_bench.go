package main

import (
	"strings"
	"testing"
)

func BenchmarkProcessWord(b *testing.B) {
	// Benchmark the processWord function with a long word to simulate real-world usage
	longWord := strings.Repeat("a", 100)
	for i := 0; i < b.N; i++ {
		processWord(longWord)
	}
}
