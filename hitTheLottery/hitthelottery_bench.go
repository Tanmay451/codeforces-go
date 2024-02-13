package main

import (
	"testing"
)

func BenchmarkNumberOfBills(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numberOfBills(123)
	}
}
